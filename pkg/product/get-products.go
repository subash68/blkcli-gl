package product

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/subash68/blkcli-gl/pkg/config"
)

type App struct {
}

func (a *App) Initialize() (*sql.DB, error) {
	conStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)

	db, err := sql.Open("postgres", conStr)

	return db, err
}

func (a *App) GetAllProducts(w http.ResponseWriter, req *http.Request) {
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
