package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./bot.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, qq from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var qq string
		err = rows.Scan(&id, &qq)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, qq)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
