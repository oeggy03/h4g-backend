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

func GetProfile(c *fiber.Ctx) error {
	profileID1, _ := strconv.Atoi(c.Params("id"))
	profileID := uint(profileID1)

	//makes sure that user has signed in
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(util.SecretKey), nil
	})

	//error handling
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	var profile models.User

	//Find that user
	connect.DB.Where("id = ?", profileID).First(&profile)

	if profile.ID == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sorry, user not found.",
		})
	}

	//Slice of activities
	var activities []models.Activity

	//Find the user's created activities
	connect.DB.Where("user_id = ?", profileID).Find(&activities)

	//Retrieve joiners where user_id = user
	var joinedActIDs []uint
	if err := connect.DB.Table("joiners").Select("activity_id").Where("user_id = ?", profileID).Find(&joinedActIDs); err != nil {
		fmt.Println("Error retrieving joiners for activity")
	}

	//Find user's joined activities
	var activitiesJoined []models.Activity

	if len(joinedActIDs) != 0 {
		// SELECT * FROM users WHERE id IN ;
		if err := connect.DB.Find(&activitiesJoined, joinedActIDs); err != nil {
			fmt.Println("Error retrieving activities joined for user")
		}
	} else {
		activities = []models.Activity{}
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message":            "User found",
		"profile":            profile,
		"activities_created": activities,
		"activities_joined":  activitiesJoined,
	})
}
