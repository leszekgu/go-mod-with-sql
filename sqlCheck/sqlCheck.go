package sqlCheck

import (
	"database/sql"
	"log"
)

func MySqlCheck() string {
	return "This is a test of module working v1.0.4"
}

func PgConnectAndClose(driverName, connectionString string) {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connection went well")
	}
	defer db.Close()
}

