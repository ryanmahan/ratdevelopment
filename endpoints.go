package main

import (
	"fmt"
	"net/http"
	"./DB"
	"strings"
)

type Env struct {
	session DB.FileBrowserDBSession
}

func (env *Env) handleGetLatestSnapshotsByTenant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tenant := r.URL.Query().Get("tenant")
	if len(tenant) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Malformed query string: tenant must have a value")
		return
	}
	snapshots, err := env.session.GetLatestSnapshotsByTenant(tenant)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
		Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "[%s]", strings.Join(snapshots, ","))
}
