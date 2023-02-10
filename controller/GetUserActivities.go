package controller

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/h4g-backend/connect"
	"github.com/oeggy03/h4g-backend/models"
	"github.com/oeggy03/h4g-backend/util"
)

func GetUserActivities(c *fiber.Ctx) error {
	profileID1, _ := strconv.Atoi(c.Params("id"))
	profileID := uint(profileID1)

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(util.SecretKey), nil
	})

	//error handling
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	//Slice of activities
	var activities []models.Activity

	//Find the user's created activities
	connect.DB.Where("user_id = ?", profileID).Find(&activities)

	c.Status(200)
	return c.JSON(activities)
}
