package Storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:123tres@localhost:5432/GoPruebas?sslmode=disable")
		if err != nil {
			log.Fatalf("Connection failed db: %v", err)
		}
		defer db.Close()
		if err = db.Ping(); err != nil {
			log.Fatalf("Connection failed, no Ping db: %v", err)
		}
	})

	fmt.Println("Conectado a Postgres")
}

// Pool retorna una única instancia de DB
func Pool() *sql.DB {
	return db
}
