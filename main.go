package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	println("Hello World!")

	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "test_space"
	if session, err := cluster.CreateSession(); err != nil {
		println("Failed to connect")
	} else {

		var comment, name string
		var datetime time.Time

		if err := session.Query(`INSERT INTO test_space.comments (comment, name, datetime) VALUES (?, ?, ?) IF NOT EXISTS`,
			"Hello World from inside Cassandra!", "Cassandra", time.Now()).Exec(); err != nil {
			println("0 Error occured...")
			fmt.Println(err)

		}

		if err := session.Query(`INSERT INTO test_space.comments (comment, name, datetime) VALUES (?, ?, ?) IF NOT EXISTS`,
			"Goodbye and gooday to you.", "Cassandra", time.Now()).Exec(); err != nil {
			println("1 Error occured...")
			fmt.Println(err)
		}
		iterator := session.Query(`SELECT * FROM test_space.comments`).Iter()
		for iterator.Scan(&comment, &datetime, &name) {
			println(name, "says:", comment, "at", datetime.String())
		}

		if err := iterator.Close(); err != nil {
			println("2 Error occured...")
			fmt.Println(err)
		}

		session.Close()
	}
}
