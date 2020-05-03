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

	file := model.File{
		ID:       rand.Int(),
		Filename: "Image",
		Mimetype: "img",
		Size:     "240gb",
		URI:      "https://test.cloud.com",
		Encoding: "UTf8",
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
		Type:        "Design",
		IsCompleted: false,
		Assignees:   nil,
	}

	talk := model.Talk{
		ID:           rand.Int(),
		Title:        "Building Modern Distributed Systems",
		TalkCoverURI: nil,
		Summary:      "Learn about Distributed systems and how they're built ",
		Description:  "Come learn how building modern distributed systems can affect performance of a Software",
		Reviewers:    nil,
		Archived:     false,
		Duration:     20,
		Tags:         nil,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	track := model.Track{
		ID:          rand.Int(),
		Name:        "Design Track",
		Duration:    "10am - 11pm",
		TotalTalks:  2,
		IsCompleted: false,
		Archived:    false,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	db.Insert(&user)
	db.Insert(&event)
	db.Insert(&preference)
	db.Insert(&team)
	db.Insert(&file)
	db.Insert(&sponsor)
	db.Insert(&task)
	db.Insert(&talk)
	db.Insert(&track)
}
