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
	db.Timeout = time.Minute * 2
	db.Keyspace = "defaultks"
	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}
	return &DatabaseSession{session}, nil
}
