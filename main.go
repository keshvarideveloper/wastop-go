package main

import (
	"github.com/keshvarideveloper/wastop/adapter/store"
	v1 "github.com/keshvarideveloper/wastop/delivery/http/v1"
	"github.com/keshvarideveloper/wastop/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	dsn := "adamak_user:adamak_pass@tcp(127.0.0.1:3306)/adamak?charset=utf8mb4&parseTime=True&loc=Local"
	// connect to database and auto migrate
	mysqlStore := store.New(dsn)

	// setup http server and router
	e := echo.New()

	// add routes

	e.POST("/users", v1.SignUpUser(mysqlStore,
		validator.ValidateCreateUser))

	e.Logger.Fatal(e.Start(":8080"))
}
