package cmd

import (
	"database/sql"
	"gcp/router"
	"log"
	"net/http"
)

// App : Struct to represent this app
type App struct {
	DB *sql.DB
}

// NewApp : to get App Struct
func NewApp(db *sql.DB) *App {
	return &App{
		DB: db,
	}
}

// Serve : to Run API Server
func (a *App) Serve(addr string) {
	router := router.Setup(a.DB)
	log.Println("App : Server is listening...")
	http.ListenAndServe(addr, router)
}
