package DB

import (
	"github.com/gocql/gocql"
	"log"
	"strings"
	"time"
)

//DatabaseSession is a container for existing cassandra session
type DatabaseSession struct {
	*gocql.Session
}


func NewDBSession(hosts ...string) (*DatabaseSession, error) {
	log.Printf("Cassandra IPs: %s", strings.Join(hosts, ", "))
	db := gocql.NewCluster(hosts...)
	db.ProtoVersion = 4
	db.ConnectTimeout = 5 * time.Second
	db.Keyspace = "defaultks"
	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}
	return &DatabaseSession{session}, nil
}
