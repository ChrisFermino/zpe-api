package main

import (
    "zpeTest/src/database"
    "zpeTest/src/server"
)

func main() {
    database.DbConnect()
    server.HandleRequests()
}
