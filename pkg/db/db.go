package db

import (
	"log"

	"github.com/gocql/gocql"
)

type DB struct {
	gocql.Session
}

func NewSession() *DB {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Printf("Error creating session: %v", err)
	}

	return &DB{*session}
}
