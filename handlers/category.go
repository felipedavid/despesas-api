package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/felipedavid/saldop/internal/filters"
	"github.com/felipedavid/saldop/internal/helpers"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func createCategory(w http.ResponseWriter, r *http.Request) error {
	params := service.NewCreateCategoryParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	if !params.Validate() {
		return ValidationError(params.Errors)
	}

	user := helpers.GetUserFromRequestContext(r)

	newCategory := &models.Category{
		Name:   *params.Name,
		UserID: &user.ID,
	}
	err = storage.InsertCategory(context.Background(), newCategory)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, newCategory)
}

func listUserCategories(w http.ResponseWriter, r *http.Request) error {
	user := helpers.GetUserFromRequestContext(r)

	filters := filters.NewQueryFilters(r)
	if !filters.Valid() {
		return QueryValidationError(filters.Errors)
	}

	categories, err := storage.ListUserCategories(context.Background(), user.ID, filters)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"metadata":   filters.Metadata(),
		"categories": categories,
	})
}

func deleteCategory(w http.ResponseWriter, r *http.Request) error {
	categoryID := r.PathValue("categoryID")

	user := helpers.GetUserFromRequestContext(r)

	err := storage.DeleteCategory(context.Background(), user.ID, categoryID)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func editCategory(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "not implemented yet")
	return nil
}
