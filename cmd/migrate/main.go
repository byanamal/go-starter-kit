package main

import (
	"base-api/internal/pkg/config"
	"base-api/internal/pkg/db"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: migrate <command> [args]")
		return
	}

	_ = godotenv.Load()

	dbConfig := config.LoadDatabase()
	db := db.Init(dbConfig.DSN())

	cmd := os.Args[1]

	if err := os.MkdirAll("migrations/", 0755); err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := os.MkdirAll("migrations/up", 0755); err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := os.MkdirAll("migrations/down", 0755); err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch cmd {

	case "create":
		if len(os.Args) < 3 {
			log.Fatal("Usage: migrate create <name>")
		}

		name := strings.ToLower(strings.ReplaceAll(os.Args[2], " ", "_"))

		ts := time.Now().Format("20060102150405")

		up := fmt.Sprintf("migrations/up/%s_%s.up.sql", ts, name)
		down := fmt.Sprintf("migrations/down/%s_%s.down.sql", ts, name)
		os.WriteFile(up, []byte("-- +migrate Up\n"), 0644)
		os.WriteFile(down, []byte("-- +migrate Down\n"), 0644)

		fmt.Println("Created:")
		fmt.Println("-", up)
		fmt.Println("-", down)

	case "up":
		files, err := os.ReadDir("migrations/up/")
		if err != nil {
			log.Fatal(err)
		}

		var ups []fs.DirEntry
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".up.sql") {
				ups = append(ups, f)
			}
		}

		sort.Slice(ups, func(i, j int) bool {
			return ups[i].Name() < ups[j].Name()
		})

		db.Exec(`
    	CREATE TABLE IF NOT EXISTS schema_migrations (
    	    version VARCHAR(14) PRIMARY KEY,
    	    name VARCHAR(255),
    	    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    	)`)
		for _, file := range ups {
			parts := strings.SplitN(file.Name(), "_", 2)
			version := parts[0]
			name := parts[1]

			var count int
			db.QueryRow(
				"SELECT COUNT(*) FROM schema_migrations WHERE version=$1",
				version,
			).Scan(&count)

			if count > 0 {
				continue
			}

			path := filepath.Join("migrations/up/", file.Name())
			sqlBytes, err := os.ReadFile(path)
			if err != nil {
				log.Fatalf("Failed to read migration file %s: %v", file.Name(), err)
			}
			sqlParts := strings.Split(string(sqlBytes), "-- +migrate Down")
			upSQL := strings.TrimPrefix(sqlParts[0], "-- +migrate Up")
			log.Printf("Executing migration %s:\n%s\n", file.Name(), upSQL)

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}

			if _, err := tx.Exec(upSQL); err != nil {
				tx.Rollback()
				log.Fatalf("Migration %s failed: %v", file.Name(), err)
			}
			raw := strings.TrimSuffix(name, ".up.sql")

			if _, err := tx.Exec(
				"INSERT INTO schema_migrations(version,name) VALUES($1,$2)",
				version, raw,
			); err != nil {
				tx.Rollback()
				log.Fatalf("Failed to log migration %s: %v", file.Name(), err)
			}
			tx.Commit()

			log.Println("Migrated:", file.Name())
		}

	case "down":
		var version, name string
		err := db.QueryRow(`
        SELECT version, name FROM schema_migrations
        ORDER BY version DESC
        LIMIT 1
    `).Scan(&version, &name)

		if err != nil {
			log.Fatal("No migration to rollback")
		}

		downFile := fmt.Sprintf("migrations/down/%s_%s.down.sql", version, name)

		sqlBytes, err := os.ReadFile(downFile)
		if err != nil {
			log.Fatalf("Down file not found: %s", downFile)
		}

		sqlParts := strings.Split(string(sqlBytes), "-- +migrate Down")
		downSQL := sqlParts[len(sqlParts)-1]

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		if _, err := tx.Exec(downSQL); err != nil {
			tx.Rollback()
			log.Fatalf("Rollback %s failed: %v", downFile, err)
		}

		tx.Exec("DELETE FROM schema_migrations WHERE version=$1", version)
		tx.Commit()

		fmt.Println("Rolled back:", downFile)

	case "status":

		files, err := os.ReadDir("migrations/up")
		if err != nil {
			log.Fatal(err)
		}

		applied := map[string]bool{}
		rows, _ := db.Query("SELECT version FROM schema_migrations")
		defer rows.Close()

		for rows.Next() {
			var v string
			rows.Scan(&v)
			applied[v] = true
		}

		fmt.Println("Migration status:")
		fmt.Println("VERSION\t\tNAME\t\t\tAPPLIED")

		sort.Slice(files, func(i, j int) bool {
			return files[i].Name() < files[j].Name()
		})

		for _, f := range files {
			parts := strings.SplitN(f.Name(), "_", 2)
			version := parts[0]
			name := strings.TrimSuffix(parts[1], ".up.sql")

			status := "❌"
			if applied[version] {
				status = "✅"
			}

			fmt.Printf("%s\t%s\t%s\n", version, name, status)
		}

	default:
		fmt.Println("Unknown command:", cmd)
	}
}
