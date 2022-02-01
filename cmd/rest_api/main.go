package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	router "github.com/jsokka/todo-app-go"
	"github.com/jsokka/todo-app-go/db"
)

func main() {
	err := godotenv.Load("../../.env.local")

	if err != nil {
		log.Print("Failed to load .env.local")
	}

	db.Init(os.Getenv("TODOAPPGO_CONNECTION_STRING"))

	addr, hasAddress := os.LookupEnv("TODOAPPGO_ADDR")

	if !hasAddress {
		addr = ":8080"
	}

	router.InitRouter().Run(addr)
}
