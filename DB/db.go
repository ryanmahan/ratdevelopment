package DB

import (
	"github.com/gocql/gocql"
)

type FileBrowserDBSession interface {
	GetLatestSnapshotsByTenant(string) ([]string, error)
}

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

func (db *DatabaseSession) GetLatestSnapshotsByTenant(tenant string) ([]string, error) {

	iter := db.Session.Query("SELECT snapshot FROM defaultks.latest_snapshots_by_tenant WHERE tenant = ?", tenant).Iter()

	snapshots := make([]string, iter.NumRows())
	var jsonBlob string
	i := 0
	for iter.Scan(&jsonBlob) {
		snapshots[i] = jsonBlob
		i++
	}

	return snapshots, nil
}
