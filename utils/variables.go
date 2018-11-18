package utils

import (
	"fmt"
	"os"
)

var (
	VERSION string = "1.0"
	AUTHOR  string = ""
	TAG     string = ""
	PORT    string = "8080"
	HOST    string = ""
	DB      string

	DATABASE_HOST     string = "localhost"
	POSTGRES_USER     string = "jorge"
	POSTGRES_DB       string = "test"
	POSTGRES_PASSWORD string = "1234abc"
	POSTGRES_PORT     string = "5432"
)

func init() {
	if db := os.Getenv("DATABASE_HOST"); db != "" {
		DATABASE_HOST = db
	}

	if db := os.Getenv("POSTGRES_USER"); db != "" {
		POSTGRES_USER = db
	}

	if db := os.Getenv("POSTGRES_DB"); db != "" {
		POSTGRES_DB = db
	}

	if db := os.Getenv("POSTGRES_PASSWORD"); db != "" {
		POSTGRES_PASSWORD = db
	}

	if db := os.Getenv("POSTGRES_PORT"); db != "" {
		POSTGRES_PORT = db
	}

	DB = fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		DATABASE_HOST,
		POSTGRES_PORT,
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		POSTGRES_DB,
	)
}
