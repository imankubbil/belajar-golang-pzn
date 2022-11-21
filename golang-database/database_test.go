package golangdatabase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {
}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_belajar_golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// gunakan DB
}
