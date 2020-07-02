package db

import (
    "fmt"
    "github.com/go-pg/pg/v9"
    "github.com/go-pg/pg/v9/orm"
    "github.com/joho/godotenv"

    "github.com/vickywane/event-server/graph/model"
)

// Hierarchy of execution of functions here is important.
// Wrong placement panics the system
// --> Establishing db connection comes before creating ORM models
// ---> Seeding the db comes next
// -----> Checking db health last!

func createSchema(db *pg.DB) error {
    for _, models := range []interface{}{(*model.User)(nil),
        (*model.User)(nil), (*model.Event)(nil), (*model.Preference)(nil),
        (*model.UserFile)(nil), (*model.EventFile)(nil), (*model.Team)(nil), (*model.Sponsor)(nil),
        (*model.Tasks)(nil), (*model.Tracks)(nil), (*model.Talk)(nil),
        (*model.Volunteer)(nil), (*model.BetaTester)(nil), (*model.Attendee)(nil),
        (*model.Category)(nil), (*model.CartItem)(nil), (*model.Purchases)(nil),
        (*model.TaskComments)(nil), (*model.MeetupGroups)(nil), (*model.EventTalk)(nil)} {
        err := db.CreateTable(models, &orm.CreateTableOptions{
            IfNotExists: true, FKConstraints: false, // turned this off because of VOLUNTEER table. Check out later!!
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

    db := pg.Connect(&pg.Options{
        Password:        Envs["POSTGRES_DB_PASSWORD"],
        User:            Envs["POSTGRES_USER"],
        Addr:            Envs["POSTGRES_DB_ADDRESS"],
        Database:        Envs["POSTGRES_DB"],
        ApplicationName: Envs["APPLICATION_NAME"],
        MaxRetries: 10,
        TLSConfig: nil,
    })

    // ========> USING CONN STR
    // opt, err := pg.ParseURL("postgres://user:pass@localhost:5432/db_name")
    // if err != nil {
    //     panic(err)
    // }
    //
    // db := pg.Connect(opt)


    if db != nil {
        fmt.Println(db)
    }

    error := createSchema(db)
    if error != nil {
        panic(err)
    }
    SeedDatabase(db)

    if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
        panic("PostgreSQL is down")
    }
    // NOTE!! this func might likely try to reseed on every restart

    return db
}
