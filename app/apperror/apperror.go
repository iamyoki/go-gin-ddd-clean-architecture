package apperror

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	BadRequest      struct{ Msg string }
	Unauthorized    struct{ Msg string }
	Forbidden       struct{ Msg string }
	NotFound        struct{ Msg string }
	Conflict        struct{ Msg string }
	RequestTooLarge struct{ Msg string }
	// Internal       struct{ Msg string }
	Validation struct {
		Msg     string
		Details map[string]string
	}
)

func (err *BadRequest) Error() string      { return err.Msg }
func (err *Unauthorized) Error() string    { return err.Msg }
func (err *Forbidden) Error() string       { return err.Msg }
func (err *NotFound) Error() string        { return err.Msg }
func (err *Conflict) Error() string        { return err.Msg }
func (err *RequestTooLarge) Error() string { return err.Msg }

// func (err *Internal) Error() string       { return "Internal Server Error" }
func (err *Validation) Error() string { return err.Msg }

func IntoResponse(err error) (int, gin.H) {
	// default
	status := http.StatusInternalServerError
	error := "Internal"
	msg := err.Error()

	// turn error into apperror
	switch err.(type) {
	case validator.ValidationErrors:
		err = &Validation{Msg: "Invalid request parameters"}
		msg = err.Error()
	case *http.MaxBytesError:
		err = &RequestTooLarge{Msg: "Request too large"}
		msg = err.Error()
	}

	// mapping
	switch e := err.(type) {
	case *BadRequest:
		status = http.StatusBadRequest
		error = "BadRequest"
	case *Unauthorized:
		status = http.StatusUnauthorized
		error = "Unauthorized"
	case *Forbidden:
		status = http.StatusForbidden
		error = "Forbidden"
	case *NotFound:
		status = http.StatusNotFound
		error = "NotFound"
	case *Conflict:
		status = http.StatusConflict
		error = "Conflict"
	case *RequestTooLarge:
		status = http.StatusRequestEntityTooLarge
		error = "RequestEntityTooLarge"
	case *Validation:
		status = http.StatusBadRequest
		error = "Validation"
	default:
		msg = "Internal server error" // hidden err reason for client
		slog.Error(e.Error())         // log err reason for server
	}

	// construct
	body := gin.H{
		"error": error,
		"msg":   msg,
	}

	return status, body
}
