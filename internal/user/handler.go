package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserResource struct {
	S *Storage
}

func (p *UserResource) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := p.S.GetAllUsers()
	res := map[string][]User{"users": users}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		fmt.Println("Failed to encode: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (p *UserResource) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, ok := p.S.GetUserById(id)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *UserResource) CreateUser(w http.ResponseWriter, r *http.Request) {
	var reqBody User

	userId, ok := p.S.CreateUser(reqBody)

	if !ok {
		fmt.Print("Error creating user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBody.ID = userId

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		fmt.Println("Failed to decode: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (p *UserResource) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var reqBody User

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		fmt.Println("Failed to encode: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := p.S.UpdateUser(id, reqBody)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
