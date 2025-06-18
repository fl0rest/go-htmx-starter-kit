package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
)

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	var (
		filePath string
		fullPath string
	)

	if strings.Contains(req.URL.Path, "css") {
		filePath = strings.TrimPrefix(req.URL.Path, "/static/css/")
		fullPath = filepath.Join("static/css", filePath)
	} else if strings.Contains(req.URL.Path, "js") {
		filePath = strings.TrimPrefix(req.URL.Path, "/static/js/")
		fullPath = filepath.Join("static/js", filePath)
	}

	ext := filepath.Ext(filePath)
	switch ext {
	case ".css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case ".js":
		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	}

	http.ServeFile(w, req, fullPath)
}
