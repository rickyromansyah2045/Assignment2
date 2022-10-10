package databases

import (
	"assignment-2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	var host = "localhost"
	var port = "5432"
	var username = "postgres"
	var password = "Noerhick_02"
	var dbname = "order_by"

	var conn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	db, err := gorm.Open(postgres.Open(conn))

	if err != nil {
		fmt.Println("Error open connection db", err)

		return Database{}, err
	}

	err = db.Debug().AutoMigrate(models.Order{}, models.Item{})

	if err != nil {
		fmt.Println("Error On migration", err)

		return Database{}, err
	}

	return Database{
		db: db,
	}, nil
}
