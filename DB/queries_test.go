package DB

import "testing"

func BenchmarkGetLatestSnapshotsByTenant(b *testing.B) {
	session, err := NewDBSession()
	if err != nil {
		b.Error(err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		session.GetLatestSnapshotsByTenant("hpe")
	}

}
