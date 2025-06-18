package handlers

import (
	"net/http"
	"path/filepath"
)

func ContactPageHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/contact" {
		http.NotFound(w, req)
		return
	}

	http.ServeFile(w, req, filepath.Join("static/pages", "contact.html"))
}
