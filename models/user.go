package models

import "time"

type User struct {
	ID           int        `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Email        string     `db:"email" json:"email"`
	Password     []byte     `db:"password" json:"-"`
	PhoneNumber  *string    `db:"phone_number" json:"phone_number"`
	BirthDate    *time.Time `db:"birth_date" json:"-"`
	JobTitle     *string    `db:"job_title" json:"-"`
	CompanyName  *string    `db:"company_name" json:"-"`
	Document     *string    `db:"document" json:"-"`
	DocumentType *string    `db:"document_type" json:"-"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}
