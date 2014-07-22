package people

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func FindAll() []string {

	conninfo := os.Getenv("HEROKU_POSTGRESQL_SILVER_URL")
	db, err := sql.Open("postgres", conninfo)
	PanicIf(err)
	defer db.Close()

	rows, err := db.Query("SELECT name FROM people;")
	PanicIf(err)
	defer rows.Close()

	var name string
	var names []string

	for rows.Next() {
		err := rows.Scan(&name)
		PanicIf(err)
		names = append(names, name)
	}
	return names
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
