package main

import database "jwt-authentication-golang/Database"

func main() {
	// Initialize Database
	database.Connect("root:root@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()
}
