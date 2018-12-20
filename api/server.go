package api

import (
	"crypto"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"ratdevelopment/DB"
	"strings"
	"strconv"
	"encoding/hex"
	"github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
)

const (
	APIAudience = "https://mousefb/api"
	APIDomain   = "https://rat-dev.auth0.com/"
	APICert     = `-----BEGIN CERTIFICATE-----
MIIC/TCCAeWgAwIBAgIJAvo3FE3TLwO+MA0GCSqGSIb3DQEBCwUAMBwxGjAYBgNV
BAMTEXJhdC1kZXYuYXV0aDAuY29tMB4XDTE4MTEwMTE3MjM1N1oXDTMyMDcxMDE3
MjM1N1owHDEaMBgGA1UEAxMRcmF0LWRldi5hdXRoMC5jb20wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQCtk25a9eiO+qjuM0bBh3F5foO0qiMG6mfYwBH1
SacA28GTX5NlA3HHdAqVAHNzqxpwC6dTsHSkbvfY1IIaMHe5nc364J+2YeshT1MB
1TuQWsx33s77QuTtOmlYXzCwT/6CGWO6IORCaJ2WnJh0wUXp667HUyKjlaP4bR/T
vEJaCRzVDBwngeGDFDJRfcciGMsR3e7N1ca/teuBAsvQV2M4Jj3FaQz6OvX+Heoo
UsE9GnODFzTsSLI2m4hnbtu2pnjgdCYlCeg4JbHjFMaXFkfWsUpNzAVxbxV6wkEM
HFNZMcqNQsl6HbaRuV/8PsMzDtJqsmg/TTjEQP34JkL3Hq0HAgMBAAGjQjBAMA8G
A1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFD6KDEmytvKbhgYCvdinUtp0IzsJMA4G
A1UdDwEB/wQEAwIChDANBgkqhkiG9w0BAQsFAAOCAQEAmOL8kkRK8onlwcVDMUny
ifVG4LOD3LbqPptSIjRiSbM19q537JnEofVemXbAPrCKtvcmPpSQ+ILnXXSdBjNX
yoNTW0sWrAdIzDR82tlV7VqkSf/Q/pQg8SN7LVigdHDU8URK4QFIIsyCR2d7Mbb9
o/kSHNPtZ1pv6ISPI4WYXTNAmWC1+ji2aNNRF1sBd39vjhV9TU+jQRBqPiLcCKMF
tU/HI8oIuySpPbwEbjid6qhdDKLErcrn0ITBO0jKjOCXuCsTDoVklmtRvPElNpDz
SW9H3FFam91En7aOmwjh1gwiELc9NiivQQGbMkcA4SuXRYlMD6452Hh0ee35kpv4
og==
-----END CERTIFICATE-----
`
)

// Server is a struct that contains DB session and router info, to better consolidate and modularize API requests.
type Server struct {
	// DBSession is essentially a wrapper for the database session, and here for modularity. In the future, defining interfaces that implement multiple databases would be a better option.
	DBSession DB.FileBrowserDBSession
	loggers   *serverLogs
	router    *requestRouter
	validator *auth0.JWTValidator
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

	// Initialize the validator to check tokens
	s.InitAuthHandlers()
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
	s.router.HandleFunc("/api/tenants/{name}/snapshots", s.getLatestSnapshotsByTenant()).Queries("searchString", "{options}").Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/snapshots", s.getLatestSnapshotsByTenant()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems", s.getTenantSystems()).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/snapshots/{timestamp}", s.getSnapshotByTenantSerialNumberAndDate(false)).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/snapshots/{timestamp}/download", s.getSnapshotByTenantSerialNumberAndDate(true)).Methods("GET")
	s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/timestamps", s.getValidTimestampsForSerialNumber()).Methods("GET")
	s.router.HandleFunc("/api/paginate/tenants/{page}", s.tenantsPaginated()).Queries("pageState", "{state}").Methods("GET")
	s.router.HandleFunc("/api/paginate/tenants/{page}", s.tenantsPaginated()).Methods("GET")
	s.router.HandleFunc("/api/paginate/tenant/{name}/snapshots/{page}", s.snapshotsPaginated()).Queries("pageState", "{state}").Methods("GET")
	s.router.HandleFunc("/api/paginate/tenant/{name}/snapshots/{page}", s.snapshotsPaginated()).Methods("GET")
	// We can wrap these handler functions in a call like this:
	// s.router.HandleFunc("/api/tenants/{name}/systems/{sernum}/timestamps", s.isAdmin(s.getValidTimestampsForSerialNumber())).Methods("GET")
	// and in isAdmin we can check for admin, and call the function contained in the parameters. This is why we return a function in all other methods,
	// in case there is some validation we need to do.
}

// Authorization functions

