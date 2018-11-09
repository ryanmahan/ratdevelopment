package DB

type FileBrowserDBSession interface {
	GetLatestSnapshotsByTenant(string) ([]string, error)
}

func (db *DatabaseSession) GetLatestSnapshotsByTenant(tenant string) ([]string, error) {
	return db.RunQuery("SELECT snapshot FROM latest_snapshots_by_tenant WHERE tenant = ?", tenant)
}

func (db *DatabaseSession) GetTimedSystemSnapshotByTenant(tenant string, time string, sernum int) (string, error) {
	snapshots, err = db.RunQuery("SELECT snapshot FROM snapshots_by_serial_number WHERE tenant = ? AND serial_number = ? AND time = ?", tenant, sernum, time)
	if err != nil {
		return nil, err
	}
	return snapshots[0], nil
}

func (db *DatabaseSession) GetValidTimestampsOfSystem(tenant string, sernum, offset, count int) ([]string, error) {
	return db.RunQuery("SELECT time FROM snapshots_by_serial_number WHERE tenant = ? AND serial_number = ?", tenant, sernum)
}

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
