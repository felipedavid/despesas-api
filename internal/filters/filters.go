package filters

import (
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/felipedavid/saldop/internal/helpers"
	"github.com/felipedavid/saldop/internal/validator"
)

type Metadata struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	FirstPage    int `json:"first_page"`
	LastPage     int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

type Filters struct {
	Page     int
	PageSize int

	totalRecords int
	*validator.Validator
}

// To generate correct metadata we need the number of records after the filtering.
// So when you get that data, probably in the storage layer, call SetTotalRecords
func (f *Filters) Metadata() *Metadata {
	m := &Metadata{}

	if f.totalRecords == 0 {
		return m
	}

	m.CurrentPage = f.Page
	m.PageSize = f.PageSize
	m.FirstPage = 1
	m.LastPage = int(math.Ceil(float64(f.totalRecords) / float64(f.PageSize)))
	m.TotalRecords = f.totalRecords

	return m
}

func (f *Filters) SetTotalRecords(totalRecords int) {
	f.totalRecords = totalRecords
}

func NewQueryFilters(r *http.Request) *Filters {
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

func (f *Filters) Limit() int {
	return f.PageSize
}

func (f *Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}

func getQueryInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}

	return i
}
