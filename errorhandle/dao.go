package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MyDB struct {
	conn *sql.DB
}

func NewMyDB(dbUser, dbPass, dbHost string, dbPort int, dbName string) (MyDB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		return MyDB{}, err
	}
	err = db.Ping()
	if err != nil {
		return MyDB{}, err
	}
	return MyDB{db}, nil
}

func (myDB MyDB) Close() {
	if myDB.conn != nil {
		myDB.conn.Close()
	}
}

func (myDB MyDB) Get(id int) (*User, error) {
	u := &User{}
	err := myDB.conn.QueryRow("select id,name from user where id=?", id).Scan(&u.Id, &u.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user not found : %d\n", id) //log for warning,ignore the error
			return u, nil
		} else {
			return u, err
		}
	}
	return u, nil
}

type User struct {
	Id   int
	Name string
}
