package main

import (
	"emailing-service/api"
	"emailing-service/database"
	"emailing-service/helpers"
	"log"
	"os"
)

// @title Emailing-Service
// @version 1.0
// @contact.name API Support
// @contact.email software@upcraft.li

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Access-Token

func main() {
	status := database.InitTables(helpers.DB)
	if !status {
		log.Println("[SETUP GORM]: an error has occurred with a table.")
		os.Exit(1)
	}
	api.Router()
}

func init() {
	helpers.DB = database.Client()
}
