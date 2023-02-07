package models

import (
	"database/sql"

	"github.com/google/uuid"
)

func (b *BlogCategory) CreateArticleCategory(db *sql.DB) error {
	category_id := (uuid.New()).String()
	b.ID = category_id

	err := db.QueryRow(
		"INSERT INTO article_categories(id, article_category) VALUES($1, $2) RETURNING id",
		category_id, b.ArticleCategory).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}

func (b *BlogCategory) GetArticleCategories(db *sql.DB) ([]BlogCategory, error) {
	rows, err := db.Query(
		`SELECT id, article_category FROM article_categories`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []BlogCategory{}

	for rows.Next() {
		var b BlogCategory

		if err := rows.Scan(&b.ID, &b.ArticleCategory); err != nil {
			return nil, err
		}

		categories = append(categories, b)
	}

	return categories, nil
}

func (b *BlogCategory) GetArticleCategory(db *sql.DB, id string) error {
	return db.QueryRow("SELECT * FROM article_categories WHERE id=$1", id).Scan(&b.ID, &b.ArticleCategory)
}

func (b *BlogCategory) UpdateArticleCategory(db *sql.DB) error {
	_, err := db.Exec(`UPDATE article_categories SET article_category=$1 WHERE id=$2`, b.ArticleCategory, b.ID)

	return err
}

func (b *BlogCategory) DeleteArticleCategory(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM article_categories WHERE id=$1", b.ID)

	return err
}
