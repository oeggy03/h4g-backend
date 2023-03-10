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

func JoinActivity(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
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
	claims := token.Claims.(*jwt.StandardClaims)
	intID, _ := strconv.Atoi(claims.Issuer)

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Join Activity: Unable to parse body")
	}

	idActivity, _ := strconv.Atoi(data["activity_id"].(string))
	//Retrieve the activity details
	var activity models.Activity

	if err := connect.DB.First(&activity, "id = ?", uint(idActivity)); err != nil {
		fmt.Println("Error retrieving activity")
	}

	if activity.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{"message": "Activity not found."})
	}

	//Check if entry already exists
	var joinerCheck models.Joiner

	if err := connect.DB.Where(&models.Joiner{ActivityID: uint(idActivity), UserID: uint(intID)}).First(&joinerCheck); err != nil {
		fmt.Println("Error checking joiner info for activity")
	}

	if joinerCheck.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "You have already joined the activity!"})
	}

	//Everything is normal, we can add in the joiner
	joiner := models.Joiner{
		UserID:     uint(intID),
		ActivityID: uint(idActivity),
	}

	if err := connect.DB.Create(&joiner).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Add activity joiner: Invalid payload",
		})
	}

	c.Status(200)
	return c.JSON(joiner.ActivityID)
}
