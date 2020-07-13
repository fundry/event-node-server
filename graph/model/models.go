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

type BugReport struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	User        []*User  `json:"user"`
	UserID      int      `json:"user_id"`
	Event       []*Event `json:"event"`
	EventID     int      `json:"event_id"`
	Status      string   `json:"status"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

type CartItem struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CreatedAt   string      `json:"createdAt"`
	Category    []*Category `json:"category"`
	CategoryID  int         `json:"category_id"`
	Quantity    int         `json:"quantity"`
	Price       *string     `json:"price"`
	IsFree      bool        `json:"isFree"`
}

type Category struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Event   []*Event    `json:"event"`
	EventID int         `json:"event_id"`
	Items   []*CartItem `json:"items"`
	ItemID  *int        `json:"item_id"`
}

type CreateAttendee struct {
	User  []*CreateUser  `json:"user"`
	Event []*CreateEvent `json:"event"`
}

type CreateBugReport struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateCartItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       *string `json:"price"`
	IsFree      bool    `json:"isFree"`
}

type CreateCategory struct {
	Name   string `json:"name"`
	ItemID *int   `json:"item_id"`
}

type CreateEvent struct {
	Name                  string    `json:"name"`
	Summary               string    `json:"summary"`
	Alias                 string    `json:"alias"`
	Description           string    `json:"description"`
	Website               string    `json:"website"`
	Email                 string    `json:"Email"`
	EventType             string    `json:"eventType"`
	Venue                 string    `json:"venue"`
	SpeakerConduct        *string   `json:"speakerConduct"`
	VolunteerID           *int      `json:"volunteer_id"`
	SponsorID             *int      `json:"sponsor_id"`
	EventDate             []*string `json:"EventDate"`
	Actions               []string  `json:"actions"`
	MediaLinks            []string  `json:"mediaLinks"`
	IsVirtual             bool      `json:"isVirtual"`
	IsArchived            bool      `json:"isArchived"`
	IsLocked              bool      `json:"isLocked"`
	IsAcceptingVolunteers bool      `json:"isAcceptingVolunteers"`
	IsAcceptingTalks      bool      `json:"isAcceptingTalks"`
}

type CreateFeatureRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateMeetupGroup struct {
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Alias       string    `json:"alias"`
	Email       string    `json:"email"`
	Website     *string   `json:"website"`
	Description string    `json:"description"`
	MediaLinks  []*string `json:"mediaLinks"`
}

type CreateNote struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateReminder struct {
	Name string `json:"name"`
	From string `json:"from"`
	Due  string `json:"due"`
}

type CreateSponsor struct {
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	ImageURL       *string `json:"image_url"`
	IsOrganization bool    `json:"isOrganization"`
}

type CreateTalk struct {
	Title        string    `json:"title"`
	TalkCoverURI *string   `json:"talkCoverUri"`
	Summary      string    `json:"summary"`
	Description  string    `json:"description"`
	Archived     bool      `json:"Archived"`
	Duration     string    `json:"duration"`
	Tags         []*string `json:"tags"`
}

type CreateTaskComment struct {
	Content   string `json:"content"`
	WrittenAt string `json:"writtenAt"`
	UserID    int    `json:"user_id"`
}

type CreateTasks struct {
	Name      string        `json:"name"`
	Category  string        `json:"category"`
	Status    string        `json:"status"`
	Priority  string        `json:"priority"`
	Assignees []*CreateUser `json:"assignees"`
	TeamID    int           `json:"team_id"`
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
	Summary     string  `json:"summary"`
	Talks       *string `json:"talks"`
	TotalTalks  int     `json:"totalTalks"`
	CreatedBy   *string `json:"createdBy"`
	IsCompleted bool    `json:"isCompleted"`
	Archived    bool    `json:"Archived"`
}

type CreateUser struct {
	Name         string             `json:"name"`
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
	ID                    int              `json:"id"`
	Name                  string           `json:"name"`
	Description           string           `json:"description"`
	Summary               string           `json:"summary"`
	Alias                 string           `json:"alias"`
	Email                 string           `json:"Email"`
	Website               string           `json:"website"`
	BucketLink            string           `json:"bucketLink"`
	BucketName            string           `json:"bucketName"`
	Venue                 string           `json:"venue"`
	EventType             string           `json:"eventType"`
	MeetupGroupID         *int             `json:"meetupGroup_id"`
	CreatedAt             time.Time        `json:"createdAt"`
	DateCreated           string           `json:"dateCreated"`
	UpdatedAt             time.Time        `json:"updatedAt"`
	AuthorID              int              `json:"author_id"`
	TrackID               *int             `json:"track_id"`
	VolunteerID           *int             `json:"volunteer_id"`
	SponsorID             *int             `json:"sponsor_id"`
	Settings              []*EventSettings `json:"settings"`
	SettingsID            *int             `json:"settings_id"`
	TotalAttendees        int              `json:"totalAttendees"`
	MediaLinks            []string         `json:"mediaLinks"`
	Tracks                []*Tracks        `json:"tracks"`
	Talk                  []*EventTalk     `json:"talk"`
	MeetupGroups          []*MeetupGroups  `json:"meetupGroups"`
	EventDate             []*string        `json:"EventDate"`
	CreatedBy             []*User          `json:"createdBy"`
	Attendees             []*Attendee      `json:"attendees"`
	SpeakerConduct        *string          `json:"speakerConduct"`
	Actions               []string         `json:"actions"`
	Teams                 []*Team          `json:"teams"`
	CartItemsCategory     []*Category      `json:"cart_items_category"`
	Volunteer             []*Volunteer     `json:"volunteer"`
	Sponsors              []*Sponsor       `json:"sponsors"`
	IsVirtual             bool             `json:"isVirtual"`
	ConfirmedEmail        bool             `json:"confirmedEmail"`
	IsAcceptingVolunteers bool             `json:"isAcceptingVolunteers"`
	IsAcceptingAttendees  bool             `json:"isAcceptingAttendees"`
	IsAcceptingTalks      bool             `json:"isAcceptingTalks"`
	IsArchived            bool             `json:"isArchived"`
	IsLocked              bool             `json:"isLocked"`
}

type EventFile struct {
	ID         int            `json:"id"`
	File       graphql.Upload `json:"file"`
	Type       string         `json:"type"`
	Event      []*Event       `json:"event"`
	EventID    *int           `json:"eventId"`
	UserID     int            `json:"userId"`
	UploadedBy []*User        `json:"uploadedBy"`
	FileURI    string         `json:"file_uri"`
	Timestamp  time.Time      `json:"timestamp"`
}

type EventSettings struct {
	ID                          int     `json:"id"`
	EventID                     int     `json:"eventId"`
	ShowWelcomeMeetupGroup      bool    `json:"showWelcomeMeetupGroup"`
	ShowTeamInstruction         bool    `json:"showTeamInstruction"`
	ShowInvitationInstruction   bool    `json:"showInvitationInstruction"`
	ShowWelcomeEventInstruction bool    `json:"showWelcomeEventInstruction"`
	EventThemeColour            *string `json:"eventThemeColour"`
}

type EventTalk struct {
	ID            int      `json:"id"`
	IsAccepted    bool     `json:"isAccepted"`
	DateSubmitted string   `json:"dateSubmitted"`
	DateAccepted  string   `json:"dateAccepted"`
	ReviewerID    *int     `json:"reviewer_id"`
	Event         []*Event `json:"event"`
	EventID       int      `json:"event_id"`
	Comment       *string  `json:"comment"`
	Reviewer      []*User  `json:"reviewer"`
	Track         *string  `json:"track"`
	DraftID       int      `json:"draftId"`
	Draft         []*Talk  `json:"draft"`
}

type FeatureRequest struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	User        []*User  `json:"user"`
	UserID      int      `json:"user_id"`
	Event       []*Event `json:"event"`
	EventID     int      `json:"event_id"`
	Status      string   `json:"status"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MakePurchases struct {
	ItemID        int     `json:"item_id"`
	UserID        int     `json:"user_id"`
	EventID       int     `json:"event_id"`
	DatePurchased string  `json:"datePurchased"`
	Quantity      int     `json:"quantity"`
	Price         *string `json:"price"`
}

