package DB

import (
	"github.com/gocql/gocql"
)

type DatabaseSession struct {
	*gocql.Session
}

func NewDBSession() (*DatabaseSession, error) {
	db := gocql.NewCluster("localhost")
	db.Keyspace = "defaultks"
	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}
	return &DatabaseSession{session}, nil
}
