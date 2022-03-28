package main

import (
	"hacktiv8_assignment_3/config"
	"hacktiv8_assignment_3/routes"
)

func main() {
	PORT := ":8080"
	config.StartDB()
	routes.Serve().Run(PORT)
}
