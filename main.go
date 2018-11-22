package main

import (
	"flag"
	"github.com/rs/cors"
	"ratdevelopment-backend/DB"
	"ratdevelopment-backend/api"
	"net/http"
	"log"
)

func init() {

	flag.StringVar(&hostIPs, "cassandra_ips", "10.10.10.31", "Pass the ips of the cassandra hosts")
	flag.Parse()
}

var hostIPs string

func main() {

	databaseSession, err := DB.NewDBSession(hostIPs)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer databaseSession.Close()

	// env := &Env{session: session}
	// mux := http.NewServeMux()

	// mux.HandleFunc("/GetLatestSnapshotsByTenant", env.handleGetLatestSnapshotByTenant)
	// mux.HandleFunc("/GetSnapshotByTenantSerialNumberAndDate", env.handleGetSnapshotByTenantSerialNumberAndDate)
	// mux.HandleFunc("/GetValidTimestampsForSerialNumber", env.handleGetValidTimestampsForSerialNumber)
	// mux.HandleFunc("/GetTenantSystems", env.handleGetTenantSystems)

	// create server with database seloool aaron is ur research hookupssion


	server := &api.Server{
		DBSession: databaseSession,
	}

	server.InitServer()

	// This redirects all setting of routes to the api package, if this is not called then no routes will be handled
	server.SetRoutes()

	handler := cors.Default().Handler(server.GetRouter())
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:8080", "http://localhost:8081"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler = c.Handler(handler)

	log.Fatal(http.ListenAndServe(":8081", handler))
}
