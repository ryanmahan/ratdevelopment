package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

type mockSession struct{}

const expectedObtainedString = "\n...expected = %#v\n...obtained = %#v"

func (db *mockSession) GetLatestSnapshotsByTenant(tenant string) ([]string, error) {

	snapshots := []string{
		"{ SerialNumberInserv: \"1231234\", tenant: { authorized:[\"%TENANT%\"]}",
		"{ SerialNumberInserv: \"7162634\", tenant: { authorized:[\"%TENANT%\"]}",
		"{ SerialNumberInserv: \"1111111\", tenant: { authorized:[\"%TENANT%\"]}",
	}
	for i, s := range snapshots {
		snapshots[i] = strings.Replace(s, "%TENANT%", tenant, 1)
	}

	return snapshots, nil
}

func (db *mockSession) GetTimedSnapshotByTenant(tenant string, time string, sysID int) (string, error) {
	snapshot := "{ SerialNumberInserv: \"%SYSID%\", tenant: { authorized:[\"%TENANT%\"]}"
	snapshot = strings.Replace(snapshot, "%TENANT%", tenant, 1)
	snapshot = strings.Replace(snapshot, "%SYSID%", strconv.Itoa(sysID), 1)
	return snapshot, nil
}

func (db *mockSession) GetValidTimestampsOfSystem(tenant string, sysID int) ([]time.Time, error) {
	return []time.Time{time.Now()}, nil
}

func TestHandleGetLatestSnapshotsByTenantWithoutTenant(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/GetLatestSnapshotsByTenant", nil)
	env := Env{session: &mockSession{}}
	env.MakeLatestSnapshotsHandler().ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf(expectedObtainedString, http.StatusBadRequest, rec.Code)
	}

	expected := "Malformed query string: not enough arguments provided"
	if expected != rec.Body.String() {
		t.Errorf(expectedObtainedString, expected, rec.Body.String())
	}
}

func TestHandleGetLatestSnapshotsByTenantWithTextTenant(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/GetLatestSnapshotsByTenant?tenant=hpe", nil)

	env := Env{session: &mockSession{}}
	env.MakeLatestSnapshotsHandler().ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Errorf(expectedObtainedString, http.StatusOK, rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf(expectedObtainedString, "application/json", rec.Header().Get("Content-Type"))
	}

	expected :=
		"[{ SerialNumberInserv: \"1231234\", tenant: { authorized:[\"hpe\"]}," +
			"{ SerialNumberInserv: \"7162634\", tenant: { authorized:[\"hpe\"]}," +
			"{ SerialNumberInserv: \"1111111\", tenant: { authorized:[\"hpe\"]}]"
	if expected != rec.Body.String() {
		t.Errorf(expectedObtainedString, expected, rec.Body.String())
	}
}

func TestHandleGetLatestSnapshotsByTenantWithTenantID(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/GetLatestSnapshotsByTenant?tenant=264593856", nil)

	env := Env{session: &mockSession{}}
	env.MakeLatestSnapshotsHandler().ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Errorf(expectedObtainedString, http.StatusOK, rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf(expectedObtainedString, "application/json", rec.Header().Get("Content-Type"))
	}

	expected :=
		"[{ SerialNumberInserv: \"1231234\", tenant: { authorized:[\"264593856\"]}," +
			"{ SerialNumberInserv: \"7162634\", tenant: { authorized:[\"264593856\"]}," +
			"{ SerialNumberInserv: \"1111111\", tenant: { authorized:[\"264593856\"]}]"
	if expected != rec.Body.String() {
		t.Errorf(expectedObtainedString, expected, rec.Body.String())
	}
}
