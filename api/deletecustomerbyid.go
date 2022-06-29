package api

import (
	"encoding/json"
	"net/http"

	"github.com/darkhunterLearning/Golang-CRUD-API/action"
	"github.com/darkhunterLearning/Golang-CRUD-API/db"
	"github.com/darkhunterLearning/Golang-CRUD-API/model"
	"github.com/gorilla/mux"
)

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	var response = model.JsonResponse{}

	if ID == "" {
		response = model.JsonResponse{Type: "error", Message: "You are missing ID parameter."}
	} else {
		db := db.Set_UpDB()

		action.PrintMessage("Deleting customer from DB")

		_, err := db.Exec("DELETE FROM customer where id = $1", ID)

		// check errors
		action.CheckErr(err)

		response = model.JsonResponse{Type: "success", Message: "Customer has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
