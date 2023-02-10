package models

//this refers to the people who join an activity
type Joiner struct {
	ID         uint     `json:"id"`
	UserID     uint     `json:"userid"`     //foreign key too joiner's ID
	ActivityID uint     `json:"activityid"` //foreign key to activityID
	User       User     `json:"user":gorm:"foreignkey:UserID"`
	Activity   Activity `json:"activity":gorm:"foreignkey:ActivityID"`
}
