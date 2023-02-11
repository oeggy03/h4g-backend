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

func CreateActivity(c *fiber.Ctx) error {
	//This section makes sure that only users who are logged in can make posts
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
	intID, _ := strconv.Atoi(claims.Issuer) //this is the userID (integer)

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Create Post: Unable to parse body")
	}

	if len(data["name"].(string)) == 0 || len(data["desc"].(string)) == 0 || len(data["location"].(string)) == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Create Activity: Please do not leave any blank fields!",
		})
	}

	//Retrieve the creator's details
	var userDetails models.User
	if err := connect.DB.Find(&userDetails, "id = ?", intID); err != nil {
		fmt.Println("Error retrieving user details")
	}

	Activity := models.Activity{
		Name:        data["name"].(string),
		Desc:        data["desc"].(string),
		Time:        data["time"].(string),
		Location:    data["location"].(string),
		CreatorType: userDetails.Type,
		UserID:      uint(intID),
	}

	//Creates the activity entry in the db
	if err := connect.DB.Create(&Activity).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Create Activity: Invalid payload",
		})
	}

	c.Status(200)
	return c.JSON(Activity)
}
