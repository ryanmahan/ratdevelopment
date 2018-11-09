package main

import (
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
	directory, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer directory.Close()

	files, err := directory.Readdir(-1)

	if err != nil {
		log.Fatal(err)
	}

	db, err := DB.NewDBSession()
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range files{

		if(!fileInfo.IsDir()) {
			sys_file, err := os.Open(path + fileInfo.Name())
			if err != nil {
				log.Fatal(err)
			}

			bytes , err := ioutil.ReadAll(sys_file)

			if err != nil {
				log.Fatal(err)
			}

			db.InsertSnapshot(bytes)
			sys_file.Close()
		}
	}

}
