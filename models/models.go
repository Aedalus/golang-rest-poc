package models

import "time"

type ReviewType string

var REVIEW_BOXER ReviewType = "REVIEW_BOXER"
var REVIEW_FIGHT_PREDICTION ReviewType = "FIGHT_PREDICTION"
var REVIEW_FIGHT ReviewType = "REVIEW_FIGHT"

type LiveStream struct {
	ID   uint
	Name string

	CreatedAt time.Time
}

type Review struct {
	ID            uint
	Reviewer      uint
	ReviewSubject uint
	ReviewType    ReviewType
	Review        string
	Likes         int
	Shares        int
	Flag          int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID       uint
	Sub      int
	Nickname string
	Gravatar string
	Email    string
	Phone    string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Boxer struct {
	ID          uint
	Description string
	FirstName   string
	LastName    string
	Wins        int
	Losses      int
	Draws       int
	KOs         int
	IG          string
	Twitter     string
	FB          string
	ProfileImg  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Article struct {
	ID       uint
	Title    string
	Subtitle string
	Author   int
	Article  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Gym struct {
	ID           uint
	Name         string
	Address      string
	Owner        string
	ProfileImage string

	CreatedAt time.Time
	UpdatedAt time.Time
}
