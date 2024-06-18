package rest

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPErrorHandler is a custom error handler for the REST API
// Stores a map of errors and their respective HTTP status code
type httpErrorHandler struct {
	errors map[error]int
}

func NewHTTPErrorHandler(errors map[error]int) *httpErrorHandler {
	return &httpErrorHandler{
		errors: errors,
	}
}

func (h *httpErrorHandler) getStatusCodeByError(err error) int {
	for k, v := range h.errors {
		if errors.Is(err, k) {
			return v
		}
	}
	return http.StatusInternalServerError
}

func (h *httpErrorHandler) getErrorMessageByStatusCode(code int, err error) string {
	if code == http.StatusInternalServerError {
		return "Service unavailable"
	}
	return err.Error()
}

func (h *httpErrorHandler) CustomHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		// Non-echo HTTPError, can be a panic or a custom business error
		sc := h.getStatusCodeByError(err)
		he = &echo.HTTPError{
			Code:    sc,
			Message: h.getErrorMessageByStatusCode(sc, err),
		}
	}

	code := he.Code
	message := he.Message

	switch m := he.Message.(type) {
	case string:
		message = map[string]interface{}{"status": code, "message": m}
	case error:
		message = map[string]interface{}{"status": code, "message": m.Error()}
	}

	c.Logger().Error(err)

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
