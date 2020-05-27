// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Attendee struct {
	ID         int      `json:"id"`
	DateJoined string   `json:"dateJoined"`
	User       []*User  `json:"user"`
	UserID     int      `json:"user_id"`
	Event      []*Event `json:"event"`
	EventID    int      `json:"event_id"`
}

type AuthResponse struct {
	ID        int       `json:"id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
	User      *User     `json:"user"`
}

type CreateAttendee struct {
	User  []*CreateUser  `json:"user"`
	Event []*CreateEvent `json:"event"`
}

type CreateEvent struct {
	Name         string             `json:"name"`
	Summary      string             `json:"summary"`
	Alias        string             `json:"alias"`
	Date         int                `json:"Date"`
	Description  string             `json:"description"`
	Website      string             `json:"website"`
	Email        string             `json:"Email"`
	EventType    string             `json:"eventType"`
	IsArchived   *bool              `json:"isArchived"`
	IsLocked     *bool              `json:"isLocked"`
	CreatedBy    *CreateUser        `json:"CreatedBy"`
	Attendees    []*CreateUser      `json:"attendees"`
	Venue        string             `json:"venue"`
	VolunteerID  *int               `json:"volunteer_id"`
	Volunteering []*CreateVolunteer `json:"volunteering"`
}

type CreatePreference struct {
	Name  string       `json:"name"`
	Color string       `json:"color"`
	Event *CreateEvent `json:"Event"`
}

type CreateSponsor struct {
	Name           string       `json:"name"`
	Type           *string      `json:"type"`
	Amount         *int         `json:"amount"`
	Event          *CreateEvent `json:"event"`
	IsOrganization *bool        `json:"isOrganization"`
}

type CreateTalk struct {
	Title        string    `json:"title"`
	TalkCoverURI *string   `json:"talkCoverUri"`
	Summary      string    `json:"summary"`
	Description  string    `json:"description"`
	Archived     bool      `json:"Archived"`
	Duration     int       `json:"duration"`
	Tags         []*string `json:"tags"`
}

type CreateTasks struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	IsCompleted bool          `json:"isCompleted"`
	Event       *CreateEvent  `json:"event"`
	Assignees   []*CreateUser `json:"assignees"`
	CreatedBy   *CreateUser   `json:"createdBy"`
	TeamID      int           `json:"team_id"`
}

type CreateTeam struct {
	Name    string        `json:"name"`
	Members []*CreateUser `json:"members"`
	Goal    string        `json:"goal"`
}

type CreateTrack struct {
	Name        string  `json:"name"`
	TrackImgURI *string `json:"trackImgUri"`
	Duration    string  `json:"duration"`
	Talks       *string `json:"talks"`
	TotalTalks  int     `json:"totalTalks"`
	CreatedBy   *string `json:"createdBy"`
	IsCompleted bool    `json:"isCompleted"`
	Archived    bool    `json:"Archived"`
}

type CreateUser struct {
	Name         string             `json:"name"`
	Role         *string            `json:"role"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	Volunteering []*CreateVolunteer `json:"volunteering"`
	Events       []*CreateEvent     `json:"events"`
}

type CreateVolunteer struct {
	Role              string       `json:"role"`
	Duration          *string      `json:"duration"`
	User              *CreateUser  `json:"user"`
	Event             *CreateEvent `json:"event"`
	VolunteerProposal string       `json:"volunteer_proposal"`
}

type DeleteFile struct {
	ID  int     `json:"id"`
	URI *string `json:"uri"`
}

type Event struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Summary        string       `json:"summary"`
	Alias          string       `json:"alias"`
	Email          string       `json:"Email"`
	ConfirmedEmail bool         `json:"confirmedEmail"`
	Website        string       `json:"website"`
	BucketLink     string       `json:"bucketLink"`
	Venue          string       `json:"venue"`
	EventType      string       `json:"eventType"`
	Date           int          `json:"Date"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
	AuthorID       int          `json:"author_id"`
	CreatedBy      []*User      `json:"createdBy"`
	Attendees      []*User      `json:"attendees"`
	Tracks         []*Tracks    `json:"tracks"`
	TrackID        *int         `json:"track_id"`
	Teams          []*Team      `json:"teams"`
	VolunteerID    *int         `json:"volunteer_id"`
	Volunteer      []*Volunteer `json:"volunteer"`
	TotalAttendees int          `json:"totalAttendees"`
	IsArchived     bool         `json:"isArchived"`
	IsLocked       bool         `json:"isLocked"`
}

type File struct {
	ID        int            `json:"id"`
	File      graphql.Upload `json:"file"`
	FileURI   string         `json:"file_uri"`
	Timestamp time.Time      `json:"timestamp"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Preference struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Event     *Event    `json:"Event"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Sponsor struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Amount         int    `json:"amount"`
	Event          *Event `json:"event"`
	IsOrganization bool   `json:"isOrganization"`
}

type Talk struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	TalkCoverURI *string   `json:"talkCoverUri"`
	Duration     int       `json:"duration"`
	Summary      string    `json:"summary"`
	Description  string    `json:"description"`
	Archived     bool      `json:"Archived"`
	Tags         []*string `json:"tags"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	EventID      *int      `json:"event_id"`
	ReviewersID  int       `json:"reviewers_id"`
	SpeakerID    int       `json:"speaker_id"`
	Event        []*Event  `json:"event"`
	Speaker      []*User   `json:"speaker"`
	Reviewers    []*User   `json:"reviewers"`
}

