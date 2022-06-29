package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darkhunterLearning/Golang-CRUD-API/action"
	"github.com/darkhunterLearning/Golang-CRUD-API/api"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func Init_Routing() {
	router := mux.NewRouter()

	// Route handles & endpoints
	router.HandleFunc("/", homePage)

	// Get all customers
	router.HandleFunc("/customers/", api.GetCustomers).Methods("GET")

	// Get a specific customer by the ID
	router.HandleFunc("/customer/{id}", api.GetCustomer).Methods("GET")

	// Create a customer
	router.HandleFunc("/customer/", api.CreateCustomer).Methods("POST")

	// Update a specific customer by the ID
	router.HandleFunc("/customers/{id}", api.UpdateCustomer).Methods("PATCH")

	// Delete a specific customer by the ID
	router.HandleFunc("/customers/{id}", api.DeleteCustomer).Methods("DELETE")

	// Delete all customers
	router.HandleFunc("/customers/", api.DeleteCustomers).Methods("DELETE")

	action.Server_Start()

	log.Fatal(http.ListenAndServe(":8000", router))
}
