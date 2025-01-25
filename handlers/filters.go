package handlers

import (
	"net/http"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/validator"
)

type Filters struct {
	Page     int
	PageSize int

	*validator.Validator
}

func newQueryFilters(r *http.Request) *Filters {
	values := r.URL.Query()
	v := validator.New(helpers.GetTranslator(r.Context()))

	return &Filters{
		Page:      getQueryInt(values, "page", 1, v),
		PageSize:  getQueryInt(values, "page_size", 20, v),
		Validator: v,
	}
}

func (f *Filters) Valid() bool {
	f.Check(f.Page > 0, "page", "should be greater than zero")
	f.Check(f.PageSize > 0, "page_size", "should be greater than zero")
	f.Check(f.PageSize <= 200, "page_size", "cannot be greater than 200")

	return len(f.Errors) == 0
}
