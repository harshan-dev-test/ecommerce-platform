package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"user-service/database"
	"user-service/router"
)

func main() {
	
	PORT := os.Getenv("APPLICATION_PORT")
	if PORT ==""{
		PORT="8080"
	}
	
	db := database.ConnectDB()
	defer db.Close()

	router := router.SetupRouter(db)

	fmt.Println("Server started at http://localhost:"+PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}