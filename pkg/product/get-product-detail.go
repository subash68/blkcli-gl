package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) GetProductDetails(w http.ResponseWriter, req *http.Request) {
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
