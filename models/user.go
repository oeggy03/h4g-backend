package models

//bcrypt is for encrypting the password
import "golang.org/x/crypto/bcrypt"

type User struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
	Type        uint   `json:"type"` //type refers to accoount type (disabled person (0) or not disabled(1))
	Password    []byte `json:"-"`    //"-" means that password is not returned when user is retrieved
}

//for sign up function
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

//for sign in function
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
