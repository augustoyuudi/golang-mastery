package httpadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"lesson/core"
)

type HttpUserAdapter struct {
	CorePort core.UserRepository
}

func (h *HttpUserAdapter) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.CorePort.GetUserById("user_id")

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching user: %s", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}