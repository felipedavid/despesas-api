package models

import "time"

type Category struct {
	ID              int    `db:"id" json:"id"`
	Name            string `db:"name" json:"name"`
	DefaultCategory string `db:"default_category" json:"default_category"`
	UserID          *int   `db:"user_id" json:"user_id"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
