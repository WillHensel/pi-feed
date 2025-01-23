package main

import (
	"net/http"
	"strings"
)

func hlsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		pathFragments := strings.Split(path, "/")
		end := pathFragments[len(pathFragments)-1]

		if strings.HasSuffix(end, ".m3u8") {
			w.Header().Set("Content-Type", "vnd.apple.mpegURL")
		} else if strings.HasSuffix(end, ".ts") {
			w.Header().Set("Content-Type", "video/MP2T")
		}

		next.ServeHTTP(w, r)
	})
}
