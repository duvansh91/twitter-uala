package writer

import (
	"encoding/json"
	"net/http"
	"twitter-uala/pkg/services/rest/dto"

	"github.com/gorilla/mux"
)

func (h *Handler) Follow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	followerID := vars["user_id"]
	w.Header().Set("Content-Type", "application/json")

	var request dto.FollowRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if request.UserToFollow == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("field user_id is required")
	}

	err = h.userService.Follow(ctx, followerID, request.UserToFollow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")
}
