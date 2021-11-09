package presentation

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func FromJson(source io.Reader, data interface{}) error {
	dec := json.NewDecoder(source)
	return dec.Decode(data)
}

func Json(res http.ResponseWriter, data interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(data); sendError(res, err, http.StatusInternalServerError) {
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(buf.Bytes())
}

func sendError(res http.ResponseWriter, err error, code int) bool {
	if err == nil {
		return false
	}

	http.Error(res, err.Error(), code)
	return true
}
