package main

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/nipeharefa/lemonilo/controller"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// lemonilo

type (
	Application interface {
		StartHTTPServer()
	}

	application struct {
		e  *echo.Echo
		db *sqlx.DB
	}
)

func NewApplication() Application {

	app := &application{}

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())

	app.e = e

	return app
}

func (a *application) connectDB() {

	db, err := sqlx.Connect("postgres", viper.GetString("application.db.url"))
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	maxidle := viper.GetInt("application.db.max_idle")
	maxConn := viper.GetInt("application.db.max_conn")

	log.Info("Database connected")
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(maxidle)
	db.SetMaxOpenConns(maxConn)

	a.db = db

}

func (a *application) StartHTTPServer() {

	loginController := controller.NewLoginController()

	a.e.POST("/login", loginController.Login)
	log.Fatal(a.e.Start(":8000"))
}
