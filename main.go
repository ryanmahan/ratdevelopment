package main

import (
	"flag"
	"github.com/rs/cors"
	"ratdevelopment/DB"
	"ratdevelopment/api"
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

	// create server with database session
	// this initialization should be done here, as the above defer must be done before the main loop
	server := &api.Server{DBSession: databaseSession}
	server.InitServer(hostIPs)

	handler := cors.Default().Handler(server.GetRouter())
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:8080", "http://localhost:8081"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler = c.Handler(handler)

	log.Fatal(http.ListenAndServe(":8081", handler))
}
