package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Flame-eth/bookstore-mng/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	var PORT = "8080"

	r := mux.NewRouter()

	r = routes.SetBookstoreRoutes(r)

	http.Handle("/", r)

	fmt.Println("Server is running on port:", PORT)
	err := http.ListenAndServe(":"+PORT, r)
	if err != nil {
		log.Fatal(err)
	}

}
