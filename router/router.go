package router

import (
	"database/sql"
	"gcp/handler"

	"github.com/gorilla/mux"
)

// Setup : Router init
func Setup(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	ph := handler.NewProductHandler(db)
	r.HandleFunc("/product", ph.PostProductHandler).Methods("POST")
	r.HandleFunc("/product", ph.GetProductHandler).Methods("GET")
	return r
}
