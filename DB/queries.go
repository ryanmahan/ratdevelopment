package DB

//FileBrowserDBSession is an interface for querying the database session
type FileBrowserDBSession interface {
	GetLatestSnapshotsByTenant(string) ([]string, error)
	GetTimedSystemSnapshotByTenant(string, string, int) (string, error)
	GetValidTimestampsOfSystem(string, int, int, int) ([]string, error)
}

//GetLatestSnapshotsByTenant returns slice of JSON blobs for the latest snapshots of all systems owned by a tenant
func (db *DatabaseSession) GetLatestSnapshotsByTenant(tenant string) ([]string, error) {
	return db.RunQuery("SELECT snapshot FROM latest_snapshots_by_tenant WHERE tenant = ?", tenant)
}

//GetTimedSystemSnapshotByTenant gets the JSON blob of a system at a specified timestamp
func (db *DatabaseSession) GetTimedSystemSnapshotByTenant(tenant, time string, sernum int) (string, error) {
	snapshots, err := db.RunQuery("SELECT snapshot FROM snapshots_by_serial_number WHERE tenant = ? AND serial_number = ? AND time = ?", tenant, sernum, time)
	if err != nil {
		return "", err
	}
	return snapshots[0], nil
}

//GetValidTimestampsOfSystem returns a slice of strings that represent valid timestamps to index by for a system
func (db *DatabaseSession) GetValidTimestampsOfSystem(tenant string, sernum, offset, count int) ([]string, error) {
	return db.RunQuery("SELECT time FROM snapshots_by_serial_number WHERE tenant = ? AND serial_number = ?", tenant, sernum)
}

//RunQuery is a helper function to easily get slice of strings that is returned from query, with error handling
func (db *DatabaseSession) RunQuery(query string, args ...interface{}) ([]string, error) {
	iter := db.Session.Query(query, args...).Iter()

	items := make([]string, iter.NumRows())
	var item string
	for i := 0; iter.Scan(&item); i++ {
		items[i] = item
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return items, nil
}
