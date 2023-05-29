package main

import (
	"zpeTest/src/database"
	"zpeTest/src/server"
)

// @title ZPE Crud API
// @version 1.0
// @description ZPE User CRUD API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	database.DbConnect()
	server.HandleRequests()
}
