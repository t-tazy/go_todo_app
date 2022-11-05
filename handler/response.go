package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// 統一したJSONフォーマットでエラーを返すための型
type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

// JSONのレスポンスを書き込むヘルパー
func RespondJSON(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=uft-8")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rsp := ErrResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			fmt.Printf("write error response error: %v", err)
		}
		return
	}

	w.WriteHeader(status)
	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		fmt.Printf("write response error: %v", err)
	}
}
