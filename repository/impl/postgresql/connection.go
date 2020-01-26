package postgresql

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

func openConn() sqlbuilder.Database {
	connSettings := postgresql.ConnectionURL{
		User:     "postgres",
		Password: "postgres",
		Host:     "localhost",
		Socket:   "",
		Database: "pilot_management",
		Options:  nil,
	}

	conn, err := postgresql.Open(connSettings)

	if err != nil {
		panic("connection could not be setup with the database")
	}

	return conn
}

func getReadConn() sqlbuilder.Database {
	return openConn()
}

func getWriteConn() sqlbuilder.Database {
	return openConn()
}
