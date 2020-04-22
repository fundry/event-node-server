package db

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"log"

	"github.com/vickywane/event-server/graph/model"
)

//Todo : Use & Load env vars here

func createSchema(db *pg.DB) error {
	for _, models := range []interface{}{(*model.User)(nil),
		(*model.User)(nil), (*model.Event)(nil), (*model.Preference)(nil)} {
		err := db.CreateTable(models, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			panic(err)
		} else {
			log.Println(err)
		}
	}
	return nil
}

func Connect() *pg.DB {
	log.Println("Db connection is starting")
	// APP_NAME := os.Getenv("APPLICATION_NAME" )
	// DB_USER := os.Getenv("POSTGRES_USER" )
	// DB_PASSWORD := os.Getenv("POSTGRES_DB_PASSWORD" )
	// DB_ADDRESS := os.Getenv("POSTGRES_DB_ADDRESS" )
	// DB_DATABASE := os.Getenv("POSTGRES_DB" )


	db := pg.Connect(&pg.Options{
		User:            "postgres",
		Password:        "postgres",
		Addr:            "localhost:5432",
		Database:        "event-database",
		ApplicationName: "event-server",
	})

	if db != nil {
		fmt.Println(db)
	}

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return db
}
