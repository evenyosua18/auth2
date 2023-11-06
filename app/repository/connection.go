package repository

import (
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/utils/db"
	"os"
)

var (
	Con *Connection
)

type Connection struct {
	MainMongoDB *db.MongoConnection
}

func InitConnection() {
	//initialize connection
	Con = &Connection{}

	//initialize connection mongodb
	Con.MainMongoDB = &db.MongoConnection{
		DatabaseName: os.Getenv(constant.MongoDB),
		Uri:          os.Getenv(constant.MongoUri),
		AppName:      os.Getenv(constant.AppName),
	}

	Con.MainMongoDB.Connect()
}
