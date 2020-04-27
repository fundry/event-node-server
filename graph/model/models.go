// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type AuthResponse struct {
	Token     string  `json:"token"`
	ExpiredAt *string `json:"expiredAt"`
}

type CreateEvent struct {
	Name        string        `json:"name"`
	Summary     string        `json:"summary"`
	Alias       string        `json:"alias"`
	Description string        `json:"description"`
	Website     string        `json:"website"`
	Email       string        `json:"Email"`
	EventType   string        `json:"eventType"`
	IsArchived  *bool         `json:"isArchived"`
	IsLocked    *bool         `json:"isLocked"`
	Attendees   []*CreateUser `json:"attendees"`
	Venue       string        `json:"venue"`
	Date        int           `json:"Date"`
}

type CreateFile struct {
	ID        int       `json:"id"`
	Filename  string    `json:"filename"`
	Mimetype  string    `json:"mimetype"`
	Size      string    `json:"size"`
	URI       string    `json:"uri"`
	Encoding  string    `json:"encoding"`
	Timestamp time.Time `json:"timestamp"`
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

type CreateTasks struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	IsCompleted bool          `json:"isCompleted"`
	Assignees   []*CreateUser `json:"assignees"`
	CreatedBy   *CreateUser   `json:"createdBy"`
}

type CreateTeam struct {
	Name      string        `json:"name"`
	Members   []*CreateUser `json:"members"`
	Goal      string        `json:"goal"`
	CreatedBy *CreateUser   `json:"createdBy"`
}

type CreateUser struct {
	Name     string         `json:"name"`
	Role     *string        `json:"role"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Events   []*CreateEvent `json:"events"`
}

type DeleteFile struct {
	ID  int     `json:"id"`
	URI *string `json:"uri"`
}

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Summary     string    `json:"summary"`
	Alias       string    `json:"alias"`
	Email       string    `json:"Email"`
	Website     string    `json:"website"`
	BucketLink  string    `json:"bucketLink"`
	Venue       string    `json:"venue"`
	EventType   string    `json:"eventType"`
	Date        int       `json:"Date"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedBy   *User     `json:"createdBy"`
	Attendees   []*User   `json:"attendees"`
	Teams       []*Team   `json:"teams"`
	IsArchived  bool      `json:"isArchived"`
	IsLocked    bool      `json:"isLocked"`
}

type File struct {
	ID        int       `json:"id"`
	Filename  string    `json:"filename"`
	Mimetype  string    `json:"mimetype"`
	Size      string    `json:"size"`
	URI       string    `json:"uri"`
	Encoding  string    `json:"encoding"`
	Timestamp time.Time `json:"timestamp"`
}

type LoginUser struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Preference struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Event     *Event    `json:"Event"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Response struct {
	Completed bool `json:"completed"`
}

type Sponsor struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Amount         int    `json:"amount"`
	Event          *Event `json:"event"`
	IsOrganization bool   `json:"isOrganization"`
}

type Tasks struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	IsCompleted bool      `json:"isCompleted"`
	Assignees   []*User   `json:"assignees"`
	CreatedBy   *User     `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Team struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Members   []*User   `json:"members"`
	Goal      string    `json:"goal"`
	CreatedBy *User     `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateEvent struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Summary     string        `json:"summary"`
	Alias       string        `json:"alias"`
	BucketLink  string        `json:"bucketLink"`
	Description string        `json:"description"`
	EventType   string        `json:"eventType"`
	IsArchived  *bool         `json:"isArchived"`
	IsLocked    *bool         `json:"isLocked"`
	Email       string        `json:"Email"`
	Website     string        `json:"website"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Attendees   []*CreateUser `json:"attendees"`
	Venue       string        `json:"venue"`
	Date        int           `json:"Date"`
	Team        *CreateTeam   `json:"team"`
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

type UpdateTask struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	IsCompleted bool          `json:"isCompleted"`
	Assignees   []*CreateUser `json:"assignees"`
}

type UpdateTeam struct {
	Name      string        `json:"name"`
	Members   []*CreateUser `json:"members"`
	Goal      string        `json:"goal"`
	CreatedBy *CreateUser   `json:"createdBy"`
}

type UpdateUser struct {
	Name      *string        `json:"name"`
	Role      *string        `json:"role"`
	Email     *string        `json:"email"`
	Password  *string        `json:"password"`
	Events    []*CreateEvent `json:"events"`
	UpdatedAt *time.Time     `json:"updatedAt"`
}

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Role       *string   `json:"role"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	BucketLink string    `json:"bucketLink"`
	Events     []*Event  `json:"events"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
