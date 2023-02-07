package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID              string       `json:"id"`
	Title           string       `json:"title"`
	Article         string       `json:"article"`
	Image           string       `json:"image"`
	CreatedDate     time.Time    `json:"created_date"`
	UpdatedDate     time.Time    `json:"updated_date"`
	ArticleCategory BlogCategory `json:"article_category"`
	//Comment     []Comment   `json:"comment"`  Add Comment
}

type BlogCategory struct {
	ID              string `json:"id"`
	ArticleCategory string `json:"category"`
}

func (b *Article) CreateArticle(db *sql.DB) error {
	b.CreatedDate = time.Now()
	b.UpdatedDate = time.Now()
	article_id := (uuid.New()).String()
	b.ID = article_id

	err := db.QueryRow(
		`INSERT INTO articles(id, title, article, image, created_date, updated_date, categories) 
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		b.ID, b.Title, b.Article, b.Image, b.CreatedDate, b.UpdatedDate, b.ArticleCategory).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}

func (p *Article) GetArticles(db *sql.DB) ([]Article, error) {
	rows, err := db.Query(`SELECT id, title, article, image, created_date, updated_date FROM articles`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Article{}

	for rows.Next() {
		var p Article
		if err := rows.Scan(&p.ID, &p.Title, &p.Article, &p.Image, &p.CreatedDate, &p.UpdatedDate); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (b *Article) GetArticle(db *sql.DB, id string) error {
	b.UpdatedDate = time.Now()

	return db.QueryRow("SELECT * FROM Articles WHERE id=$1", id).Scan(&b.ID, &b.Title, &b.Article,
		&b.Image, &b.CreatedDate, &b.UpdatedDate, &b.ArticleCategory)
}

func (b *Article) UpdateArticle(db *sql.DB) error {
	_, err := db.Exec(`UPDATE articles SET title=$1, article=$2, image=$3, created_date=$4, 
	updated_date=$5 WHERE id=$6`, &b.Title, &b.Article, &b.Image, &b.CreatedDate, &b.UpdatedDate, &b.ID)

	return err
}

func (p *Article) DeleteArticle(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM articles WHERE id=$1", p.ID)

	return err
}

func (bc BlogCategory) Value() (driver.Value, error) {
	return json.Marshal(bc)
}

func (a *BlogCategory) Scan(value interface{}) error {
	b, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
