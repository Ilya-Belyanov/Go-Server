package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func isExistUser(name string) bool {
	db, err := sql.Open("mysql", "root:Hox3taBev@/MusicMan")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from MusicMan.users where name = ?", name)
	if err != nil {
		panic(err)
	}
	return rows.Next()
}

func addNewUser(name string, password string) (int64, bool) {
	db, err := sql.Open("mysql", "root:Hox3taBev@/MusicMan")
	if err != nil {
		panic(err)
		return -1, true
	}
	defer db.Close()
	sqlScript := "INSERT INTO MusicMan.users(name, password) VALUES (?, ?)"
	res, err := db.Exec(sqlScript, name, password)
	if err != nil {
		panic(err)
		return -1, true
	}
	lastId, _ := res.LastInsertId()
	return lastId, false
}
