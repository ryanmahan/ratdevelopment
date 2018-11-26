package api

import (
	"log"
	"os"
)

// serverOutput is a struct designed to encapsulate the Trace, Info, Warning, and Error loggers that need to be used by the server and handling functions.
type serverLogs struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func (s *serverLogs) initLogs() {

	// Create Trace, Info, Warning, Error loggers
	s.Trace = log.New(os.Stdout,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	s.Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	s.Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	s.Error = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
