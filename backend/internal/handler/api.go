package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Pertsaa/pokkne/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type APIHandler struct {
	ctx     context.Context
	queries *database.Queries
}

func NewAPIHandler(ctx context.Context, pool *pgxpool.Pool) *APIHandler {
	return &APIHandler{
		ctx:     ctx,
		queries: database.New(pool),
	}
}

type APIError struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%d: %v", e.Code, e.Message)
}

func NewAPIError(code int, message any) APIError {
	return APIError{
		Code:    code,
		Message: message,
	}
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func Make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				writeJSON(w, apiErr.Code, apiErr)
			} else {
				internalErr := map[string]any{
					"code":    http.StatusInternalServerError,
					"message": "internal server error",
				}
				writeJSON(w, http.StatusInternalServerError, internalErr)
			}
			slog.Error("handler error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func writeJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}
