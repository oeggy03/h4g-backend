package models

import "time"

type Activity struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"postdate"`
	Name        string    `json:"name"`
	Desc        string    `json:"desc"`
	Time        string    `json:"time"`
	Location    string    `json:"location"`
	CreatorType uint      `json:"creatortype"` //account type of the creator.
	UserID      uint      `json:"userid"`      //foreign key. This refers to the person who made the activity
	User        User      `json:"user":gorm:"foreignkey:UserID"`
}
