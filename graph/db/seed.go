package db

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"math/rand"

	"github.com/vickywane/event-server/graph/model"
)

// this func helps me seed the db each time its dropped.
// Todo: Use a proper SQL file to seed.

func SeedDatabase(db *pg.DB) {

	fmt.Println("Seeding Database ....")
	user := model.User{
		ID:         rand.Int(),
		Name:       "John Doe",
		Email:      "Johndoe@gmail.com",
		BucketLink: "https://google.cloud.com/storage",
	}

	event := model.Event{
		ID:          rand.Int(),
		Name:        "Test Event",
		Description: "A test event to seed the database",
		Summary:     "A test event to seed the database",
		Alias:       "TE",
		Email:       "test@gmail.com",
		Website:     "test.com",
		BucketLink:  "test.com/storage",
		Venue:       "Test, Lagos, Nigeria",
		EventType:   "Conference",
		Date:        12 - 12 - 12,
		IsArchived:  false,
		IsLocked:    false,
		Attendees:   nil,
		CreatedBy:   nil,
	}

	preference := model.Preference{
		ID:    rand.Int(),
		Name:  "Preference 1",
		Color: "Yellow",
		Event: nil,
	}

	team := model.Team{
		ID:        rand.Int(),
		Name:      "Technical Team",
		Members:   nil,
		Goal:      "To make sure test team works",
		CreatedBy: nil,
	}

	file := model.File{
		ID:       rand.Int(),
		Filename: "Image",
		Mimetype: "img",
		Size:     "240gb",
		URI:      "https://test.cloud.com",
		Encoding: "UTf8",
	}

	db.Insert(&user)
	db.Insert(&event)
	db.Insert(&preference)
	db.Insert(&team)
	db.Insert(&file)

}
