package main

import (
	"github.com/slayer321/bookstore/pkg/db"
	"github.com/slayer321/bookstore/pkg/server"
)

func main() {

	session := db.NewSession()

	server.CreateServer(&session.Session)
}
