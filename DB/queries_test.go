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
		session.GetLatestSnapshotsByTenant("hpe")
	}

}
