package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/slayer321/bookstore/pkg/db"
	"github.com/slayer321/bookstore/pkg/server"
)

func main() {

	session := db.NewSession()

	var id gocql.UUID
	var text string

	if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tweet:", id, text)

	server.CreateServer()
}
