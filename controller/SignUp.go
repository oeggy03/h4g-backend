package controller

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/h4g-backend/connect"
	"github.com/oeggy03/h4g-backend/models"
)

func SignUp(c *fiber.Ctx) error {
	var data map[string]interface{}
	var checkUser models.User
	var checkEmail models.User
	var checkPhone models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Create User: Unable to parse body")
	}

	//Checks if any field is blank
	if len(data["username"].(string)) == 0 || len(data["password"].(string)) == 0 || len(data["email"].(string)) == 0 ||
		len(data["name"].(string)) == 0 || len(data["phone"].(string)) == 0 || (data["type"].(float64) == 2) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Please do not leave any field empty.",
		})
	}

	connect.DB.Where("username = ?", data["username"].(string)).First(&checkUser)

	//If username is taken
	if checkUser.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sorry, the username " + data["username"].(string) + " is taken.",
		})
	}

	//Checks if email has been used
	connect.DB.Where("email = ?", data["email"].(string)).First(&checkEmail)

	if checkEmail.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sorry, the email " + data["email"].(string) + " has already been used.",
		})
	}

	//Checks if phone number has been used
	connect.DB.Where("phone_number = ?", data["phone"].(string)).First(&checkPhone)

	if checkPhone.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Sorry, the number " + data["phone"].(string) + " has already been used.",
		})
	}

	user := models.User{
		Username:    data["username"].(string),
		Email:       strings.TrimSpace(data["email"].(string)),
		PhoneNumber: strings.TrimSpace(data["phone"].(string)),
		Type:        uint(data["type"].(float64)),
	}

	user.SetPassword(data["password"].(string))

	if err := connect.DB.Create(&user).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Create User: Invalid payload.",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Success! Please close this popup and sign in.",
	})
}
