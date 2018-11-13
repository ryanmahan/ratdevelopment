package DB

import (
	"encoding/json"
	"log"
	"time"
)

// The main method of loading a file into ALL tables
func (db *DatabaseSession) InsertSnapshot(jsonBytes []byte) (error){

	// Here we need to create a map with string keys and interface{} values,
	// then convert and store the byte array containing the json data into the said map.
	var values map[string]interface{}
	err := json.Unmarshal(jsonBytes, &values)

	if err != nil { return err }

	// Functions per table
	db.insertLatestByTenant(&values, &jsonBytes)
	db.insertSnapshotsBySerialNumber(&values, &jsonBytes)
	return nil
}

// Updates the latest_snapshot_by_tenant table with the info from a single json file
func (db *DatabaseSession) insertLatestByTenant(info *map[string]interface{}, jsonBlob *[]byte) (error) {

	serialNum := (*info)["serialNumberInserv"]

	// since the tenants is in an array, nested within tenants, nested within authorized
	// we have to do first access authorized,
	// then cast the contents to the same map as info before we can get at the tenants array
	tenants := (*info)["authorized"].(map[string]interface{})["tenants"]

	// to iterate over the array, we have to cast it to an array of interface
	// and only then cast each individual tenant to a string (casting to an array of strings breaks things)
	for _, tenant := range tenants.([]interface{}){
		tenant := tenant.(string)

		// Update the record
		// if none match the WHERE clause, this creates a new record
		err := db.Query("UPDATE latest_snapshots_by_tenant " +
						"SET snapshot = textAsBlob(?)" +
						"WHERE tenant = ? AND serial_number = ?",
						*jsonBlob, tenant, serialNum).Exec()
		if err != nil { log.Fatal(err) }
	}

	return nil
}

// Creates a record for each file for each tenant
// this works exactly the same as insertLastestByTenant,
// except it adds to a different table that also requires the date of the file
func (db *DatabaseSession) insertSnapshotsBySerialNumber(info *map[string]interface{}, jsonBlob *[]byte) (error) {

	serialNum := (*info)["serialNumberInserv"]

	// We cant directly convert to a time object, we have to convert it into a string
	// then parse the string following a specific format (the first argument of Parse)
	sysDate, err := time.Parse(time.RFC3339, (*info)["date"].(string))

	if err != nil { return err }

	tenants := (*info)["authorized"].(map[string]interface{})["tenants"]

	for _, tenant := range tenants.([]interface{}){
		tenant := tenant.(string)
		err := db.Query("INSERT INTO snapshots_by_serial_number " +
						"(tenant, serial_number, time, snapshot)" +
						"VALUES (?, ?, ?, textAsBlob(?)) IF NOT EXISTS",
						tenant, serialNum, sysDate, *jsonBlob).Exec()

		if err != nil { log.Fatal(err) }
	}

	return nil
}