package models

import "time"

type Category struct {
	ID              string  `db:"id" json:"id"`
	Name            string  `db:"name" json:"name"`
	DefaultCategory bool    `db:"default_category" json:"default_category"`
	UserID          *string `db:"user_id" json:"user_id"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type CategoryNullable struct {
	ID              *string `db:"id"`
	Name            *string `db:"name"`
	DefaultCategory *bool   `db:"default_category"`
	UserID          *string `db:"user_id"`

	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
