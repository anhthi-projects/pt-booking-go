package app

import (
	"anhthi-projects/pt-booking-go/app/middlewares"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func InitDatabase(connStr string) *bun.DB {
	config, err := pgx.ParseConfig(connStr)

	if err != nil {
		panic(err)
	}

	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	var sqlDB *sql.DB = stdlib.OpenDB(*config)

	db := bun.NewDB(sqlDB, pgdialect.New())
	
	return db
}

func InitServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middlewares.WriteToConsole)

	return e
}

