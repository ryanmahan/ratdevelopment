package DB

import (
	"flag"
	"testing"
)

func init() {
	flag.StringVar(&host, "cassandra_ip", "10.10.10.31", "Pass the IP of the Cassandra host")
	flag.Parse()
}

var host string

func TestTenantSerialNumbers(t *testing.T) {
	session, err := NewDBSession(host)
	if err != nil {
		t.Error(err)
		return
	}

	serialNums, err := session.GetSystemsOfTenant("hpe")
	if err != nil {
		t.Error(err)
		return
	}
	if len(serialNums) == 0 {
		t.Error("Expected serialNums to be populated but was empty")
		return
	}
	if serialNums[0] != "9996788" {
		t.Error("expecting serial number 9996788, got ", serialNums[0])
		return
	}
}

func TestGetLatestSnapshotsByTenant(t *testing.T) {
	session, err := NewDBSession(host)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = session.GetLatestSnapshotsByTenant("hpe")
	if err != nil {
		t.Error(err)
		return
	}

}

func TestGetValidTimestamps(t *testing.T) {
	session, err := NewDBSession(host)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = session.GetValidTimestampsOfSystem("hpe", "9996788")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetTimedSnapshot(t *testing.T) {
	session, err := NewDBSession(host)
	if err != nil {
		t.Error(err)
		return
	}

	timestamps, err := session.GetValidTimestampsOfSystem("hpe", "9996788")
	if err != nil {
		t.Error(err)
		return
	}
	stamps := TimestampsToStrings(timestamps)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = session.GetSnapshotByTenantSerialNumberAndDate("hpe", "9996788", stamps[0])
	if err != nil {
		t.Error(err)
		return
	}
}
