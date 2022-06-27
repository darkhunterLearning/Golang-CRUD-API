package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "Customer"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	DB, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return DB
}

type Customer struct {
	CustomerID       int    `json:"id"`
	CustomerUniqueId string `json:"unique_id"`
	CustomerName     string `json:"customer_name"`
	CustomerPhone    string `json:"customer_phone"`
	CustomerAddress  string `json:"customer_address"`
	CustomerPassword string `json:"customer_password"`
}

type JsonResponse struct {
	Type    string     `json:"type"`
	Data    []Customer `json:"data"`
	Message string     `json:"message"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	router.HandleFunc("/", homePage)

	// Get all customers
	router.HandleFunc("/customers/", GetCustomers).Methods("GET")

	// Create a customer
	router.HandleFunc("/customer/", CreateCustomer).Methods("POST")

	// Update a specific customer by the ID
	router.HandleFunc("/customers/{id}", UpdateCustomer).Methods("PATCH")

	// Delete a specific customer by the ID
	router.HandleFunc("/customers/{id}", DeleteCustomer).Methods("DELETE")

	// Delete all customers
	router.HandleFunc("/customers/", DeleteCustomers).Methods("DELETE")

	// serve the app
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// response and request handlers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting customers...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM customer")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var customers []Customer

	// Foreach movie
	for rows.Next() {
		var (
			id               int
			uniqueID         string
			customerName     string
			customerPhone    string
			customerAddress  string
			customerPassword string
		)

		err = rows.Scan(&id, &uniqueID, &customerName, &customerPhone, &customerAddress, &customerPassword)

		// check errors
		checkErr(err)

		customers = append(customers, Customer{CustomerID: id, CustomerUniqueId: uniqueID, CustomerName: customerName, CustomerPhone: customerPhone,

			CustomerAddress: customerAddress, CustomerPassword: customerPassword})
	}

	var response = JsonResponse{Type: "success", Data: customers, Message: "200"}

	json.NewEncoder(w).Encode(response)

}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("id")
	Unique_Id := r.FormValue("unique_id")
	Customer_Name := r.FormValue("customer_name")
	Customer_Phone := r.FormValue("customer_phone")
	Customer_Address := r.FormValue("customer_address")
	Customer_Password := r.FormValue("customer_password")

	var response = JsonResponse{}

	if ID == "" || Customer_Name == "" {
		response = JsonResponse{Type: "error", Message: "You are missing ID or Customer_Name parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting customer into DB")

		fmt.Println("Inserting new customer with ID: " + ID + " and name: " + Customer_Name)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO customer(id, unique_id, customer_name, customer_phone, customer_address, customer_password) VALUES($1, $2, $3, $4, $5, $6) returning id;",
			ID, Unique_Id, Customer_Name, Customer_Phone, Customer_Address, Customer_Password).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "Customer has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	var response = JsonResponse{}

	if ID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing ID parameter."}
	} else {
		db := setupDB()
		Unique_Id := r.FormValue("unique_id")
		Customer_Name := r.FormValue("customer_name")
		Customer_Phone := r.FormValue("customer_phone")
		Customer_Address := r.FormValue("customer_address")
		Customer_Password := r.FormValue("customer_password")

		sqlUpdate := `
		UPDATE customer
		SET unique_id = $1, customer_name = $2, customer_phone = $3, customer_address = $4, customer_password = $5
		WHERE id = $6;`
		_, err := db.Exec(sqlUpdate, Unique_Id, Customer_Name, Customer_Phone, Customer_Address, Customer_Password, ID)

		checkErr(err)

		response = JsonResponse{Type: "success", Message: "Customer's Info has been updated successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	var response = JsonResponse{}

	if ID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing ID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting customer from DB")

		_, err := db.Exec("DELETE FROM customer where id = $1", ID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "Customer has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteCustomers(w http.ResponseWriter, r *http.Request) {
	db := setupDB()
	printMessage("Deleting all customer!")
	_, err := db.Exec("DELETE FROM customer")
	checkErr(err)
	printMessage("All customers have been deleted successfully")
	var response = JsonResponse{Type: "success", Message: "All customers have been deleted successfully"}
	json.NewEncoder(w).Encode(response)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Status code: 200 - Success"))
}