type MeetupGroups struct {
	ID            int        `json:"id"`
	ImgURI        *string    `json:"img_uri"`
	Name          string     `json:"name"`
	Summary       string     `json:"summary"`
	Email         string     `json:"email"`
	Website       *string    `json:"website"`
	Facilitators  []*User    `json:"facilitators"`
	FacilitatorID *int       `json:"facilitator_id"`
	SponsorID     *int       `json:"sponsor_id"`
	Event         []*Event   `json:"event"`
	EventID       int        `json:"event_id"`
	Lead          []*User    `json:"lead"`
	Description   string     `json:"description"`
	Sponsors      []*Sponsor `json:"sponsors"`
	MediaLinks    []*string  `json:"mediaLinks"`
	Actions       []*string  `json:"actions"`
	Members       []*User    `json:"members"`
	MembersID     *int       `json:"members_id"`
	LeadID        int        `json:"lead_id"`
	Location      string     `json:"location"`
	Alias         string     `json:"alias"`
	CreatedAt     string     `json:"createdAt"`
}

type Notes struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Talk    []*Talk `json:"talk"`
	TalkID  int     `json:"talk_id"`
}

type Purchases struct {
	ID            int         `json:"id"`
	Item          []*CartItem `json:"item"`
	ItemID        int         `json:"item_id"`
	User          []*User     `json:"user"`
	UserID        int         `json:"user_id"`
	Event         []*Event    `json:"event"`
	EventID       int         `json:"event_id"`
	DatePurchased string      `json:"datePurchased"`
	Quantity      int         `json:"quantity"`
	Price         *string     `json:"price"`
}

