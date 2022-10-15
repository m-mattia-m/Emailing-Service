package database

import (
	"emailing-service/helpers"
	"emailing-service/models"
	"errors"
	"fmt"
	"log"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Set this variabels in the init-Func from the enviroment
var (
	host     = ""
	port     = ""
	username = ""
	password = ""
	name     = ""
)

func Client() *gorm.DB {
	config := mysqlDriver.Config{
		User:   username,
		Passwd: password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", host, port),
		DBName: name,
		Params: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "True",
			"loc":       "Local",
		},
		AllowNativePasswords: true,
	}
	dsn := config.FormatDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(dsn)
		log.Println("[SETUP GORM]: can not open the connection to the database\t -> Error: " + err.Error())
		return nil
	}
	return db
}

func InitTables(db *gorm.DB) bool {
	err := db.AutoMigrate(&models.Sender{})
	if err != nil {
		log.Println("[SETUP GORM]: create Table was failed \t\t -> Error: " + err.Error())
		return false
	}
	return true
}

func init() {
	log.Println("[SETUP GORM]: Status -> Grom-init is loaded")
	host, _ = helpers.GetEnv("DB_HOST")
	port, _ = helpers.GetEnv("DB_PORT")
	username, _ = helpers.GetEnv("DB_USERNAME")
	password, _ = helpers.GetEnv("DB_PASSWORD")
	name, _ = helpers.GetEnv("DB_NAME")
}

func Auth(APIKey string) (bool, error) {
	var sender models.Sender
	err := helpers.DB.First(&sender, "token = ?", APIKey).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("sender not found")
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
