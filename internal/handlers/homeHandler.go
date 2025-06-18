package handlers

import (
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	http.ServeFile(w, req, filepath.Join("static/pages", "index.html"))
}
