package middlewares

import "breeze/internal/db"

var dbClient *db.Db = nil

func getDb() *db.Db {
	if dbClient == nil {
		dbClient = db.NewDb()
	}
	return dbClient
}
