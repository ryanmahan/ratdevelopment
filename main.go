package main

import (
	"log"
	"net/http"
	"os"
	"ratdevelopment-backend/DB"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {

	Trace = log.New(os.Stdout,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	session, err := DB.NewDBSession()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer session.Close()
	env := &Env{session: session}
	//http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/GetLatestSnapshotsByTenant", env.handleGetLatestSnapshotsByTenant)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
