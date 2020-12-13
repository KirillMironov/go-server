package main

import (
	"database/sql"
	"testing"
)

const (
	connectionString = "postgres://postgres:postgres@35.210.228.180:5432/postgres"
)

func TestInsertUser(t *testing.T) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	//Begin tx

	user:= &User{}
	user.Username="Lisa"
	user.Password="123"
	user.Email="afs@gmail.com"

	insert(user, db)

	//Rollback tx
}
