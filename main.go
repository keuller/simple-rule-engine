package main

import (
	"github.com/keuller/pricing-service/internal/infrastructure/database"
	"github.com/keuller/pricing-service/internal/infrastructure/http"
)

func main() {
	if err := database.NewConnection(); err != nil {
		panic(err)
	}

	server := http.BuildHTTPServer("0.0.0.0", "8000")
	go server.Start()
	server.Stop()
}
