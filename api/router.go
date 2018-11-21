package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"ratdevelopment-backend/DB"
	"encoding/json"
	"fmt"
	"strings"
	"log"
	"os"
)

// serverOutput is a struct designed to encapsulate the Trace, Info, Warning, and Error loggers that need to be used by the server and handling functions.
type serverOutput struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

// requestRouter is a wrapper for the router. Change this if you'd like to use a different router such as httprouter or roll your own net/http router.
type requestRouter struct {
	*mux.Router
}

// Server is a struct that contains DB session and router info, to better consolidate and modularize API requests.
type Server struct {
	DBSession *DB.DatabaseSession
	loggers   *serverOutput

	router    requestRouter
}
// DBSession is essentially a wrapper for the database session, and here for modularity. In the future, defining interfaces that implement multiple databases would be a better option.

// InitServer initializes the router and the loggers for the server.
func (s *Server) InitServer() {

	// Create Trace, Info, Warning, Error loggers
	Trace := log.New(os.Stdout,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info := log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning := log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error := log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	// add loggers to struct
	s.loggers = &serverOutput{
		Trace:   Trace,
		Info:    Info,
		Warning: Warning,
		Error:   Error,
	}

	// set router of server to gorilla multiplexer
	s.router = requestRouter{mux.NewRouter()}
}

// GetRouter is a function that allows for access to the HTTP Router / Multiplexer. The router is not public since it shouldn't be changed once the server is initialized.
func (s *Server) GetRouter() http.Handler {
	return s.router
}

// SetRoutes sets the routes that the server will handle, such as the /api/tenant or /api/system GET requests. How this is written is dependent on the type wrapped in the requestRouter struct.
func (s *Server) SetRoutes() {
	s.router.HandleFunc("/api/", s.handleAPI()).Methods("GET")
	s.router.HandleFunc("/api/tenants/", s.tenants()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/", s.getTenant()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/snapshots/", s.getLatestSnapshotsByTenant()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/", s.getTenantSystems()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/{timestamp}", s.getSnapshotByTenantSerialNumberAndDate(false)).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/{timestamp}/download", s.getSnapshotByTenantSerialNumberAndDate(true)).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/timestamps/", s.getValidTimestampsForSerialNumber()).Methods("GET")
}

// handleAPI is just a test response to the /api/ request
func (s *Server) handleAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The API is working! Nothing else to see here. TODO: put documentation here, or example requests.")
		return
	}
}

// tenants returns the handler function for hte /api/tenants/ request, which should be a list of tenants
func (s *Server) tenants() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		return
	}
}

// getTenant gets information about a specific tenant with name {name}
func (s *Server) getTenant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		// TODO refactor this at some point!!
		type tenantStructure struct {
			Tenant        string `json:"tenant"`
			SystemCount   int    `json:"systemCount"`
			SnapshotCount int    `json:"snapshotCount"`
		}
		var tenantData tenantStructure

		// Get the name of the tenant from the request
		params := mux.Vars(r)
		tenantData.Tenant = params["name"]

		// Get number of systems
		systems, err := (*s.DBSession).GetSystemsOfTenant(tenantData.Tenant)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The system had a database error when obtaining the systems of that tenant. Does that tenant exist?")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		tenantData.SystemCount = len(systems)

		// Get the tenant's snapshots
		snapshots, err := (*s.DBSession).GetLatestSnapshotsByTenant(tenantData.Tenant)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The system had a database error when obtaining the snapshots of that tenant. Does that tenant exist?")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		tenantData.SnapshotCount = len(snapshots)
		marshalledData, err := json.Marshal(tenantData)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The system had an issue marshalling the tenant data json. Contact admins.")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		fmt.Fprintf(w, "%s", marshalledData)

	}
}

func (s *Server) getLatestSnapshotsByTenant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		tenantName := params["name"]

		snapshots, err := (*s.DBSession).GetLatestSnapshotsByTenant(tenantName)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "[%s]", strings.Join(snapshots, ","))
	}
}

func (s *Server) getTenantSystems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		tenantName := params["name"]

		serialNumberStrings, err := (*s.DBSession).GetSystemsOfTenant(tenantName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "[\"%s\"]", strings.Join(serialNumberStrings, "\",\""))
	}
}

func (s *Server) getSnapshotByTenantSerialNumberAndDate(download bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		tenantName := params["name"]
		serialNumberString := params["sernum"]
		timestamp := params["timestamp"]

		snapshot, err := (*s.DBSession).GetSnapshotByTenantSerialNumberAndDate(tenantName, serialNumberString, timestamp)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if download {
			w.Header().Del("Content-Disposition")
			w.Header().Set("Content-Disposition", "attachment; filename=\""+serialNumberString+"-"+timestamp+".json\"")
		}

		fmt.Fprint(w, snapshot)
	}
}

func (s *Server) getValidTimestampsForSerialNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		tenantName := params["name"]
		serialNumber := params["sernum"]

		timestamps, err := (*s.DBSession).GetValidTimestampsOfSystem(tenantName, serialNumber)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		timestampStrings := DB.TimestampsToStrings(timestamps)
		fmt.Fprintf(w, "[\"%s\"]", strings.Join(timestampStrings, "\",\""))
	}
}
