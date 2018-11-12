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

func BenchmarkGetLatestSnapshotsByTenant(b *testing.B) {
	session, err := NewDBSession(host)
	if err != nil {
		b.Error(err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := session.GetLatestSnapshotsByTenant("hpe")
		if err != nil {
			b.Error(err)
			return
		}
	}

}

func BenchmarkGetValidTimestamps(b *testing.B) {
	session, err := NewDBSession()
	if err != nil {
		b.Error(err)
		return
	}
	b.ResetTimer()
	_, err = session.GetValidTimestampsOfSystem("hpe", 9996788)
	if err != nil {
		b.Error(err)
		return
	}
}

func BenchmarkGetTimedSnapshot(b *testing.B) {
	session, err := NewDBSession()
	if err != nil {
		b.Error(err)
		return
	}
	b.ResetTimer()
	timestamps, err := session.GetValidTimestampsOfSystem("hpe", 9996788)
	if err != nil {
		b.Error(err)
		return
	}
	stamps := TimestampsToStrings(timestamps)
	if err != nil {
		b.Error(err)
		return
	}
	_, err = session.GetTimedSystemSnapshotByTenant("hpe", stamps[0], 9996788)
	if err != nil {
		b.Error(err)
		return
	}
}