// Initialize the token validator and apply the authentication middleware
func (s *Server) InitAuthHandlers() {
	secret, err := loadPublicKey([]byte(APICert))

	if err != nil {
		panic(err)
	}

	secretProvider := auth0.NewKeyProvider(secret)
	configuration := auth0.NewConfiguration(secretProvider, []string{APIAudience}, APIDomain, jose.RS256)
	s.validator = auth0.NewValidator(configuration, nil)
	s.router.Use(s.authorizeUserExists)
}

// Get the public key from a certificate
func loadPublicKey(data []byte) (crypto.PublicKey, error) {
	input := data

	block, _ := pem.Decode(data)
	if block != nil {
		input = block.Bytes
	}

	cert, err1 := x509.ParseCertificate(input)
	if err1 == nil {
		return cert.PublicKey, nil
	}

	return nil, fmt.Errorf("certificate parse error: '%s'", err1)
}

// Check to see if the token is from an authorized user
func (s *Server) authorizeUserExists(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := s.validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Printf("Token is not valid: %#v\n", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
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
		fmt.Fprintf(w, "I am a teapot! Have some tea! :)")
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
		snapshots, err := (s.DBSession).GetLatestSnapshotsByTenant(tenantData.Tenant, "")

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
		searchString := params["options"]

		snapshots, err := (s.DBSession).GetLatestSnapshotsByTenant(tenantName, searchString)

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

func (s *Server) tenantsPaginated() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		params := s.router.getParams(r)

		w.Header().Set("Content-Type", "application/json")

		// TODO: cache the page state in the queries.go file so that you can just get the next page super smoothly, currently if you wanna get page 2 it doesn't care what page you're on
		// We'll have to keep maybe the page state on the frontend to pass in a request, so we can tell the database to start from there
		// This should definitely be done in a query parameter, not like our current /api/asdas/{name} fashion, since it's not necessary and you won't always have it
		type tenantPageJSON struct {
			TenantPage int `json:"tenantPage"`
			LastPage bool `json:"lastPage"`
			TenantCount int `json:"tenantCount"`
			Tenants []string `json:"tenants"`
		}
		var tenantData tenantPageJSON

		tenantPage := params["page"]
		page, err := strconv.Atoi(tenantPage)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		}

		pageOfTenants, pageReturned, lastPage, err := (s.DBSession).GetTenantPage(100, page)

		tenantData.TenantPage = pageReturned
		// lastPage is to help the frontend not allow the "next page" button to be active if this is the last page
		tenantData.LastPage = lastPage
		tenantData.Tenants = pageOfTenants
		tenantData.TenantCount = len(pageOfTenants)


		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The pagination returned an error! Please check that you're expecting this to have this many pages.")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		}

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

func (s *Server) snapshotsPaginated() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		params := s.router.getParams(r)

		w.Header().Set("Content-Type", "application/json")

		// TODO: cache the page state in the queries.go file so that you can just get the next page super smoothly, currently if you wanna get page 2 it doesn't care what page you're on
		// We'll have to keep maybe the page state on the frontend to pass in a request, so we can tell the database to start from there
		// This should definitely be done in a query parameter, not like our current /api/asdas/{name} fashion, since it's not necessary and you won't always have it
		type snapshotPageJSON struct {
			PageState string `json:"pageState"`
			SnapshotPage int `json:"snapshotPage"`
			LastPage bool `json:"lastPage"`
			SnapshotCount int `json:"snapshotCount"`
			Snapshots []Snapshot `json:"snapshots"`
		}

		var snapshotData snapshotPageJSON

		var state []byte
		if stateString, found := params["state"]; found {
			var err error
			state, err = hex.DecodeString(stateString)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
				s.loggers.Error.Fatal(err)
				s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			}
		}

		tenant := params["name"]
		snapshotPage := params["page"]

		page, err := strconv.Atoi(snapshotPage)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Potentially malformed API call, or internal application error!")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		}

		pageOfSnapshots, pageReturned, lastPage, pageState, err := (s.DBSession).GetSnapshotPageByTenant(tenant, 100, page, state)

		snapshotData.PageState = hex.EncodeToString(pageState)
		snapshotData.SnapshotPage = pageReturned
		// LastPage is to help the frontend not allow the "next page" button to be active if this is the last page

		for _, pg := range pageOfSnapshots  {

			var marshalledPage Snapshot
			err := json.Unmarshal([]byte(pg), &marshalledPage)
			if err != nil {
				println(pg)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "The system had an issue marshalling the tenant data json. Contact admins.")
				s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
				return
			}

			snapshotData.Snapshots = append(snapshotData.Snapshots, marshalledPage)
		}

		snapshotData.LastPage = lastPage
		snapshotData.SnapshotCount = len(pageOfSnapshots)


		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The pagination returned an error! Please check that you're expecting this to have this many pages.")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
		}

		marshalledData, err := json.Marshal(snapshotData)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "The system had an issue marshalling the tenant data json. Contact admins.")
			s.loggers.Error.Printf("Request:\n%#v\nError:\n%#v", r, err)
			return
		}

		fmt.Fprintf(w, "%s", marshalledData)
	}

}

func (s *Server) search() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
	}

}
