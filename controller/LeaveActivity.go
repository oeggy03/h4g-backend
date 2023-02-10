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

func LeaveActivity(c *fiber.Ctx) error {
	activityID, _ := strconv.Atoi(c.Params("id"))

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(util.SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	//Access issuer through claims.Issuer
	claims := token.Claims.(*jwt.StandardClaims)
	intUserID, _ := strconv.Atoi(claims.Issuer)

	var testJoiner models.Joiner
	//Find the joiner with given activity id
	if err := connect.DB.Where(&models.Joiner{ActivityID: uint(activityID), UserID: uint(intUserID)}).First(&testJoiner); err != nil {
		fmt.Println("Error checking joiner info for activity")
	}

	if int(testJoiner.ID) == 0 {
		c.JSON(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "You have not joined this activity!",
		})
	}

	//Deletes the joiner entry
	var joinerFormat models.Joiner

	connect.DB.Delete(&joinerFormat, testJoiner.ID)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Successfully left activity.",
	})

}
