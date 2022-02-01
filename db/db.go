package db

import (
	"errors"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ErrDbRecordNotFound = errors.New("db record not found")

type TodoAppDb struct {
	*gorm.DB
}

var Db TodoAppDb

func Init(connectionString string) {

	if len(connectionString) == 0 {
		log.Fatal("Connection string is empty")
	}

	log.Println("Initializing database connection. ConnectionString: " + connectionString)

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})

	if err != nil {
		log.Fatalln(err)
		return
	}

	Db = TodoAppDb{db}
}
