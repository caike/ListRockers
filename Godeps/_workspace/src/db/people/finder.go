package people

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func FindAll() []string {
	db, err := sql.Open("postgres", "dbname=go_people sslmode=disable")
	PanicIf(err)

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
