package models

import "time"

type Comment struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"postdate"`
	Content    string    `json:"content"`
	ActivityID uint      `json:"activityid"` //foreign key to activityID
	UserID     uint      `json:"userid"`     //foreign key. This refers to the person who made the activity
	Creator    string    `json:"creator"`
	Activity   Activity  `json:"activity":gorm:"foreignkey:ActivityID"`
	User       User      `json:"user":gorm:"foreignkey:UserID"`
}
