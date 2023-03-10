package controller

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/h4g-backend/connect"
	"github.com/oeggy03/h4g-backend/models"
	"github.com/oeggy03/h4g-backend/util"
)

func GetActivity(c *fiber.Ctx) error {
	activityID := c.Params("id")

	//retrieving the jwt so that we may verify if our user is the creator of the activity.
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(util.SecretKey), nil
	})

	//error handling
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "You are unauthorized! Please sign in first.",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	intID, _ := strconv.Atoi(claims.Issuer) //this is the userID (integer)

	//Retrieve the activity details
	var activity models.Activity

	if err := connect.DB.First(&activity, "id = ?", activityID); err != nil {
		fmt.Println("Error retrieving activity")
	}

	if activity.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{"message": "Activity not found."})
	}

	//Retrieve the activity's creator
	var user models.User

	if err := connect.DB.First(&user, "id = ?", activity.UserID); err != nil {
		fmt.Println("Error retrieving creator for activity")
	}
	fmt.Println("hello")

	//Retrieve the ID of the joiners for the activity
	var joinerIDs []uint
	if err := connect.DB.Table("joiners").Select("user_id").Where("activity_id = ?", activity.ID).Find(&joinerIDs); err != nil {
		fmt.Println("Error retrieving joiners for activity")
	}

	//Retrieve all participant details for the activity.
	var participants []models.User

	if len(joinerIDs) != 0 {
		// SELECT * FROM users WHERE id IN joinerIDs;
		if err := connect.DB.Find(&participants, joinerIDs); err != nil {
			fmt.Println("Error retrieving participant details for activity")
		}
	} else {
		participants = []models.User{}
	}

	//Checks if our user has joined the activity
	var joined bool = false
	for i := range joinerIDs {
		if joinerIDs[i] == uint(intID) {
			joined = true
		}
	}

	//Retrieve the comments
	//.Order("id desc") add in front of .Find for descending
	var comments []models.Comment
	if err := connect.DB.Order("id desc").Find(&comments, "activity_id = ?", activity.ID); err != nil {
		fmt.Println("Error retrieving comments for activity")
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"user_id":      intID,
		"creator":      user.Username,
		"activity":     activity,
		"participants": participants,
		"owner":        activity.UserID == uint(intID),
		"joined":       joined,
		"comments":     comments,
		"message":      "Success retrieving activity!",
	})
}
