package DB

import (
	"github.com/gocql/gocql"
)

//DatabaseSession is a container for existing cassandra session
type DatabaseSession struct {
	*gocql.Session
}

//NewDBSession creates a new database session.  In the future replace localhost with AWS cassandra instance
func NewDBSession() (*DatabaseSession, error) {
	db := gocql.NewCluster("localhost")
	db.Keyspace = "defaultks"
	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}
	return &DatabaseSession{session}, nil
}
