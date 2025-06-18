package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
)

func IconHandler(w http.ResponseWriter, r *http.Request) {
	assetPath := strings.TrimPrefix(r.URL.Path, "/assets/icons/")

	if strings.Contains(assetPath, "..") || strings.Contains(assetPath, "/") {
		http.NotFound(w, r)
		return
	}

	if !strings.HasSuffix(assetPath, ".svg") {
		http.NotFound(w, r)
		return
	}

	fullPath := filepath.Join("assets/icons", assetPath)
	http.ServeFile(w, r, fullPath)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	assetPath := strings.TrimPrefix(r.URL.Path, "/assets/images/")

	if strings.Contains(assetPath, "..") || strings.Contains(assetPath, "/") {
		http.NotFound(w, r)
		return
	}

	if !strings.HasSuffix(assetPath, ".png") {
		http.NotFound(w, r)
		return
	}

	fullPath := filepath.Join("assets/images", assetPath)
	http.ServeFile(w, r, fullPath)
}
