package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
}

// Define product here
type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"product_name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (a *App) Initialize() (*sql.DB, error) {
	conStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", DbUser, DbPass, DbHost, DbPort, DbName)

	db, err := sql.Open("postgres", conStr)

	return db, err
}

func (a *App) getAllProducts(w http.ResponseWriter, req *http.Request) {
	con, err := a.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	defer con.Close()

	var products []Product
	rows, err := con.Query("SELECT id, product_name, quantity, price from products")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}

	data, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (a *App) getProductDetails(w http.ResponseWriter, req *http.Request) {
	conn, err := a.Initialize()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id := mux.Vars(req)["id"]

	var product Product
	query := fmt.Sprintf("SELECT id, product_name, quantity, price FROM products WHERE id=%v", id)
	row, err := conn.Query(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for row.Next() {
		err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)

}

func AddProduct(w http.ResponseWriter, req *http.Request) {}

func main() {
	app := App{}
	router := mux.NewRouter().StrictSlash(true) // FIXME: what does the strict slash does ?

	app.Initialize()
	router.HandleFunc("/products", app.getAllProducts)
	router.HandleFunc("/product/{id}", app.getProductDetails)
	http.ListenAndServe(":8000", router)
}
