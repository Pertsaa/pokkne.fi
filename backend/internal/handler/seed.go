package handler

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func (h *APIHandler) SeedHandler(w http.ResponseWriter, r *http.Request) error {
	content, err := os.ReadFile("data/out.txt")
	if err != nil {
		log.Fatal(err)
	}
	textContent := string(content)

	sentences := strings.Split(textContent, "\n")

	count, err := h.queries.BatchCreateSentences(h.ctx, sentences)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, SentenceBatchCreateResponse{Count: count})
}
