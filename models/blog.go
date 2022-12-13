package models

import (
	"database/sql"
	"time"
)

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Article     string    `json:"article"`
	Image       string    `json:"image"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	//Comment     []Comment   `json:"comment"`
}

func (b *Article) CreateArticle(db *sql.DB) error {
	b.CreatedDate = time.Now()
	b.UpdatedDate = time.Now()

	err := db.QueryRow(
		`INSERT INTO articles(title, article, image, created_date, updated_date) 
		VALUES($1, $2, $3, $4, $5) RETURNING id`,
		b.Title, b.Article, b.Image, b.CreatedDate, b.UpdatedDate).Scan(&b.ID)

	if err != nil {
		return err
	}
	return nil
}

func (b *Article) GetArticles(db *sql.DB) ([]Article, error) {
	rows, err := db.Query(
		`SELECT * FROM articles`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	blogs := []Article{}

	for rows.Next() {
		var b Article
		if err := rows.Scan(&b.ID, &b.Title, &b.Article, &b.Image, &b.CreatedDate, &b.UpdatedDate); err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}
	return blogs, nil
}
