package DB

import (
	"encoding/json"
	"log"
	"time"
)

func (db *DatabaseSession) InsertSnapshot(jsonBytes []byte) (error){

	var values map[string]interface{}
	err := json.Unmarshal(jsonBytes, &values)

	if err != nil {
		return err
	}

	db.insertLastestByTenant(&values, &jsonBytes)
	db.insertSnapshotsBySerialNumber(&values, &jsonBytes)
	return nil
}


func (db *DatabaseSession) insertLastestByTenant(info *map[string]interface{}, jsonBlob *[]byte) (error) {

	serialNum := (*info)["serialNumberInserv"]

	tenants := (*info)["authorized"].(map[string]interface{})["tenants"]

	for _, tenant := range tenants.([]interface{}){
		tenant := tenant.(string)
		err := db.Query("UPDATE latest_snapshots_by_tenant " +
						"SET snapshot = textAsBlob(?)" +
						"WHERE tenant = ? AND serial_number = ?",
						*jsonBlob, tenant, serialNum).Exec()

		if err != nil { log.Fatal(err) }
	}

	return nil
}


func (db *DatabaseSession) insertSnapshotsBySerialNumber(info *map[string]interface{}, jsonBlob *[]byte) (error) {

	serialNum := (*info)["serialNumberInserv"]

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