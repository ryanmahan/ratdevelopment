package main

import (
	"fmt"
	"net/http"
	"ratdevelopment-backend/DB"
	"strings"
)

//Env is a type holding the environment, particularly the cassandra session in use
type Env struct {
	session DB.FileBrowserDBSession
}

func (env *Env) handleGetLatestSnapshotByTenant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	tenant := q.Get("tenant")
	//limit := q.Get("limit")

	if len(tenant) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Must supply a tenant ID")
		return
	}

	snapshots, err := env.session.GetLatestSnapshotsByTenant(tenant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
		Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "[%s]", strings.Join(snapshots, ","))
}

func (env *Env) handleGetSnapshotByTenantSerialNumberAndDate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	tenant := q.Get("tenant")
	serialNumberString := q.Get("serialNumber")
	timestamp := q.Get("timestamp")
	download := q.Get("download")

	if len(tenant) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Must supply a tenant ID")
		return
	}
	if len(serialNumberString) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Must supply a tenant ID")
		return
	}
	if len(timestamp) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Must supply a timestamp")
		return
	}

	snapshot, err := env.session.GetSnapshotByTenantSerialNumberAndDate(tenant, serialNumberString, timestamp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
		Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if download == "1" {
		w.Header().Del("Content-Disposition")
		w.Header().Set("Content-Disposition", "attachment; filename="+serialNumberString+".json")
	}

	fmt.Fprint(w, snapshot)
}

func (env *Env) handleGetValidTimestampsForSerialNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	tenant := q.Get("tenant")
	serialNumber := q.Get("serialNumber")

	//limit := q.Get("limit")

	if len(tenant) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Must supply a tenant ID")
		return
	}
	if len(serialNumber) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Must supply a timestamp")
		return
	}

	timestamps, err := env.session.GetValidTimestampsOfSystem(tenant, serialNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
		Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	timestampStrings := DB.TimestampsToStrings(timestamps)
	fmt.Fprintf(w, "[\"%s\"]", strings.Join(timestampStrings, "\",\""))
}

func (env *Env) handleGetTenantSystems(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	tenant := q.Get("tenant")

	if len(tenant) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Must supply a tenant ID")
		return
	}

	serialNumberStrings, err := env.session.GetSystemsOfTenant(tenant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
		Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "[\"%s\"]", strings.Join(serialNumberStrings, "\",\""))
}
