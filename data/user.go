package data

import (
	"database/sql"
	// "fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	// "os"
)

type User struct {
	Id  int64
	Qq  string
	Tag string
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./assets/bot.sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateUser(user *User) error {
	db := openDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into users(qq, tag) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Qq, user.Tag)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func FindUserByTag(t string) *User {
	db := openDb()
	defer db.Close()

	stmt, err := db.Prepare("select id, qq, tag from users where tag = ?")
	if err != nil {
		return nil
	}
	defer stmt.Close()

	u := User{}

	err = stmt.QueryRow(t).Scan(&u.Id, &u.Qq, &u.Tag)
	if err != nil {
		return nil
	}
	return &u
}

func FindUserByQq(t string) *User {
	db := openDb()
	defer db.Close()

	stmt, err := db.Prepare("select id, qq, tag from users where qq = ?")
	if err != nil {
		return nil
	}
	defer stmt.Close()

	u := User{}

	err = stmt.QueryRow(t).Scan(&u.Id, &u.Qq, &u.Tag)
	if err != nil {
		return nil
	}
	return &u
}
