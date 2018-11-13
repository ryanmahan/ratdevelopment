package DB

type FileBrowserDBSession interface {
	GetLatestSnapshotsByTenant(string) ([]string, error)
}

func (db *DatabaseSession) GetLatestSnapshotsByTenant(tenant string) ([]string, error) {
	//	This line makes a query to the database and puts the rows returned into an iterator. The Iter object, as per the
	//	gocql documentation, will page in 5000 rows at once.
	iter := db.Session.Query("SELECT snapshot FROM latest_snapshots_by_tenant WHERE tenant = ?", tenant).Iter()
	// Initially we make a slice of 0 strings with capacity iter.NumRows(). This will at max be 5000 because of the
	// paging discussed above.
	snapshots := make([]string, 0, iter.NumRows())
	var jsonBlob string
	//	This line will scan each row and put the returned value into the string jsonBlob. It may also make more requests
	// 	to the database and page in those results. (This is why the initial iter.NumRows() may not be equal to the
	// number of rows the query will eventually return.)
	for iter.Scan(&jsonBlob) {
		//	append will grow the slice if it is not big enough to hold the next string.
		snapshots = append(snapshots, jsonBlob)
	}
	//	This checks for errors upon closing the iterator
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return snapshots, nil
}