type Tasks struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	IsCompleted bool      `json:"isCompleted"`
	Assignees   []*User   `json:"assignees"`
	CreatedBy   *User     `json:"createdBy"`
	Event       *Event    `json:"event"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	TeamID      int       `json:"team_id"`
}

type Team struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Members   []*User   `json:"members"`
	Goal      string    `json:"goal"`
	CreatedBy []*Event  `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	EventID   int       `json:"event_id"`
	Event     []*Event  `json:"event"`
}

type Tracks struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	TrackImgURI string    `json:"trackImgUri"`
	Duration    string    `json:"duration"`
	TotalTalks  int       `json:"totalTalks"`
	IsCompleted bool      `json:"isCompleted"`
	Archived    bool      `json:"Archived"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	EventID     int       `json:"event_id"`
	Talks       []*Talk   `json:"talks"`
	CreatedBy   []*Event  `json:"createdBy"`
}

type UpdateAttendee struct {
	User []*CreateUser `json:"user"`
}

type UpdateEvent struct {
	Name           *string            `json:"name"`
	Type           *string            `json:"type"`
	Summary        *string            `json:"summary"`
	Alias          *string            `json:"alias"`
	BucketLink     *string            `json:"bucketLink"`
	Description    *string            `json:"description"`
	EventType      *string            `json:"eventType"`
	IsArchived     *bool              `json:"isArchived"`
	TotalAttendees *int               `json:"totalAttendees"`
	IsLocked       *bool              `json:"isLocked"`
	Email          *string            `json:"Email"`
	Website        *string            `json:"website"`
	TrackID        *int               `json:"track_id"`
	UpdatedAt      *time.Time         `json:"updatedAt"`
	Attendees      []*CreateUser      `json:"attendees"`
	Venue          *string            `json:"venue"`
	Date           int                `json:"Date"`
	Team           *CreateTeam        `json:"team"`
	VolunteerID    *int               `json:"volunteer_id"`
	Volunteering   []*CreateVolunteer `json:"volunteering"`
}

type UpdatePreference struct {
	Name      string       `json:"name"`
	Color     string       `json:"color"`
	Event     *CreateEvent `json:"Event"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

type UpdateSponsor struct {
	Name           string       `json:"name"`
	Type           *string      `json:"type"`
	Amount         *string      `json:"amount"`
	Event          *CreateEvent `json:"event"`
	IsOrganization *bool        `json:"isOrganization"`
}

type UpdateTalk struct {
	Title        string        `json:"title"`
	TalkCoverURI *string       `json:"talkCoverUri"`
	Summary      string        `json:"summary"`
	Description  string        `json:"description"`
	Reviewers    []*CreateUser `json:"reviewers"`
	Archived     bool          `json:"Archived"`
	Duration     int           `json:"duration"`
	Tags         []*string     `json:"tags"`
}

type UpdateTask struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	IsCompleted bool          `json:"isCompleted"`
	Assignees   []*CreateUser `json:"assignees"`
}

type UpdateTeam struct {
	Name    string        `json:"name"`
	Members []*CreateUser `json:"members"`
	Goal    string        `json:"goal"`
}

type UpdateTrack struct {
	Name        string    `json:"name"`
	TrackImgURI *string   `json:"trackImgUri"`
	Duration    string    `json:"duration"`
	TotalTalks  int       `json:"totalTalks"`
	IsCompleted bool      `json:"isCompleted"`
	Archived    bool      `json:"Archived"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UpdateUser struct {
	Name         *string            `json:"name"`
	Role         *string            `json:"role"`
	Email        *string            `json:"email"`
	Password     *string            `json:"password"`
	Volunteering []*CreateVolunteer `json:"volunteering"`
	Events       []*CreateEvent     `json:"events"`
	UpdatedAt    *time.Time         `json:"updatedAt"`
}

type UpdateVolunteer struct {
	Role              *string     `json:"role"`
	ApprovalStatus    string      `json:"approvalStatus"`
	Duration          *string     `json:"duration"`
	Team              *CreateTeam `json:"team"`
	TeamID            *int        `json:"team_id"`
	VolunteerProposal string      `json:"volunteer_proposal"`
}

type UploadFile struct {
	File    graphql.Upload `json:"file"`
	FileURI string         `json:"file_uri"`
}

type User struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Role         *string      `json:"role"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	BucketLink   string       `json:"bucketLink"`
	Talks        []*Talk      `json:"talks"`
	Events       []*Event     `json:"events"`
	VolunteerID  int          `json:"volunteer_id"`
	Volunteering []*Volunteer `json:"volunteering"`
	EventID      int          `json:"event_id"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
}

type Volunteer struct {
	ID                int      `json:"id"`
	Role              string   `json:"role"`
	Duration          string   `json:"duration"`
	ApprovalStatus    string   `json:"approvalStatus"`
	VolunteerProposal string   `json:"volunteer_proposal"`
	DateApplied       string   `json:"dateApplied"`
	Team              []*Team  `json:"team"`
	TeamID            *int     `json:"team_id"`
	Event             []*Event `json:"event"`
	EventID           int      `json:"event_id"`
	User              []*User  `json:"user"`
	UserID            int      `json:"user_id"`
}
