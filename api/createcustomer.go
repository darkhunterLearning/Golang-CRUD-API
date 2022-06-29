package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/darkhunterLearning/Golang-CRUD-API/action"
	"github.com/darkhunterLearning/Golang-CRUD-API/db"
	"github.com/darkhunterLearning/Golang-CRUD-API/model"
	"github.com/google/uuid"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	Unique_Id := uuid.New()
	Customer_Name := r.FormValue("customer_name")
	Customer_Phone := r.FormValue("customer_phone")
	Customer_Address := r.FormValue("customer_address")
	Customer_Email := r.FormValue("customer_email")
	Customer_Password := r.FormValue("customer_password")

	var response = model.JsonResponse{}

	if Customer_Email == "" || Customer_Name == "" {
		response = model.JsonResponse{Type: "error", Message: "You are missing Customer_Email or Customer_Name parameter."}
	} else {
		db := db.Set_UpDB()

		action.PrintMessage("Inserting customer into DB")

		fmt.Println("Inserting new customer:" + Customer_Name)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO customer(unique_id, customer_name, customer_phone, customer_address, customer_email, customer_password) VALUES($1, $2, $3, $4, $5, $6) returning id;",
			Unique_Id.String(), Customer_Name, Customer_Phone, Customer_Address, Customer_Email, Customer_Password).Scan(&lastInsertID)

		// check errors
		action.CheckErr(err)

		response = model.JsonResponse{Type: "success", Message: "Customer has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
