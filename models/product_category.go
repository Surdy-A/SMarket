package models

import (
	"database/sql"
)

// type Category struct {
// 	ID           int    `json:"id"`
// 	MainCategory string `json:"main_category"`
// 	SubCategory  string `json:"sub_category"`
// }

func (c *Category) CreateCategory(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO categories(main_category, sub_category) VALUES($1, $2) RETURNING id",
		c.MainCategory, c.SubCategory).Scan(&c.ID)

	if err != nil {
		return err
	}

	return nil
}

func (c *Category) GetCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query(
		`SELECT id, main_category, sub_category FROM categories`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.MainCategory, &c.SubCategory); err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}

	return categories, nil
}

func (c *Category) GetCategory(db *sql.DB, id int) error {
	return db.QueryRow("SELECT * FROM categories WHERE id=$1",
		id).Scan(&c.ID, &c.MainCategory, &c.SubCategory)
}

func (c *Category) UpdateCategory(db *sql.DB) error {
	_, err :=
		db.Exec(`UPDATE categories SET main_category=$1, sub_category=$2 WHERE id=$3`,
			c.MainCategory, c.SubCategory, c.ID)

	return err
}

func (c *Category) DeleteCategory(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM categories WHERE id=$1", c.ID)

	return err
}