type Reminder struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	From      string  `json:"from"`
	Due       string  `json:"due"`
	User      []*User `json:"user"`
	UserID    int     `json:"user_id"`
	CreatedAt string  `json:"createdAt"`
}

type Sponsor struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	ImageURL          *string `json:"image_url"`
	Type              string  `json:"type"`
	Event             *Event  `json:"event"`
	EventID           int     `json:"event_id"`
	SponsorshipStatus string  `json:"sponsorshipStatus"`
	IsOrganization    bool    `json:"isOrganization"`
}

type SubmitEventTalk struct {
	IsAccepted bool `json:"isAccepted"`
}

type Talk struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	TalkCoverURI *string   `json:"talkCoverUri"`
	Duration     string    `json:"duration"`
	Summary      string    `json:"summary"`
	Description  string    `json:"description"`
	Archived     bool      `json:"Archived"`
	Tags         []*string `json:"tags"`
	Notes        []*Notes  `json:"notes"`
	CreatedAt    string    `json:"createdAt"`
	UpdatedAt    string    `json:"updatedAt"`
	EventID      *int      `json:"event_id"`
	SpeakerID    int       `json:"speaker_id"`
	Event        []*Event  `json:"event"`
	Speaker      []*User   `json:"speaker"`
}

type TaskComments struct {
	ID        int      `json:"id"`
	Content   string   `json:"content"`
	WrittenBy []*User  `json:"writtenBy"`
	Task      []*Tasks `json:"task"`
	UserID    int      `json:"user_id"`
	TaskID    int      `json:"task_id"`
	WrittenAt string   `json:"writtenAt"`
}

type Tasks struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Category  string          `json:"category"`
	Status    string          `json:"status"`
	Priority  string          `json:"priority"`
	Assignees []*User         `json:"assignees"`
	CreatedBy []*User         `json:"createdBy"`
	Event     []*Event        `json:"event"`
	CreatedAt string          `json:"createdAt"`
	Comments  []*TaskComments `json:"comments"`
	CommentID *int            `json:"comment_id"`
	UpdatedAt time.Time       `json:"updatedAt"`
	AuthorID  int             `json:"author_id"`
	TeamID    int             `json:"team_id"`
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
	Tasks     []*Tasks  `json:"tasks"`
	TaskID    *int      `json:"task_id"`
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
	Summary     string    `json:"summary"`
	UpdatedAt   time.Time `json:"updatedAt"`
	EventID     int       `json:"event_id"`
	Talks       []*Talk   `json:"talks"`
	CreatedBy   []*Event  `json:"createdBy"`
}

type UpdateAttendee struct {
	User []*CreateUser `json:"user"`
}

type UpdateBugReport struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateCartItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       *string `json:"price"`
	IsFree      bool    `json:"isFree"`
}

type UpdateEvent struct {
	Name                  *string   `json:"name"`
	Summary               *string   `json:"summary"`
	Alias                 *string   `json:"alias"`
	Description           *string   `json:"description"`
	EventType             *string   `json:"eventType"`
	TotalAttendees        *int      `json:"totalAttendees"`
	Email                 *string   `json:"Email"`
	SpeakerConduct        *string   `json:"speakerConduct"`
	Website               *string   `json:"website"`
	TrackID               *int      `json:"track_id"`
	Venue                 *string   `json:"venue"`
	EventDate             []*string `json:"EventDate"`
	VolunteerID           *int      `json:"volunteer_id"`
	SponsorID             *int      `json:"sponsor_id"`
	MediaLinks            []string  `json:"mediaLinks"`
	Actions               []string  `json:"actions"`
	IsAcceptingAttendees  *bool     `json:"isAcceptingAttendees"`
	IsVirtual             *bool     `json:"isVirtual"`
	ConfirmedEmail        *bool     `json:"ConfirmedEmail"`
	IsArchived            *bool     `json:"isArchived"`
	IsLocked              *bool     `json:"isLocked"`
	IsAcceptingVolunteers *bool     `json:"isAcceptingVolunteers"`
	IsAcceptingTalks      *bool     `json:"isAcceptingTalks"`
}

