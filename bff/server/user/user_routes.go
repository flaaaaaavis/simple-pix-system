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

func (r *RouterImplementation) GetUserById(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var body types.GetUserByIdRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error decoding req body: %v", err)
		return
	}

	res, err := r.backend.GetUserById(ctx, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error calling GetUserById: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func (r *RouterImplementation) ListUsers(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	res, err := r.backend.ListUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error calling ListUsers: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func (r *RouterImplementation) UpdateUserById(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var body types.UpdateUserRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error decoding req body: %v", err)
		return
	}

	res, err := r.backend.UpdateUserById(ctx, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error calling UpdateUserById: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func (r *RouterImplementation) CreateContact(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var body types.CreateContactRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error decoding req body: %v", err)
		return
	}

	res, err := r.backend.CreateContact(ctx, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error calling CreateContact: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func (r *RouterImplementation) GetContactById(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var body types.GetContactByIdRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error decoding req body: %v", err)
		return
	}

	res, err := r.backend.GetContactById(ctx, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error calling GetContactById: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func (r *RouterImplementation) UpdateContactById(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var body types.UpdateContactByIdRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error decoding req body: %v", err)
		return
	}

	res, err := r.backend.UpdateContactById(ctx, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error calling UpdateContactById: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}
