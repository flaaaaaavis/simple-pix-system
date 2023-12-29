package user

import (
	"context"
	"encoding/json"
	"log"
	"mentoria/bff/types"
	"net/http"
)

func (r *RouterImplementation) CreateUser(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var body types.CreateUserRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error decoding req body: %v", err)
		return
	}

	res, err := r.backend.CreateUser(ctx, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error calling CreateUser: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}
