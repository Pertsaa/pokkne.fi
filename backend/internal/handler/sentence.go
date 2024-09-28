package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Pertsaa/pokkne/internal/database"
)

func (h *APIHandler) SentenceListHandler(w http.ResponseWriter, r *http.Request) error {
	data, err := h.queries.ListSentences(h.ctx)
	if err != nil {
		return err
	}

	sentences := make([]database.Sentence, 0)
	sentences = append(sentences, data...)

	return writeJSON(w, http.StatusOK, sentences)
}

func (h *APIHandler) SentenceCreateHandler(w http.ResponseWriter, r *http.Request) error {
	var body SentenceCreate
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return NewAPIError(http.StatusBadRequest, "invalid json body")
	}
	defer r.Body.Close()

	if body.Sentence == "" {
		return writeJSON(w, http.StatusBadRequest, "body contains empty sentence")
	}

	sentence, err := h.queries.CreateSentence(h.ctx, body.Sentence)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, sentence)
}

func (h *APIHandler) SentenceBatchCreateHandler(w http.ResponseWriter, r *http.Request) error {
	var body SentenceBatchCreate
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return NewAPIError(http.StatusBadRequest, "invalid json body")
	}
	defer r.Body.Close()

	for _, s := range body.Sentences {
		if s == "" {
			return writeJSON(w, http.StatusBadRequest, "body contains empty sentences")
		}
	}

	count, err := h.queries.BatchCreateSentences(h.ctx, body.Sentences)
	if err != nil {
		return err
	}

	response := SentenceBatchCreateResponse{Count: count}

	return writeJSON(w, http.StatusCreated, response)
}
