package main

import (
	"github.com/darkhunterLearning/Golang-CRUD-API/route"
	_ "github.com/lib/pq"
)

func main() {
	// Init the mux router
	route.Init_Routing()
}
