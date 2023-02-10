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

func UpdateActivity(c *fiber.Ctx) error {
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

	//Access issuer through claims.Issuer
	// claims := token.Claims.(*jwt.StandardClaims)
	// userIDint, _ := strconv.Atoi(claims.Issuer)

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Update Activity: Unable to parse body")
	}

	var activityUpdate models.Activity

	//convert the activityid received to int
	activityIDint, _ := strconv.Atoi(data["activity_id"].(string))

	//retrieves the activity with given id
	if err := connect.DB.First(&activityUpdate, activityIDint); err != nil {
		fmt.Println("Error retrieving activity for update")
	}

	//Assign the updated vars to activityUpdate
	activityUpdate.Name = data["name"].(string)
	activityUpdate.Desc = data["desc"].(string)
	activityUpdate.Location = data["location"].(string)

	//updates the activity
	if err := connect.DB.Save(&activityUpdate); err != nil {
		fmt.Println("Error saving updated activity")
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Activity updated successfully!"})
}
