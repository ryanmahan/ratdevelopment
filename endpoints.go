package main

import (
	"errors"
	"fmt"
	"net/http"
	"ratdevelopment-backend/DB"
	"strings"
)

//Env is a type holding the environment, particularly the cassandra session in use
type Env struct {
	session DB.FileBrowserDBSession
}

//MakeHandler creates a handler function for each query type.  query is the function to be handled, and display is the function to use the result of the query
func (env *Env) MakeHandler(query func([]interface{}) ([]string, error), contentType string, display func(http.ResponseWriter, []string), propertyNames ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		//get properties needed in query
		properties := make([]interface{}, len(propertyNames))
		q := r.URL.Query()
		for i, name := range propertyNames {
			properties[i] = q.Get(name)
		}

		popErr, index := CheckPopulation(w, properties...)
		if popErr {
			w.WriteHeader(http.StatusBadRequest)
			if index == -1 {
				fmt.Fprint(w, "Malformed query string: not enough arguments provided")
			} else {
				fmt.Fprint(w, "Malformed query string: property number ", index, " must have a value")
			}
			return
		}
		results, err := query(properties)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
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
	if len(vals) == 0 || forAll(vals, func(s string) bool { return len(s) == 0 }) {
		return true, -1
	}
	for ind, val := range vals {
		i, ok := val.(int)
		if ok && i == 0 {
			return true, ind
		}
		s, ok := val.(string)
		if ok && len(s) == 0 {
			return true, ind
		}
	}
	return false, 0
}

func forAll(strings []interface{}, fn func(string) bool) bool {
	for _, val := range strings {
		if !fn(val.(string)) {
			return false
		}
	}
	return true
}

//DrawSingleJSON is a utility function to display one JSON file to the HTTP response
func (env *Env) DrawSingleJSON(w http.ResponseWriter, json []string) {
	fmt.Fprint(w, json[0])
}

//DrawMultipleJSON is a utility function to display a list of JSON files to the HTTP response
func (env *Env) DrawMultipleJSON(w http.ResponseWriter, json []string) {
	fmt.Fprintf(w, "[%s]", strings.Join(json, ","))
}

//DrawDirect is a utility function to directly write a string to the HTTP response
func (env *Env) DrawDirect(w http.ResponseWriter, items []string) {
	fmt.Fprint(w, items)
}

//translateGetLatestSnapshotsByTenant is a translator for makeHandler of GetLatestSnapshotsByTenant
func (env *Env) translateGetLatestSnapshotsByTenant(props []interface{}) ([]string, error) {
	if len(props) == 0 {
		return nil, errors.New("Malformed query string: tenant must have a value")
	}
	return env.session.GetLatestSnapshotsByTenant(props[0].(string))
}

//translateGetTimedSnapshotByTenant is a translator for makeHandler of GetLatestSnapshotsByTenant
func (env *Env) translateGetTimedSnapshotByTenant(props []interface{}) ([]string, error) {
	snapshot, error := env.session.GetTimedSnapshotByTenant(props[0].(string), props[1].(string), props[2].(int))
	if error == nil {
		return DB.MakeSingleStringSlice(snapshot), error
	}
	return nil, error
}

//translateGetValidTimestampsOfSystem is a translator for makeHandler of GetValidTimestampsOfSystem
func (env *Env) translateGetValidTimestampsOfSystem(props []interface{}) ([]string, error) {
	times, err := env.session.GetValidTimestampsOfSystem(props[0].(string), props[1].(int))
	timeStrings := DB.TimestampsToStrings(times)
	return timeStrings, err
}

//---------------------------------------------FOR CONCISION-------------------------------------------------
//MakeLatestSnapshotsHandler is the top-level concision to make the handler for GetLatestSnapshotsByTenant
func (env *Env) MakeLatestSnapshotsHandler() http.HandlerFunc {
	return env.MakeHandler(env.translateGetLatestSnapshotsByTenant, "application/json", env.DrawMultipleJSON, "tenant")
}

//MakeTimedSnapshotHandler is the top-level concision to make the handler for GetTimedSnapshotByTenant
func (env *Env) MakeTimedSnapshotHandler() http.HandlerFunc {
	return env.MakeHandler(env.translateGetTimedSnapshotByTenant, "application/json", env.DrawSingleJSON, "tenant", "time", "serNum")
}

//MakeTimestampHandler is the top-level concision to make the handler for GetValidTimestampsOfSystem
func (env *Env) MakeTimestampHandler() http.HandlerFunc {
	return env.MakeHandler(env.translateGetValidTimestampsOfSystem, "text/plain", env.DrawDirect, "tenant", "sernum")
}
