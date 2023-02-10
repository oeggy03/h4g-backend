package controller

import (
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
		c.JSON(fiber.Map{
			"message": "Sorry, user not found.",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "User found",
		"profile": profile,
	})
}
