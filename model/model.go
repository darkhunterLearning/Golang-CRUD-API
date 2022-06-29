package model

type Customer struct {
	CustomerID       int    `json:"id"`
	CustomerUniqueId string `json:"unique_id"`
	CustomerName     string `json:"customer_name"`
	CustomerPhone    string `json:"customer_phone"`
	CustomerAddress  string `json:"customer_address"`
	CustomerEmail    string `json:"customer_email"`
	CustomerPassword string `json:"customer_password"`
}

type JsonResponse struct {
	Type    string     `json:"type"`
	Data    []Customer `json:"data"`
	Message string     `json:"message"`
}
