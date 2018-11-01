package main

import (
	"fmt"
	"log"
	"net/http"
	"ratdevelopment-backend/DB"
	"strings"
)

type Env struct {
	session DB.FileBrowserDBSession
}

func main() {
	session, err := DB.NewDBSession()
	if err != nil {
		log.Fatal(err)
		return
	}
	env := &Env{session: session}
	//http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/GetLatestSnapshotsByTenant", env.handleGetLatestSnapshotsByTenant)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func (env *Env) handleGetLatestSnapshotsByTenant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tenant := r.URL.Query().Get("tenant")
	if len(tenant) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	snapshots, err := env.session.GetLatestSnapshotsByTenant(tenant)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "[%s]", strings.Join(snapshots, ","))
}
