package handler

import (
	"net/http"
	"strings"
)

func (h handler) handlerHash(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusBadRequest)
		return
	}
	path := r.URL.Path
	hash := strings.Split(path, "/")[1]
	url, err := h.encoderUrl.DecodeUrl(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := h.cfg.BaseURL + "/"+ url
	w.Header().Set("Location", result)	
	w.WriteHeader(http.StatusTemporaryRedirect)
}