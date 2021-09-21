package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/febrielven/go-crudapi/db"
)

func main() {
	// addr := ":8081"
	// listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	log.Fatalf("Error occurred: %s", err.Error())
	// }
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := db.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	fmt.Println("Starting server connection the port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))

}
