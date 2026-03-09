package helper

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("failed to encode json response", "error", err)
	}
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	if status == http.StatusInternalServerError {
		slog.Error("internal server error", "message", msg)
	}

	WriteJSON(w, status, map[string]interface{}{
		"code":    status,
		"message": msg,
	})
}

func WriteValidationError(w http.ResponseWriter, status int, errs map[string]string) {
	WriteJSON(w, status, map[string]interface{}{
		"code":   status,
		"errors": errs,
	})
}
func ReadJSON(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}
