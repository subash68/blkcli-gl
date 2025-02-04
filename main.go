package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subash68/blkcli-gl/pkg/product"
)

func main() {
	app := product.App{}
	router := mux.NewRouter().StrictSlash(true) // FIXME: what does the strict slash does ?

	app.Initialize()
	router.HandleFunc("/products", app.GetAllProducts)
	router.HandleFunc("/product/{id}", app.GetProductDetails)
	http.ListenAndServe(":8000", router)
}
