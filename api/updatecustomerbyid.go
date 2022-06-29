package api

import (
	"encoding/json"
	"net/http"

	"github.com/darkhunterLearning/Golang-CRUD-API/action"
	"github.com/darkhunterLearning/Golang-CRUD-API/db"
	"github.com/darkhunterLearning/Golang-CRUD-API/model"
	"github.com/gorilla/mux"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	var response = model.JsonResponse{}

	if ID == "" {
		response = model.JsonResponse{Type: "error", Message: "You are missing ID parameter."}
	} else {
		db := db.Set_UpDB()
		Customer_Name := r.FormValue("customer_name")
		Customer_Phone := r.FormValue("customer_phone")
		Customer_Address := r.FormValue("customer_address")
		Customer_Email := r.FormValue("customer_email")
		Customer_Password := r.FormValue("customer_password")

		sqlUpdate := `
		UPDATE customer
		SET customer_name = $1, customer_phone = $2, customer_address = $3, customer_email = $4, customer_password = $5
		WHERE id = $6;`
		_, err := db.Exec(sqlUpdate, Customer_Name, Customer_Phone, Customer_Address, Customer_Email, Customer_Password, ID)

		action.CheckErr(err)

		response = model.JsonResponse{Type: "success", Message: "Customer's Info has been updated successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
