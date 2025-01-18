package storage

import (
	"context"
	"errors"

	"github.com/felipedavid/saldop/models"
	"github.com/jackc/pgx/v5/pgconn"
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
		) RETURNING id, created_at, updated_at
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
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicatedEmail
		}
		return err
	}

	return nil
}

func DeleteUser(ctx context.Context, userID int) error {
	query := `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1`
	_, err := conn.Exec(ctx, query, userID)
	return err
}

func GetUser(ctx context.Context, userID int) (*models.User, error) {
	query := `
		SELECT
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
			updated_at,
			deleted_at
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := conn.QueryRow(ctx, query, userID).Scan(
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
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT
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
			updated_at,
			deleted_at
		FROM users
		WHERE email = $1
	`

	var user models.User
	err := conn.QueryRow(ctx, query, email).Scan(
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
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
