package main

import (
	"fmt"
	"net/http"
	"ratdevelopment-backend/DB"
	"strconv"
	"strings"
)

//Env is a type holding the environment, particularly the cassandra session in use
type Env struct {
	session DB.FileBrowserDBSession
}

//MakeHandler creates a handler function for each query type.  query is the function to be handled, and display is the function to use the result of the query
func (env *Env) MakeHandler(query func([]interface{}) ([]string, error), contentType string, display func(http.ResponseWriter, []string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		//get properties needed in query
		var properties []interface{}
		popErr, index := CheckPopulation(w, properties...)
		if popErr {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Malformed query string: property number", index, "must have a value")
			return
		}
		results, err := query(properties)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
			Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		w.Header().Set("Content-Type", contentType)
		display(w, results)
	}
}

//CheckPopulation ensures that properties queried in URL have been properly populated, and also returns the index of the unpopulated field
func CheckPopulation(w http.ResponseWriter, vals ...interface{}) (bool, int) {
	for ind, val := range vals {
		i, ok := val.(int)
		if ok && i == 0 {
			return true, ind
		} else {
			s, ok := val.(string)
			if ok && len(s) == 0 {
				return true, ind
			}
		}
	}
	return false, 0
}

//DrawSingleJSON is a utility function to display one JSON file to the HTTP response
func (env *Env) DrawSingleJSON(w http.ResponseWriter, json []string) {
	fmt.Fprint(w, json[0])
}

//DrawMultipleJSON is a utility function to display a list of JSON files to the HTTP response
func (env *Env) DrawMultipleJSON(w http.ResponseWriter, json []string) {
	fmt.Fprintf(w, "[%s]", strings.Join(json, ","))
}

//translateGetLatestSnapshotsByTenant is a translator for makeHandler of GetLatestSnapshotsByTenant
func (env *Env) translateGetLatestSnapshotsByTenant(props []interface{}) ([]string, error) {
	return env.session.GetLatestSnapshotsByTenant(props[0].(string))
}

//translateGetLatestSnapshotsByTenant is a translator for makeHandler of GetLatestSnapshotsByTenant
func (env *Env) translateGetTimedSystemSnapshotByTenant(props []interface{}) ([]string, error) {
	snapshot, error := env.session.GetTimedSystemSnapshotByTenant(props[0].(string), props[1].(string), props[2].(int))
	if error == nil {
		output := make([]string, 1)
		output[0] = snapshot
		return output, error
	}
	return nil, error
}

//translateGetValidTimestampsOfSystem is a translator for makeHandler of GetValidTimestampsOfSystem
func (env *Env) translateGetValidTimestampsOfSystem(props []interface{}) ([]string, error) {
	times, err := env.session.GetValidTimestampsOfSystem(props[0].(string), props[1].(int))
	timeStrings := env.session.TimestampsToStrings(times)
	return timeStrings, err
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

func (env *Env) handleGetTimedSnapshotsByTenant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()
	tenant := query.Get("tenant")
	stamp := query.Get("timestamp")
	var sysID int
	sysID, err := strconv.Atoi(query.Get("systemID"))
	if len(tenant) == 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Malformed query string: tenant must have a value")
		return
	}
	snapshot, err := env.session.GetTimedSystemSnapshotByTenant(tenant, stamp, sysID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
		Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "[%s]", snapshot)
}
