package DB

type FileBrowserDBSession interface {
	GetLatestSnapshotsByTenant(string) ([]string, error)
}

func (db *DatabaseSession) GetLatestSnapshotsByTenant(tenant string) ([]string, error) {

	iter := db.Session.Query("SELECT snapshot FROM latest_snapshots_by_tenant WHERE tenant = ?", tenant).Iter()

	snapshots := make([]string, iter.NumRows())
	var jsonBlob string
	i := 0
	for iter.Scan(&jsonBlob) {
		snapshots[i] = jsonBlob
		i++
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return snapshots, nil
}
