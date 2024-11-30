package main

import (
	"github.com/Dryluigi/golang-todos/controller"
	"github.com/Dryluigi/golang-todos/database"
	"github.com/labstack/echo"
)

func main() {
	db := database.InirDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	controller.NewGetAllTodosController(e, db)
	controller.NewCreateAllTodosController(e, db)
	controller.NewDeleteAllTodosController(e, db)
	controller.NewUpdateAllTodosController(e, db)
	controller.NewCheckAllTodosController(e, db)

	e.Start(":8080")
}
