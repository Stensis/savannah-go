package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-service/db"
	"go-service/model"
	"go-service/sms"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Handle order creation
func CreateOrder(res http.ResponseWriter, req *http.Request) {

	var order model.Order

	if err := json.NewDecoder(req.Body).Decode(&order); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(res, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO orders (customer_id, item, amount) VALUES ($1, $2, $3)", order.CustomerID, order.Item, order.Amount)
	if err != nil {
		log.Printf("Failed to insert order into database: %v", err)
		http.Error(res, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// Fetch the customer's phone number
	phoneNumber, err := GetCustomerPhoneNumber(db.DB, order.CustomerID)
	if err != nil {
		log.Printf("failed to get customer phone number: %v", err)
		http.Error(res, "failed to get customer phone number", http.StatusInternalServerError)
		return
	}

	apiOrderResponse := model.ApiResponse{
		Message: "Order created successfully",
		Status:  http.StatusCreated,
	}

	// send sms
	smsMessage := "Order item " + order.Item + " of amount KES." + fmt.Sprint(order.Amount) + " had be added successfully!"
	smsResponse, err := sms.SendSMS(phoneNumber, smsMessage, os.Getenv("ENVIRONMENT"))
	if err != nil {
		log.Printf("Failed to send order sms: %v", err)
		http.Error(res, "Failed to send order using sms service: ", http.StatusInternalServerError)
		return
	}

	log.Println("[SMS SERVICE]: SMS Details: ", smsResponse)

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(apiOrderResponse)
}

// GetCustomerPhoneNumber fetches the phone number of the customer based on customer ID
func GetCustomerPhoneNumber(db *sql.DB, customerID int) (string, error) {
	var phoneNumber string
	err := db.QueryRow("SELECT phone_number FROM customers WHERE id = $1", customerID).Scan(&phoneNumber)
	if err != nil {
		return "", fmt.Errorf("failed to get phone number: %v", err)
	}
	return phoneNumber, nil
}
