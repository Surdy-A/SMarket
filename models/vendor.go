package models

import (
	"database/sql"
	"time"
)

type Vendor struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	LogoUrl     string    `json:"logo_url"`
	CreatedDate time.Time `json:"created_date"`
}

func (v *Vendor) CreateVendor(db *sql.DB) error {
	v.CreatedDate = time.Now()

	err := db.QueryRow(`INSERT INTO vendors(name, email, phone, address, logo_url, created_date) 
		VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		v.Name, v.Email, v.Phone, v.Address, v.LogoUrl, v.CreatedDate).Scan(&v.ID)

	if err != nil {
		return err
	}
	return nil
}

func (v *Vendor) GetVendors(db *sql.DB) ([]Vendor, error) {
	rows, err := db.Query(`SELECT * FROM vendors`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	vendors := []Vendor{}

	for rows.Next() {
		var v Vendor
		if err := rows.Scan(&v.ID, &v.Name, &v.Email, &v.Phone, &v.Address, &v.LogoUrl, &v.CreatedDate); err != nil {
			return nil, err
		}
		vendors = append(vendors, v)
	}
	return vendors, nil
}

func (v *Vendor) GetVendor(db *sql.DB, id int) error {
	return db.QueryRow("SELECT * FROM Vendors WHERE id=$1",
		id).Scan(&v.ID, &v.Name, &v.Email, &v.Phone, &v.Address, &v.LogoUrl, &v.CreatedDate)
}

func (v *Vendor) UpdateVendor(db *sql.DB) error {
	_, err :=
		db.Exec(`UPDATE vendors SET name=$1, email=$2, phone=$3, address=$4, logo_url=$5 WHERE id=$6`,
			v.Name, v.Email, v.Phone, v.Address, v.LogoUrl, v.ID)

	return err
}

func (v *Vendor) DeleteVendor(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM vendors WHERE id=$1", v.ID)

	return err
}
