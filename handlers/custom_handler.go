package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/felipedavid/saldop/internal/helpers"
)

type Error interface {
	error
	Status() int
	AdditionalParams() map[string]any
}

type ErrorResponse struct {
	code             int
	message          string
	additionalParams map[string]any
}

func (e ErrorResponse) Error() string {
	return e.message
}

func (e ErrorResponse) Status() int {
	return e.code
}

func (e ErrorResponse) AdditionalParams() map[string]any {
	return e.additionalParams
}

func ErrorRes(status int, message string, additionalParams map[string]any) Error {
	return &ErrorResponse{code: status, message: message, additionalParams: additionalParams}
}

func ValidationError(paramsErrors map[string]string) Error {
	return &ErrorResponse{
		code:    http.StatusBadRequest,
		message: "failed validation",
		additionalParams: map[string]any{
			"param_errors": paramsErrors,
		}}
}

func QueryValidationError(paramsErrors map[string]string) Error {
	return &ErrorResponse{
		code:    http.StatusBadRequest,
		message: "invalid query params",
		additionalParams: map[string]any{
			"param_errors": paramsErrors,
		}}
}

func BadRequestError(ctx context.Context, message string) Error {
	return &ErrorResponse{code: http.StatusBadRequest, message: message}
}

func UnauthenticatedError(ctx context.Context) Error {
	return &ErrorResponse{code: http.StatusBadRequest, message: "user is not authenticated"}
}

type customHandler func(w http.ResponseWriter, r *http.Request) error

func handleErrors(h customHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := h(w, r)
		if err != nil {
			switch e := err.(type) {
			case Error:
				slog.Error(e.Error())

				t := helpers.GetTranslator(r.Context())

				resBody := map[string]any{
					"title":   http.StatusText(e.Status()),
					"status":  e.Status(),
					"message": t.Translate(e.Error()),
				}

				for k, v := range e.AdditionalParams() {
					resBody[k] = v
				}

				writeJSON(w, e.Status(), resBody)
			default:
				slog.Error("unexpected error", "error", err)
				debug.PrintStack()

				resBody := map[string]any{
					"title":          http.StatusText(http.StatusInternalServerError),
					"status":         http.StatusInternalServerError,
					"message":        "Something unexpected happen",
					"internal_error": err.Error(),
				}

				writeJSON(w, http.StatusInternalServerError, resBody)
			}
		}
	}
}
