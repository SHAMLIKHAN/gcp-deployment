package repository

import (
	"database/sql"
	"errors"
	"gcp/models"
	"log"
)

// PostProduct :
func PostProduct(db *sql.DB, product *models.Product) (*models.Product, error) {
	query := `
		INSERT INTO products (name, brand) VALUES ($1, $2) RETURNING id;
	`
	err := db.QueryRow(query, product.Name, product.Brand).Scan(&product.ID)
	if err != nil {
		log.Println(err, " in PostProduct repository")
		return nil, errors.New("error creating product")
	}
	return product, nil
}

// GetProduct :
func GetProduct(db *sql.DB) ([]models.Product, error) {
	var products []models.Product
	var product models.Product
	query := `
		SELECT * FROM products ORDER BY id DESC LIMIT 10;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err, " in GetProduct repository")
		return nil, errors.New("error finding product")
	}
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Brand)
		if err != nil {
			log.Println(err, " in GetProduct repository")
			return nil, errors.New("error scanning product")
		}
		products = append(products, product)
	}
	return products, nil
}
