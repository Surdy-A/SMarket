package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Vendor struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	LogoUrl     string    `json:"logo_url"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

func (v *Vendor) CreateVendor(db *sql.DB) error {
	v.CreatedDate = time.Now()
	vendor_id := (uuid.New()).String()
	v.ID = vendor_id
	v.UpdatedDate = v.CreatedDate

	err := db.QueryRow(`INSERT INTO vendors(id, name, email, phone, address, logo_url, created_date, updated_date) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		v.ID, v.Name, v.Email, v.Phone, v.Address, v.LogoUrl, v.CreatedDate, v.UpdatedDate).Scan(&v.ID)

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
		if err := rows.Scan(&v.ID, &v.Name, &v.Email, &v.Phone, &v.Address, &v.LogoUrl, &v.CreatedDate, &v.UpdatedDate); err != nil {
			return nil, err
		}
		vendors = append(vendors, v)
	}
	return vendors, nil
}

func (v *Vendor) GetVendor(db *sql.DB, id string) error {
	return db.QueryRow("SELECT * FROM Vendors WHERE id=$1", id).Scan(
	&v.ID, &v.Name, &v.Email, &v.Phone, &v.Address, &v.LogoUrl, &v.CreatedDate, &v.UpdatedDate)
}

func (v *Vendor) UpdateVendor(db *sql.DB) error {
	v.UpdatedDate = time.Now() 
	_, err :=
		db.Exec(`UPDATE vendors SET name=$1, email=$2, phone=$3, address=$4, logo_url=$5, updated_date=$6 WHERE id=$7`,
			v.Name, v.Email, v.Phone, v.Address, v.LogoUrl, v.UpdatedDate, v.ID)

	return err
}

func (v *Vendor) DeleteVendor(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM vendors WHERE id=$1", v.ID)

	return err
}
