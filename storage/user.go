package storage

import (
	"context"

	"github.com/felipedavid/saldop/models"
)

func InsertUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (
			name,
			email,
			password,
			phone_number,
			birth_date,
			job_title,
			company_name,
			document,
			document_type
		) VALUES (
		 	$1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING
		 	id,
			name,
			email,
			password,
			phone_number,
			birth_date,
			job_title,
			company_name,
			document,
			document_type,
			created_at,
			updated_at
	`

	err := conn.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.PhoneNumber,
		user.BirthDate,
		user.JobTitle,
		user.CompanyName,
		user.Document,
		user.DocumentType,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.BirthDate,
		&user.JobTitle,
		&user.CompanyName,
		&user.Document,
		&user.DocumentType,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(ctx context.Context, userID int) error {
	return nil
}

func GetUser(ctx context.Context, userID int) (*models.User, error) {
	return nil, nil
}
