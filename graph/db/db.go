package db

import (
    "github.com/go-pg/pg/v9"
    "github.com/go-pg/pg/v9/orm"
    "github.com/joho/godotenv"
	"os"

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
        (*model.UserFile)(nil), (*model.EventFile)(nil), (*model.Team)(nil),
        (*model.Tasks)(nil), (*model.Tracks)(nil), (*model.Talk)(nil),
        (*model.Volunteer)(nil), (*model.BetaTester)(nil), (*model.Attendee)(nil),
        (*model.Category)(nil), (*model.CartItem)(nil), (*model.Purchases)(nil),
        (*model.TaskComments)(nil), (*model.MeetupGroups)(nil), (*model.EventTalk)(nil),
        (*model.BugReport)(nil), (*model.FeatureRequest)(nil), (*model.Sponsor)(nil),
        (*model.Reminder)(nil),} {
        if err := db.CreateTable(models, &orm.CreateTableOptions{
            IfNotExists: true, FKConstraints: false, // Todo: turned this off because of VOLUNTEER table. Check out later!!
        }); err != nil {
            panic(err)
        }
    }
    return nil
}

func Connect() *pg.DB {
    godotenv.Load(".env")
    Envs, err := godotenv.Read(".env")

    db := pg.Connect(&pg.Options{
        Password:        os.GetEnv("PROD_POSTGRES_DB_PASSWORD"),
        User:            os.GetEnv("PROD_POSTGRES_USER"),
        Addr:            os.GetEnv("PROD_POSTGRES_DB_ADDRESS"),
        Database:        os.GetEnv("PROD_POSTGRES_DB"),
        ApplicationName: os.GetEnv"PROD_APPLICATION_NAME"),
        MaxRetries:      10,
        TLSConfig:       nil,
    })

    // ========> USING CONN STR
    // opt, err := fmt.SprintF("pg.ParseURL("postgres://%v:pass@%v/%v")", Addr, Usr, Db )
    // if err != nil {
    //     panic(err)
    // }
    // db := pg.Connect(opt)

    if schemaErr := createSchema(db); schemaErr != nil {
        panic(err)
    }

    SeedDatabase(db)
    // NOTE!! this func might likely try to reseed on every restart

    if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
        panic("PostgreSQL is down")
    }

    return db
}
