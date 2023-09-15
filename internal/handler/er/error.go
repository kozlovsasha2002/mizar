package er

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	HTTPStatusCode int    `json:"status_code"`
	UserMessage    string `json:"message"`
	Error          string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusNotFound,
		UserMessage:    "Data not found",
		Error:          err.Error(),
	}
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusBadRequest,
		UserMessage:    "Invalid request",
		Error:          err.Error(),
	}
}

func ErrInternal(err error) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		UserMessage:    "Error during creation",
		Error:          err.Error(),
	}
}

func ErrUnprocessableEntity(err error) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusUnprocessableEntity,
		UserMessage:    "Error rendering response",
		Error:          err.Error(),
	}
}
