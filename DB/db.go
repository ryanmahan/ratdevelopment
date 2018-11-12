package DB

import (
	"github.com/gocql/gocql"
	"log"
	"strings"
	"time"
)

type DatabaseSession struct {
	*gocql.Session
}

func NewDBSession(hosts ...string) (*DatabaseSession, error) {
	log.Printf("Cassandra IPs: %s", strings.Join(hosts, ", "))
	db := gocql.NewCluster(hosts...)
	db.ProtoVersion = 4
  db.ConnectTimeout = time.Minute * 10
	db.Keyspace = "defaultks"
	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}
	return &DatabaseSession{session}, nil
}
