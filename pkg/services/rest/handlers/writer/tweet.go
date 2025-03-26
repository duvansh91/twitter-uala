package writer

import (
	"encoding/json"
	"net/http"
	"twitter-uala/pkg/services/rest/dto"

	"github.com/gorilla/mux"
)

func (h *Handler) PublishTweet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	userID := vars["user_id"]

	var request dto.PublishTweetRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	w.Header().Set("Content-Type", "application/json")

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("param user_id is required")
	}

	err = h.userService.PublishTweet(ctx, request.Content, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")
}