type UpdateEventSettings struct {
	ShowWelcomeMeetupGroup      *bool   `json:"showWelcomeMeetupGroup"`
	ShowTeamInstruction         *bool   `json:"showTeamInstruction"`
	ShowInvitationInstruction   *bool   `json:"showInvitationInstruction"`
	ShowWelcomeEventInstruction *bool   `json:"showWelcomeEventInstruction"`
	EventThemeColour            *string `json:"eventThemeColour"`
}

type UpdateFeatureRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateMeetupGroup struct {
	Name          *string   `json:"name"`
	Description   string    `json:"description"`
	Location      *string   `json:"location"`
	SponsorID     *int      `json:"sponsor_id"`
	FacilitatorID *int      `json:"facilitator_id"`
	Website       *string   `json:"website"`
	Alias         *string   `json:"alias"`
	MediaLinks    []*string `json:"mediaLinks"`
	Actions       []*string `json:"actions"`
	LeadID        int       `json:"lead_id"`
}

type UpdateNote struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateSponsor struct {
	Name             string  `json:"name"`
	ImageURL         *string `json:"image_url"`
	Type             string  `json:"type"`
	SponsorhipStatus string  `json:"sponsorhipStatus"`
	IsOrganization   bool    `json:"isOrganization"`
}

type UpdateSubmittedTalk struct {
	IsAccepted bool      `json:"isAccepted"`
	Track      *string   `json:"track"`
	Comment    *string   `json:"comment"`
	MediaLinks []*string `json:"mediaLinks"`
}

type UpdateTalk struct {
	Title        string    `json:"title"`
	TalkCoverURI *string   `json:"talkCoverUri"`
	Summary      *string   `json:"summary"`
	Description  *string   `json:"description"`
	Archived     *bool     `json:"Archived"`
	Duration     *string   `json:"duration"`
	Tags         []*string `json:"tags"`
}

type UpdateTask struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
}

type UpdateTeam struct {
	Name    string        `json:"name"`
	Members []*CreateUser `json:"members"`
	Goal    string        `json:"goal"`
	TaskID  *int          `json:"task_id"`
}

type UpdateTrack struct {
	Name        string    `json:"name"`
	TrackImgURI *string   `json:"trackImgUri"`
	Summary     *string   `json:"summary"`
	Duration    string    `json:"duration"`
	TotalTalks  int       `json:"totalTalks"`
	IsCompleted bool      `json:"isCompleted"`
	Archived    bool      `json:"Archived"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UpdateUser struct {
	Name         *string            `json:"name"`
	Email        *string            `json:"email"`
	Password     *string            `json:"password"`
	ImgURI       *string            `json:"img_uri"`
	FileURI      *string            `json:"file_uri"`
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
	EventID *int           `json:"eventId"`
	Type    string         `json:"type"`
	UserID  *int           `json:"userId"`
}

type User struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	BucketName   string       `json:"bucketName"`
	Talks        []*Talk      `json:"talks"`
	Events       []*Event     `json:"events"`
	EventID      int          `json:"event_id"`
	FileID       *int         `json:"file_id"`
	Reminders    []*Reminder  `json:"reminders"`
	ImgURI       *string      `json:"img_uri"`
	CreatedAt    time.Time    `json:"createdAt"`
	Files        []*UserFile  `json:"files"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	VolunteerID  int          `json:"volunteer_id"`
	Volunteering []*Volunteer `json:"volunteering"`
	Attending    []*Attendee  `json:"attending"`
}

type UserFile struct {
	ID        int            `json:"id"`
	File      graphql.Upload `json:"file"`
	Type      string         `json:"type"`
	User      []*User        `json:"user"`
	UserID    int            `json:"userId"`
	FileURI   string         `json:"file_uri"`
	Timestamp time.Time      `json:"timestamp"`
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
