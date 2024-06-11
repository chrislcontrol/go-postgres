package db

import (
	"fmt"

	"github.com/chrislcontrol/go-postgres/internal/entity"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	dbname    = "postgres"
	password  = "postgres"
	sslmode   = "disable"
	timezone  = "UTC"
)

var dbSessionCache *gorm.DB = nil

func GetDBSession() *gorm.DB {
	if dbSessionCache != nil {
		return dbSessionCache
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	dbSessionCache = db

	return db
}


func StartDBSession() {
	GetDBSession()
}
