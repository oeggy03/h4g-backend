package models

type Activity struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Location    string `json:"location"`
	CreatorType uint   `json:"creatortype"` //account type of the creator.
	UserID      uint   `json:"userid"`      //foreign key. This refers to the person who made the activity
	User        User   `json:"user":gorm:"foreignkey:UserID"`
}
