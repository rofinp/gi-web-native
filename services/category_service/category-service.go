package category_service

import (
	"fmt"

	"github.com/rofinp/go-web-native/configs"
	"github.com/rofinp/go-web-native/models"
)

func GetAllCategories() ([]models.Category, error) {
	rows, err := configs.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	categories := []models.Category{}

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func CreateCategory(c models.Category) error {
	err := configs.DB.QueryRow(`
		INSERT INTO categories (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id`,
		c.Name, c.CreatedAt, c.UpdatedAt,
	).Scan(&c.ID)

	if err != nil {
		return fmt.Errorf("failed to create category: %v", err)
	}

	if c.ID < 1 {
		return fmt.Errorf("failed to create category: ID is not valid (%d)", c.ID)
	}

	return nil
}

func DetailsCategory(id int) (*models.Category, error) {
	// Execute SQL query to select the id and name of the category from the categories table
	row := configs.DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", id)

	// Create an empty Category struct to store the retrieved category details
	category := models.Category{}

	// Populate the category struct with the values from the query result
	if err := row.Scan(&category.ID, &category.Name); err != nil {
		// Return an error if there was a problem scanning the result
		return nil, fmt.Errorf("failed to scan category: %v", err)
	}

	// Return the populated category struct
	return &category, nil
}

func UpdateCategory(id int, c models.Category) error {
	// Execute the SQL query to update the category with the given ID
	result, err := configs.DB.Exec("UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3", c.Name, c.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("failed to update category: %v", err)
	}

	// Get the number of affected rows from the query result
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to update category: %v", err)
	}

	// If no rows were affected, return an error indicating that the category with the given ID does not exist
	if rowsAffected == 0 {
		return fmt.Errorf("failed to update category: category with ID %d does not exist", id)
	}

	return nil
}

func DeleteCategory(id int) error {
	// Execute the SQL query to delete the category with the given ID
	result, err := configs.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}

	// Get the number of affected rows from the query result
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}

	// If no rows were affected, return an error indicating that the category with the given ID does not exist
	if rowsAffected == 0 {
		return fmt.Errorf("failed to delete category: category with ID %d does not exist", id)
	}

	return nil
}
