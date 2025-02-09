package storage

import (
	"context"

	"github.com/felipedavid/saldop/internal/filters"
	"github.com/felipedavid/saldop/models"
)

func InsertCategory(ctx context.Context, c *models.Category) error {
	query := `
        INSERT INTO category (name, user_id)
        VALUES ($1, $2)
        RETURNING id, created_at, updated_at
    `

	err := conn.QueryRow(
		ctx,
		query,
		c.Name,
		c.UserID,
	).Scan(
		&c.ID,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCategory(ctx context.Context, category *models.Category) error {
	return nil
}

func ListUserCategories(ctx context.Context, userID int, f *filters.Filters) ([]models.Category, error) {
	query := `
		SELECT
            count(*) OVER(),
			id,
            name,
            default_category,
            user_id,
			created_at,
			updated_at,
			deleted_at
		FROM category
		WHERE (default_category IS true OR user_id = $1) AND deleted_at IS NULL
        LIMIT $2 OFFSET $3
	`

	rows, err := conn.Query(ctx, query, userID, f.Limit(), f.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalRecords := 0
	categories := make([]models.Category, 0, f.PageSize)
	for rows.Next() {
		var c models.Category
		err := rows.Scan(
			&totalRecords,
			&c.ID,
			&c.Name,
			&c.DefaultCategory,
			&c.UserID,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}

	f.SetTotalRecords(totalRecords)

	return categories, nil
}

func DeleteCategory(ctx context.Context, userID, categoryID int) error {
	query := `
		UPDATE category
		SET deleted_at = NOW()
		WHERE user_id = $1 AND id = $2
	`

	_, err := conn.Exec(ctx, query, userID, categoryID)
	if err != nil {
		return err
	}

	return nil
}
