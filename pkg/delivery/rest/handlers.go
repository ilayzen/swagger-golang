package rest

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

type Handler struct {
}

func NewHandlers() *Handler {
	return &Handler{}
}

func (h *Handler) GetSwaggerByReadFiles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		distDir := "./dist"
		filePath := distDir + r.URL.Path[len("/dir"):]

		if r.URL.Path == "/dir" || r.URL.Path == "/dir/" {
			filePath = distDir + "/index.html"
		}

		render, err := os.ReadFile(filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		ext := filepath.Ext(filePath)
		contentType := mime.TypeByExtension(ext)
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}

		w.Write(render)

	}
}

func (h *Handler) GetSwaggerByDir() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		distDir := "./dist"
		fileServer := http.StripPrefix("/static", http.FileServer(http.Dir(distDir)))
		fileServer.ServeHTTP(w, r)
	}
}
