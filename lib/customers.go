package lib

import (
	"encoding/json"
	"go-service/db"
	"go-service/model"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Handle customer creation
func CreateCustomer(res http.ResponseWriter, req *http.Request) {

	var customer model.Customer

	if err := json.NewDecoder(req.Body).Decode(&customer); err != nil {
		http.Error(res, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO customers (name, code, phone_number) VALUES ($1, $2, $3)", customer.Name, customer.Code, customer.PhoneNUmber)
	if err != nil {
		log.Println("Error inserting customer:", err)
		http.Error(res, "Failed to create customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	apiCustomerResponse := model.ApiResponse{
		Message: "Customer added successfully",
		Status:  http.StatusCreated,
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(apiCustomerResponse)
}
