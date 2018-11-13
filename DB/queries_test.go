package DB

import (
	"testing"
)

func BenchmarkGetLatestSnapshotsByTenant(b *testing.B) {
	session, err := NewDBSession()
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
	_, err = session.GetTimedSnapshotByTenant("hpe", stamps[0], 9996788)
	if err != nil {
		b.Error(err)
		return
	}
}
