package api

import (
	"net/http"
	"ratdevelopment/DB"
	"encoding/json"
	"fmt"
	"strings"
)

// Server is a struct that contains DB session and router info, to better consolidate and modularize API requests.
type Server struct {
	// DBSession is essentially a wrapper for the database session, and here for modularity. In the future, defining interfaces that implement multiple databases would be a better option.
	DBSession DB.FileBrowserDBSession
	loggers   *serverLogs
	router    *requestRouter
}

// InitServer initializes the router and the loggers for the server.
func (s *Server) InitServer(hostIPs string) {

	// Initalize server loggers. Initialization code for the loggers is contained in serverLogs.go
	s.loggers = &serverLogs{}
	s.loggers.initLogs()

	// set router of server to gorilla multiplexer
	s.router = &requestRouter{}
	s.router.routerInit()
	s.SetRoutes()
}

// GetRouter is a function that allows for access to the HTTP Router / Multiplexer. The router is not public since it shouldn't be changed once the server is initialized.
func (s *Server) GetRouter() http.Handler {
	return s.router
}

// SetRoutes sets the routes that the server will handle, such as the /api/tenant or /api/system GET requests. How this is written is dependent on the type wrapped in the requestRouter struct.
func (s *Server) SetRoutes() {
	// Please do not remove teapot code, I really like tea! :)
	s.router.HandleFunc("/api/teapot", s.teapot()).Methods("GET", "PUT", "POST", "HEAD", "TRACE", "OPTIONS", "DELETE", "CONNECT")
	// Thank you! :)
	// - Dan
	s.router.HandleFunc("/api", s.handleAPI()).Methods("GET")
	s.router.HandleFunc("/api/tenants", s.tenants()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}", s.getTenant()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/snapshots", s.getLatestSnapshotsByTenant()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems", s.getTenantSystems()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/snapshots/{timestamp}", s.getSnapshotByTenantSerialNumberAndDate(false)).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/snapshots/{timestamp}/download", s.getSnapshotByTenantSerialNumberAndDate(true)).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/timestamps", s.getValidTimestampsForSerialNumber()).Methods("GET")
	// We can wrap these handler functions in a call like this:
	// s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/timestamps", s.isAdmin(s.getValidTimestampsForSerialNumber())).Methods("GET")
	// and in isAdmin we can check for admin, and call the function contained in the parameters. This is why we return a function in all other methods,
	// in case there is some validation we need to do.
}

// Start of handler definitions. These should be identical to the old handler definitions, could potentially be put in their own file if we use something else for our "SetRoutes" method.

// handleAPI is just a test response to the /api/ request
func (s *Server) handleAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The API is working! Nothing else to see here. TODO: put documentation here, or example requests.")
		return
	}
}

func (s *Server) teapot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Please do not remove teapot code, I really like tea! :)
		w.WriteHeader(418)
		fmt.Fprintf(w, "I am a teapot! Have some tea! :)\n" + Teapot())
		// Thank you! :)
		// - Dan

	}
}

// tenants returns the handler function for hte /api/tenants/ request, which should be a list of tenants
func (s *Server) tenants() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO

		type tenantListStructure struct {
			TenantList []string `json:"tenants"`
		}
		w.Header().Set("Content-Type", "application/json")

		tenants, err := s.DBSession.GetValidTenants()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The system had a database error when obtaining the tenants of the system. The database may not have any data. Contact admins.")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		// Remake a set in O(n) to prove that go doesn't need a built-in set
		tenantMap := make(map[string]bool)
		for _, tenant := range tenants {
			if _, found := tenantMap[tenant]; !found {
				tenantMap[tenant] = true
			}
		}

		tenantSet := make([]string, len(tenantMap))
		i := 0
		for k := range tenantMap {
			tenantSet[i] = k
			i++
		}

		retTenantList := tenantListStructure{
			TenantList: tenantSet,
		}

		marshalledData, err := json.Marshal(retTenantList)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The system had an issue marshalling the tenant list json. Contact admins.")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		fmt.Fprintf(w, "%s", marshalledData)

		return
	}
}

// getTenant gets information about a specific tenant with name {name}
func (s *Server) getTenant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		// TODO refactor this at some point!!
		// But also this is probably an easy way to return stuff as JSON if we want to convert to that format with other endpoints.
		type tenantStructure struct {
			Tenant        string `json:"tenant"`
			SystemCount   int    `json:"systemCount"`
			SnapshotCount int    `json:"snapshotCount"`
		}
		var tenantData tenantStructure

		// Get the name of the tenant from the request
		params := s.router.getParams(r)
		tenantData.Tenant = params["name"]

		// Get number of systems
		systems, err := (s.DBSession).GetSystemsOfTenant(tenantData.Tenant)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The system had a database error when obtaining the systems of that tenant. Does that tenant exist?")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		tenantData.SystemCount = len(systems)

		// Get the tenant's snapshots
		snapshots, err := (s.DBSession).GetLatestSnapshotsByTenant(tenantData.Tenant)

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
		params := s.router.getParams(r)

		tenantName := params["name"]

		snapshots, err := (s.DBSession).GetLatestSnapshotsByTenant(tenantName)

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
		params := s.router.getParams(r)

		// TODO: this should eventually (in a steel thread way) be written so it returns a JSON, since JSON is
		// much more predictable as an API return type, and much easier to process for those using the API.
		// Consistency is key!
		tenantName := params["name"]

		serialNumberStrings, err := (s.DBSession).GetSystemsOfTenant(tenantName)
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
		params := s.router.getParams(r)

		tenantName := params["name"]
		serialNumberString := params["sernum"]
		timestamp := params["timestamp"]

		s.loggers.Info.Println(timestamp)

		snapshot, err := (s.DBSession).GetSnapshotByTenantSerialNumberAndDate(tenantName, serialNumberString, timestamp)

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
		params := s.router.getParams(r)

		tenantName := params["name"]
		serialNumber := params["sernum"]

		timestamps, err := (s.DBSession).GetValidTimestampsOfSystem(tenantName, serialNumber)
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
