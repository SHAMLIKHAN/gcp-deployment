package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func preparePostgres() (*sql.DB, error) {
	url, err := getDatabaseURL()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDatabaseURL() (string, error) {
	env, err := getEnv()
	if err != nil {
		return "", err
	}
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_USER"],
		env["DB_PASSWORD"],
		env["DB_NAME"],
	)
	return psql, nil
}

func getEnv() (map[string]string, error) {
	env := make(map[string]string)
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, errors.New("DB_HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, errors.New("DB_PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return nil, errors.New("DB_USER environment variable required but not set")
	}
	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, errors.New("DB_PASSWORD environment variable required but not set")
	}
	database, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, errors.New("DB_NAME environment variable required but not set")
	}
	env["DB_HOST"] = host
	env["DB_PORT"] = port
	env["DB_USER"] = user
	env["DB_PASSWORD"] = password
	env["DB_NAME"] = database
	return env, nil
}
