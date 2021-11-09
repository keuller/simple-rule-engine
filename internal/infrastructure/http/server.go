package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type webServer struct {
	Server *http.Server
	Host   string
	Port   string
}

func BuildHTTPServer(host, port string) webServer {
	return webServer{nil, host, port}
}

func (ws *webServer) SetMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Heartbeat("/info"))
}

func (ws *webServer) Start() {
	router := chi.NewRouter()
	ws.SetMiddlewares(router)
	registerRoutes(router)

	ws.Server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", ws.Host, ws.Port),
		Handler:      router,
		IdleTimeout:  15 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("[INFO] Server is ready on http://%s\n", ws.Server.Addr)

	err := ws.Server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("[ERROR] Cannot initiate the server, reason: %s \n", err.Error())
	}
}

func (ws *webServer) Stop() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	<-ctx.Done()

	log.Println("[INFO] Server is shutting down...")
	// stop receiving any request.
	if err := ws.Server.Shutdown(context.Background()); err != nil {
		log.Fatalf("[FAIL] Fail on stop the server, reason: %s\n", err.Error())
	}
	log.Println("[INFO] Server has been stopped.")
	stop()
}
