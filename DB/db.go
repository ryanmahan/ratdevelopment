package DB

import (
	"github.com/gocql/gocql"
	"time"
)

type DatabaseSession struct {
	*gocql.Session
}

func NewDBSession() (*DatabaseSession, error) {
	db := gocql.NewCluster("localhost")
	// Datadumps are big and gocql times out by default while loading it in
	// so we extend the time out duration
	db.ConnectTimeout = time.Minute * 10
	db.Keyspace = "defaultks"
	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}
	return &DatabaseSession{session}, nil
}
