package handler

import (
	"fmt"
	"io"
	"net/http"
)

func (h handler) handlerUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusBadRequest)
		return
	}

	if (r.Body == nil) {
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	shortURL := h.encoderUrl.EncodeURL(string(body))
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	link := h.cfg.BaseURL + "/" + shortURL
	fmt.Fprint(w, link)
}