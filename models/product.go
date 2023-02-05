package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Image        string    `json:"image"`
	Details      string    `json:"details"`
	Sizes        []string  `json:"sizes"`
	Colours      []string  `json:"colours"`
	VideoUrl     string    `json:"video_url"`
	Availability bool      `json:"availability"`
	Star         []int64   `json:"star"`
	Labels       []string  `json:"labels"`
	Discount     float64   `json:"discount"`
	Brands       []string  `json:"brands"`
	Categories   Category  `json:"categories"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}

type Category struct {
	ID string `json:"id"`
	//Add Menu
	MainCategory string `json:"main_category"`
	SubCategory  string `json:"sub_category"` // Make subcategory ann array
}

func (p *Product) ProductDiscount(percent float64) *float64 {
	percent = percent / 100
	pprice := percent * p.Price
	discountPrice := p.Price - pprice
	return &discountPrice
}

func (p *Product) CreateProduct(db *sql.DB) error {
	p.CreatedDate = time.Now()
	product_id := (uuid.New()).String()
	p.ID = product_id
	p.UpdatedDate = p.CreatedDate

	err := db.QueryRow(
		"INSERT INTO products(id, name, image, details, sizes, colours, video_url, availability, star, labels, discount, price, brands, categories, created_date, updated_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id",
		p.ID, p.Name, p.Image, p.Details, pq.Array(p.Sizes), pq.Array(p.Colours), p.VideoUrl, p.Availability, pq.Array(p.Star), pq.Array(p.Labels), p.ProductDiscount(p.Discount), p.Price, pq.Array(p.Brands), p.Categories, p.CreatedDate, p.UpdatedDate).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) GetProducts(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query(
		`SELECT id, name, image, details, sizes, colours, video_url, availability, 
		star, labels, discount, price, brands, categories, created_date, updated_date FROM products LIMIT $1 OFFSET $2`,
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Image, &p.Details, pq.Array(&p.Sizes), pq.Array(&p.Colours), &p.VideoUrl,
			&p.Availability, pq.Array(&p.Star), pq.Array(&p.Labels), &p.Discount, &p.Price, pq.Array(&p.Brands), &p.Categories, &p.CreatedDate, &p.UpdatedDate); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (p *Product) GetProduct(db *sql.DB, id string) error {
	return db.QueryRow("SELECT * FROM products WHERE id=$1",
		id).Scan(&p.ID, &p.Name, &p.Image, &p.Details, pq.Array(&p.Sizes), pq.Array(&p.Colours), &p.VideoUrl,
		&p.Availability, pq.Array(&p.Star), pq.Array(&p.Labels), &p.Discount, &p.Price, pq.Array(&p.Brands),
		&p.Categories, &p.CreatedDate, &p.UpdatedDate)
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	p.UpdatedDate = time.Now()
	_, err :=
		db.Exec(`UPDATE products SET name=$1, image=$2, details=$3, sizes=$4, colours=$5,
		 	video_url=$6, availability=$7, star=$8, labels=$9, discount=$10, price=$11, brands=$12, categories=$13, updated_date=$14 WHERE id=$15`,
			p.Name, p.Image, p.Details, pq.Array(p.Sizes), pq.Array(p.Colours), p.VideoUrl, p.Availability,
			pq.Array(p.Star), pq.Array(p.Labels), p.ProductDiscount(p.Discount), p.Price, pq.Array(p.Brands), p.Categories, p.UpdatedDate, p.ID)

	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (a Category) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Category) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
