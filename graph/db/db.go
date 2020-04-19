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
		(*model.Usecase)(nil), (*model.Case)(nil), (*model.Jotter)(nil),
		(*model.Organization)(nil)} {
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

	db := pg.Connect(&pg.Options{
		User:            "postgres",
		Password:        "postgres",
		Addr:            "localhost:5432",
		Database:        "usecase-database",
		ApplicationName: "Usecase-server",
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
