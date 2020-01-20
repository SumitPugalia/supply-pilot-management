package sqlite

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

func openConn(file string) sqlbuilder.Database {
	connSettings := sqlite.ConnectionURL{
		Database: file,
		Options:  nil,
	}

	conn, err := sqlite.Open(connSettings)

	if err != nil {
		panic("SHIT NO DB")
	}

	return conn
}

func getReadConn() sqlbuilder.Database {
	return openConn("/Users/sumit/Workspace/CAFU/supply/pilot-management/repository/impl/sqlite/db")
}

func getWriteConn() sqlbuilder.Database {
	return openConn("/Users/sumit/Workspace/CAFU/supply/pilot-management/repository/impl/sqlite/db")
}
