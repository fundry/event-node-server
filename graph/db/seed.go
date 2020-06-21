package db

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"math/rand"
	"time"

	"github.com/vickywane/event-server/graph/model"
)

// this func helps me seed the db each time its dropped.
// Todo: Use a proper SQL file to seed.

func SeedDatabase(db *pg.DB) {
	fmt.Println("Seeding Database ....")
	user := model.User{
		ID:         rand.Int(),
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
		// EventDate:        []*string{"12 - 12 - 12"},
		IsArchived:  false,
		IsLocked:    false,
		Attendees:   nil,
		// UserID:   nil,
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

	sponsor := model.Sponsor{
		ID:             rand.Int(),
		Name:           "John&sons.co",
		Type:           "Platinum",
		Amount:         1000,
		Event:          nil,
		IsOrganization: false,
	}

	task := model.Tasks{
		ID:          rand.Int(),
		Name:        "Sketch and Create Event Media Assets",
		Category:        "Design",
		IsCompleted: false,
		Assignees:   nil,
	}

	talk := model.Talk{
		ID:           rand.Int(),
		Title:        "Building Modern Distributed Systems",
		TalkCoverURI: nil,
		Summary:      "Learn about Distributed systems and how they're built ",
		Description:  "Come learn how building modern distributed systems can affect performance of a Software",
		Archived:     false,
		Duration:     "5 minutes",
		Tags:         nil,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	track := model.Tracks{
		ID:          rand.Int(),
		Name:        "Design Track",
		Duration:    "10am - 11pm",
		TotalTalks:  2,
		IsCompleted: false,
		Archived:    false,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	betaTester := model.BetaTester{
		ID:          rand.Int(),
		Name:        "John Doe",
		Email:       "johndoe@gmail.com",
		DateApplied: time.Now().Format("01-02-2006"),
	}

	db.Insert(&user)
	db.Insert(&event)
	db.Insert(&preference)
	db.Insert(&team)
	db.Insert(&sponsor)
	db.Insert(&task)
	db.Insert(&talk)
	db.Insert(&track)
	db.Insert(&betaTester)
}
