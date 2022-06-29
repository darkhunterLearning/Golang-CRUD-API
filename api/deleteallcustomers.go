package api

import (
	"encoding/json"
	"net/http"

	"github.com/darkhunterLearning/Golang-CRUD-API/action"
	"github.com/darkhunterLearning/Golang-CRUD-API/db"
	"github.com/darkhunterLearning/Golang-CRUD-API/model"
)

func DeleteCustomers(w http.ResponseWriter, r *http.Request) {
	db := db.Set_UpDB()
	action.PrintMessage("Deleting all customer!")
	_, err := db.Exec("DELETE FROM customer")
	action.CheckErr(err)
	action.PrintMessage("All customers have been deleted successfully")
	var response = model.JsonResponse{Type: "success", Message: "All customers have been deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
