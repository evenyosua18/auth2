package main

import (
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/utils/db"
	"github.com/evenyosua18/auth2/cmd/seeder/seeds"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// 1: seeder name

const (
	OauthClient = "oauth_client"
)

func main() {
	//setup environment variable (for local)
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	seeder := ""

	//get argument
	args := os.Args

	if len(args) == 2 {
		log.Println("run seeder with argument ", args[1])
		seeder = args[1]
	}

	// initialize connection mongodb
	con := &db.MongoConnection{
		DatabaseName: os.Getenv(constant.MongoDB),
		Uri:          os.Getenv(constant.MongoUri),
		AppName:      os.Getenv(constant.AppName),
	}

	con.Connect()

	// run seeder
	switch seeder {
	case OauthClient:
		log.Println("run seed " + OauthClient)
		if err := seeds.GenerateOauthClient(con); err != nil {
			panic(err)
		}
		break
	default:
		log.Println("run default seed")
		break
	}
}
