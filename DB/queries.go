package DB

import (
	"time"
)

//FileBrowserDBSession is an interface for querying the database session
type FileBrowserDBSession interface {
	GetLatestSnapshotsByTenant(string) ([]string, error)
	GetTimedSnapshotByTenant(string, string, int) (string, error)
	GetValidTimestampsOfSystem(string, int) ([]time.Time, error)
}

//GetLatestSnapshotsByTenant returns slice of JSON blobs for the latest snapshots of all systems owned by a tenant
func (db *DatabaseSession) GetLatestSnapshotsByTenant(tenant string) ([]string, error) {
	return db.RunQuery("SELECT snapshot FROM latest_snapshots_by_tenant WHERE tenant = ?", tenant)
}

//GetTimedSystemSnapshotByTenant gets the JSON blob of a system at a specified timestamp
func (db *DatabaseSession) GetTimedSnapshotByTenant(tenant, time string, sernum int) (string, error) {
	stamp, err := StringToTimestamp(time)
	if err != nil {
		return "", err
	}
	snapshots, err := db.RunQuery("SELECT snapshot FROM snapshots_by_serial_number WHERE tenant = ? AND serial_number = ? AND time = ?", tenant, sernum, stamp)
	if err != nil {
		return "", err
	}
	return snapshots[0], nil
}

const timeFormat string = time.RFC1123

//TimestampsToStrings converts a time slice to string slice for convenience
func TimestampsToStrings(times []time.Time) []string {
	timestamps := make([]string, len(times))
	for i, stamp := range times {
		timestamps[i] = stamp.Format(timeFormat)
	}
	return timestamps
}

//StringToTimestamp parses a timestamp to a time object for use in queries
func StringToTimestamp(stamp string) (time.Time, error) {
	return time.Parse(timeFormat, stamp)
}

//MakeSingleStringSlice is a util function to package a string in a slice quickly
func MakeSingleStringSlice(item string) []string {
	output := make([]string, 1)
	output[0] = item
	return output
}

//MakeSingleTimeSlice is a util function to package a time.Tim in a slice quickly
func MakeSingleTimeSlice(item time.Time) []time.Time {
	output := make([]time.Time, 1)
	output[0] = item
	return output
}

//GetValidTimestampsOfSystem returns a slice of strings that represent valid timestamps to index by for a system
func (db *DatabaseSession) GetValidTimestampsOfSystem(tenant string, sernum int) ([]time.Time, error) {
	iter := db.Session.Query("SELECT time FROM snapshots_by_serial_number WHERE tenant = ? AND serial_number = ?", tenant, sernum).Iter()
	items := make([]time.Time, iter.NumRows())
	for i := 0; i < len(items); i++ {
		var stamp time.Time
		iter.Scan(&stamp)
		items[i] = stamp
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return items, nil
}

//RunQuery is a helper function to easily get slice of strings that is returned from query, with error handling
func (db *DatabaseSession) RunQuery(query string, args ...interface{}) ([]string, error) {
	iter := db.Session.Query(query, args...).Iter()
	items := make([]string, iter.NumRows())
	for i, item := range items {
		iter.Scan(&item)
		items[i] = item
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return items, nil
}
