package main

import (
	"gogin/database"
	"gogin/routes"
)

func main() {
	database.ConectaBanco()
	routes.HandleRequests()
}
