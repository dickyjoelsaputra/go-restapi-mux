package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vanjul123/go-restapi-mux/controller/productcontroller"
	"github.com/vanjul123/go-restapi-mux/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/products",productcontroller.Index).Methods("GET")
	r.HandleFunc("/products/{id}",productcontroller.Show).Methods("GET")
	r.HandleFunc("/products",productcontroller.Create).Methods("POST")
	r.HandleFunc("/products/{id}",productcontroller.Update).Methods("PUT")
	r.HandleFunc("/products",productcontroller.Delete).Methods("DELETE")

	fmt.Print("server running at http://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe(":8080",r))

}