package db

import (
    "fmt"
    "github.com/go-pg/pg/v9"
    "github.com/go-pg/pg/v9/orm"
    "github.com/joho/godotenv"
    "github.com/satori/go.uuid"

    "github.com/vickywane/event-server/graph/model"
)

func createSchema(db *pg.DB) error {
    for _, models := range []interface{}{(*model.User)(nil),
        (*model.User)(nil), (*model.Event)(nil), (*model.Preference)(nil),
        (*model.File)(nil), (*model.Team)(nil), (*model.Sponsor)(nil),
        (*model.Tasks)(nil), (*model.Tracks)(nil), (*model.Talk)(nil)} {
        err := db.CreateTable(models, &orm.CreateTableOptions{
            IfNotExists: true, FKConstraints: true,
        })
        if err != nil {
            panic(err)
        }
    }
    return nil
}

func Connect() *pg.DB {
    godotenv.Load(".env")

    Envs, err := godotenv.Read(".env")

    _ = uuid.NewV4()
    db := pg.Connect(&pg.Options{
        User:            Envs["POSTGRES_USER"],
        Password:        Envs["POSTGRES_DB_PASSWORD"],
        Addr:            Envs["POSTGRES_DB_ADDRESS"],
        Database:        Envs["POSTGRES_DB"],
        ApplicationName: Envs["APPLICATION_NAME"],
    })

    if db != nil {
        fmt.Println(db)
    }

    error := createSchema(db)
    if error != nil {
        panic(err)
    }

    // NOTE!! this func might likely try to reseed on every restart
    SeedDatabase(db)

    return db
}
