package handler

import (
	"database/sql"
	"encoding/json"
	"gcp/models"
	"gcp/repository"
	"gcp/utils"
	"net/http"
)

// ProductHandler :
type ProductHandler struct {
	DB *sql.DB
}

// NewProductHandler :
func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{
		DB: db,
	}
}

// PostProductHandler :
func (p *ProductHandler) PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.Fail(w, 400, 101, "decode error")
		return
	}
	response, err := repository.PostProduct(p.DB, &product)
	if err != nil {
		utils.Fail(w, 500, 102, err.Error())
		return
	}
	utils.Send(w, 200, response)
}

// GetProductHandler :
func (p *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	products, err := repository.GetProduct(p.DB)
	if err != nil {
		utils.Fail(w, 500, 102, err.Error())
		return
	}
	utils.Send(w, 200, products)
}
