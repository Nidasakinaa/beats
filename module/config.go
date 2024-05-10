package module

import (
	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "ATS",
}

var MongoConn = atdb.MongoConnect(MongoInfo)
