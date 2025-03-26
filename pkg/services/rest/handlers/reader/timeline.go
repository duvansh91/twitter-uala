package reader

import (
	"encoding/json"
	"net/http"
	"twitter-uala/pkg/services/rest/dto"

	"github.com/gorilla/mux"
)

func (h *Handler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	userID := vars["user_id"]

	w.Header().Set("Content-Type", "application/json")

	tweets, err := h.userService.GetTweetsTimeline(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	content := dto.NewTimelineResponse(tweets)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(content)
}
