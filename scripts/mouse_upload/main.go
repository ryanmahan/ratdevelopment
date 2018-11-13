package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"ratdevelopment-backend/DB"
)

func main() {

	//TODO: use a flag for accepting the path
	if(len(os.Args) < 2){
		log.Fatal("Must supply path to directory containing data dump")
	}
	path := os.Args[1]

	// Open the directory...
	directory, err := os.Open(path)
	if err != nil { log.Fatal(err) }

	// close the directory when main finishes
	defer directory.Close()

	// Readdir returns an array of FileInfo objects, we use those to open each file in the datadump
	// Readdir takes in an integer representing the number of files/directories to get the FileInfo of
	//	any value <= 0 will make Readdir get the info for all files/directories
	files, err := directory.Readdir(-1)
	if err != nil { log.Fatal(err) }

	// Open a single db session for all operations
	db, err := DB.NewDBSession()
	if err != nil { log.Fatal(err) }

	count := 0

	// Iterate over all of the fileInfo, ignoring the index with a _
	for _, fileInfo := range files{

		// make sure its actually a file and not a directory...
		if(!fileInfo.IsDir()) {

			count++
			fmt.Printf("Processessing file #%d\n", count )

			// open the file, need to include the path as the fileInfo doesnt include it
			sys_file, err := os.Open(path + fileInfo.Name())
			if err != nil { log.Fatal(err) }

			// read in the contents of the file into a byte array
			bytes, err := ioutil.ReadAll(sys_file)
			if err != nil { log.Fatal(err) }

			db.InsertSnapshot(bytes)
			sys_file.Close()
		}
	}

}
