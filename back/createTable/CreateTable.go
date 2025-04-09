package createTable

import "database/sql"

func CreateTable(db *sql.DB) {

	// query
	q := `CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err)
	}
}
