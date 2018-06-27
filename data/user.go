package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id  int64
	Qq  string
	Tag string
	db  *sql.Db
}

func NewUser(qq string, tag string) *User {

	db, err := sql.Open("sqlite3", "./bot.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &User{Qq: qq, Tag: tag, db: db}
}

func CreateUser(user *User) *User {
	return user
}
