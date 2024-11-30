package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUpdateAllTodosController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var Request UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&Request)

		_, err := db.Exec(
			"UPDATE todos SET title = ?, description = ? WHERE id = ?",
			Request.Title,
			Request.Description,
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
