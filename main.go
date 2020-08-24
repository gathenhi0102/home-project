package main

import (
	"fmt"
	"log"
	"net/http"

	auth "go_modules/auth"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := httprouter.New()

	auth.AuthRouter(router)

	fmt.Println("Listening to port 6699")
	log.Fatal(http.ListenAndServe(":6699", router))
}
