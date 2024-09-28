package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"math/rand"
)

type Word struct {
	Word string
	End  bool
}

func (h *APIHandler) ChatHandler(w http.ResponseWriter, r *http.Request) error {
	var body ChatRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return NewAPIError(http.StatusBadRequest, "invalid json body")
	}
	defer r.Body.Close()

	sentences, err := h.queries.ListScoredSentences(h.ctx, body.Prompt)
	if err != nil {
		return err
	}

	ids := make([]int32, 0)
	for _, s := range sentences {
		ids = append(ids, s.ID+1)
	}

	responses, err := h.queries.BatchGetSentencesByIds(h.ctx, ids)
	if err != nil {
		return err
	}

	response := ""

	if body.Mode == "randomized" {
		indexedWords := make([][]Word, 0)

		for _, r := range responses {
			lower := strings.ToLower(r.Content)

			normalized := regexp.MustCompile(`[^a-zA-ZåäöÅÄÖ0-9 ]+`).ReplaceAllString(lower, "")

			words := strings.Fields(normalized)
			for i, w := range words {
				if len(indexedWords) < i+1 {
					indexedWords = append(indexedWords, make([]Word, 0))
				}
				indexedWords[i] = append(indexedWords[i], Word{Word: w, End: i == len(words)-1})
			}
		}

		for _, words := range indexedWords {
			word := words[rand.Intn(len(words))]
			response = fmt.Sprintf("%s %s", response, word.Word)
			if word.End && rand.Intn(100) < 50 {
				break
			}
		}
	} else if len(responses) > 0 {
		response = responses[rand.Intn(len(responses))].Content
	}

	response = strings.TrimSpace(response)
	response = strings.ToUpper(response[:1]) + response[1:]

	return writeJSON(w, http.StatusOK, ChatResponse{Message: response})
}
