package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Define product here
type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var products []Product

func getAllProducts(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting All products")
	json.NewEncoder(w).Encode(products)
}

func getProductDetails(w http.ResponseWriter, req *http.Request) {
	queryParams := mux.Vars(req)
	id := queryParams["id"]

	for _, item := range products {
		itemId, _ := strconv.Atoi(id)
		if item.Id == itemId {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func AddProduct(w http.ResponseWriter, req *http.Request) {}

func main() {

	products = []Product{
		Product{Id: 1, Name: "Product1", Quantity: 100, Price: 10.34},
		Product{Id: 2, Name: "Product2", Quantity: 1000, Price: 1.55},
	}

	// Define router here
	router := mux.NewRouter().StrictSlash(true) // FIXME: what does the strict slash does ?

	router.HandleFunc("/products", getAllProducts)
	router.HandleFunc("/product/{id}", getProductDetails)
	http.ListenAndServe(":8000", router)
}
