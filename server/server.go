package server

import (
	"net/http"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

func DefaultHandler(html string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "text/html")
		w.Write([]byte(html))
		return
	}
}
