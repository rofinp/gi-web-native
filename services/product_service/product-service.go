package product_service

import (
	"fmt"

	"github.com/rofinp/go-web-native/configs"
	"github.com/rofinp/go-web-native/models"
)

func GetAllProducts() ([]models.Product, error) {
	rows, err := configs.DB.Query(`
		SELECT 
			products.id,
			products.name,
			categories.name as category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products
		JOIN categories ON products.category_id = categories.id
	`)

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	products := []models.Product{}

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		products = append(products, product)
	}

	return products, nil
}

func CreateProduct(p models.Product) error {
	err := configs.DB.QueryRow(`
		INSERT INTO products (name, category_id, stock, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		p.Name, p.Category.ID, p.Stock, p.Description, p.CreatedAt, p.UpdatedAt,
	).Scan(&p.ID)

	if err != nil {
		return fmt.Errorf("failed to create product: %v", err)
	}

	if p.ID < 1 {
		return fmt.Errorf("failed to create product: ID is not valid (%d)", p.ID)
	}

	return nil
}

func DetailProduct(id int) (*models.Product, error) {
	row := configs.DB.QueryRow(`
		SELECT 
			products.id,
			products.name,
			categories.name as category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products
		JOIN categories ON products.category_id = categories.id
		WHERE products.id = $1`,
		id,
	)

	product := models.Product{}

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Category.Name,
		&product.Stock,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan product: %v", err)
	}
	return &product, nil
}

func UpdateProduct(id int, p models.Product) error {
	result, err := configs.DB.Exec(`
	UPDATE products 
	SET name = $1, category_id = $2, stock = $3, description = $4, updated_at = $5 
	WHERE id = $6`,
		p.Name, p.Category.ID, p.Stock, p.Description, p.UpdatedAt, id,
	)

	if err != nil {
		return fmt.Errorf("failed to update products: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to update products: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("failed to update product: product with ID %d does not exist", id)
	}

	return nil
}

func DeleteProduct(id int) error {
	result, err := configs.DB.Exec(`DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("failed to delete product: product with ID %d does not exist", id)
	}

	return nil
}
