package main

import (
	"log"
	"net/http"
	"os"

	"github.com/heroku/TransitGet/golib"
)

//*****************************************************
func main() {
	log.Printf("Server started")
	// リクエストの待ち受け
	golib.NewRouter()
	log.Printf("PORT:", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
