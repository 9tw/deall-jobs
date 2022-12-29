package main

import (
	"deall/config"
	dUser "deall/features/user/delivery"
	rUser "deall/features/user/repository"
	sUser "deall/features/user/services"
	"deall/utils/database"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	validator := validator.New()

	mdlUser := rUser.New(db)
	serUser := sUser.New(mdlUser, validator)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	dUser.New(e, serUser)

	e.Logger.Fatal(e.Start(":8000"))

}
