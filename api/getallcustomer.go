package api

import (
	"encoding/json"
	"net/http"

	"github.com/darkhunterLearning/Golang-CRUD-API/action"
	"github.com/darkhunterLearning/Golang-CRUD-API/db"
	"github.com/darkhunterLearning/Golang-CRUD-API/model"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	db := db.Set_UpDB()

	action.PrintMessage("Getting customers...")

	// Get all customers from customer table
	rows, err := db.Query("SELECT * FROM customer")

	// check errors
	action.CheckErr(err)

	// var response []JsonResponse
	var customers []model.Customer

	// Foreach customer
	for rows.Next() {
		var (
			id               int
			uniqueID         string
			customerName     string
			customerPhone    string
			customerAddress  string
			customerEmail    string
			customerPassword string
		)

		err = rows.Scan(&id, &uniqueID, &customerName, &customerPhone, &customerAddress, &customerEmail, &customerPassword)

		// check errors
		action.CheckErr(err)

		customers = append(customers, model.Customer{CustomerID: id, CustomerUniqueId: uniqueID, CustomerName: customerName, CustomerPhone: customerPhone,

			CustomerAddress: customerAddress, CustomerEmail: customerEmail, CustomerPassword: customerPassword})
	}

	var response = model.JsonResponse{Type: "success", Data: customers, Message: "200"}

	json.NewEncoder(w).Encode(response)

}
