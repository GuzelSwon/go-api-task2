package src

import (
	"app/src/logger"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	username          = os.Getenv("DBUSER")
	hostname          = os.Getenv("DBHOST")
	password          = os.Getenv("DBPASS")
	dbname            = os.Getenv("DBNAME")
	port              = os.Getenv("SERVER_PORT")
	callEntitiesDbDsn = EntitiesDbDsn
)

type EntitiesRepo struct {
	Db *gorm.DB
}

func EntitiesDbDsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, hostname, port, dbname)
}

func connectDB() *gorm.DB {
	dbDsn := fmt.Sprintf("postgres://%s:%s@%s:%s/", username, password, hostname, port)

	db := openDB(dbDsn)
	checkIfDbExists(db)
	entitiesDb := openDB(callEntitiesDbDsn())
	addOtelPlugin(entitiesDb)
	return entitiesDb
}

func addOtelPlugin(db *gorm.DB) {
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		logger.Info("Error adding OpenTelemetryPlugin : error="+err.Error(),
			logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryDatabase})
	}
}

func openDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Info("Error connecting to database: error="+err.Error(),
			logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryDatabase})
		return nil
	}
	return db
}

func checkIfDbExists(db *gorm.DB) {
	db.Exec("CREATE DATABASE IF NOT EXISTS entities;")
	db.Exec("USE entities;")
	db.Exec("CREATE TABLE IF NOT EXISTS entities(id INT NOT NULL,slug VARCHAR(30) NOT NULL,url VARCHAR(60) NOT NULL,title VARCHAR(80)  NOT NULL,content VARCHAR(1000)  NOT NULL,image VARCHAR(60)  NOT NULL,thumbnail VARCHAR(60)  NOT NULL,status VARCHAR(60)  NOT NULL,category VARCHAR(60)  NOT NULL,published_at VARCHAR(60)  NOT NULL,updated_at VARCHAR(60)  NOT NULL,user_id VARCHAR(60)  NOT NULL,PRIMARY KEY(id));")
}

func CreateNewEntitiesRepo() *EntitiesRepo {
	database := connectDB()
	return &EntitiesRepo{Db: database}
}
