package db

import (
	bolt "go.etcd.io/bbolt"
	"log"
)

// Local database created by Bolt
var Local *bolt.DB

func init() {
	var err error
	// Open the app_crupi.db data file in your current directory.
	// It will be created if it doesn't exist.
	Local, err = bolt.Open("./data/app.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}
