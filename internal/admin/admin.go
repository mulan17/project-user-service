package admin

import (
	"encoding/json"
	"net/http"
	"internal/user"
)

// Функція для перегляду списку покупців
func ViewCustomers(w http.ResponseWriter, r *http.Request) {
	customers := user.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

// Функція для блокування покупця
func BlockCustomer(w http.ResponseWriter, r *http.Request) {
	var data struct {
		CustomerID string `json:"customer_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.BlockCustomer(data.CustomerID)
	w.Write([]byte("Customer blocked successfully"))
}

