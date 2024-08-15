package db

import (
	"database/sql"
	"fmt"
	"log"
	"optimal_vi/tg_anki/pkg/conf"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	err  error
)

func GetDB() *sql.DB {
	if db == nil || db.Stats().OpenConnections == 0 {
		db, err = sql.Open("sqlite3", conf.DBPath)
		if err != nil {
			log.Fatal(err) // Handle the error appropriately in your application
		}
	}
	return db
}

func Migrate() {
	log.Println("Applying migrations...")

	GetDB()
	defer db.Close()

	existsMigrations := make(map[string]bool)

	files, err := os.ReadDir(conf.DBMigrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("Migration: %s", file.Name())
		if existsMigrations[file.Name()] {
			fmt.Println(" - Already applied")
			continue
		}
		sqlStmt, err := os.ReadFile(fmt.Sprintf("%s/%s", conf.DBMigrationsPath, file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.Exec(string(sqlStmt))
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		smtp, err := tx.Prepare("INSERT INTO migrations (name) VALUES (?)")
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		defer smtp.Close()
		_, err = smtp.Exec(file.Name())
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		tx.Commit()
		fmt.Println("- Applied")
	}

	log.Print("Migrations applied successfully")
}

func GetExistsMigration(name string) map[string]bool {
	existsMigrations := make(map[string]bool)
	GetDB()
	var count int
	var rows *sql.Rows

	rows, err = db.Query("SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name=?;", "migrations")
	if err != nil {
		log.Fatal(err)
	}
	rows.Scan(&count)
	rows.Close()
	if count == 0 {
		return existsMigrations
	}

	rows, err = db.Query("SELECT name FROM migrations")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		existsMigrations[name] = true
	}

	rows.Close()
	return existsMigrations
}
