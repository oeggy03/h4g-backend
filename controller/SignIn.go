package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/h4g-backend/connect"
	"github.com/oeggy03/h4g-backend/models"
	"github.com/oeggy03/h4g-backend/util"
)

func SignIn(c *fiber.Ctx) error {

	//stores the parsed JSON
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		fmt.Println("Unable to parse body")
	}

	var user models.User

	//check if username exists in our database
	connect.DB.Where("username=?", data["username"]).First(&user)

	//means that user not found
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Sorry, user not found.",
		})
	}

	//compares the password with the one stored in database
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password.",
		})
	}

	//Itoa converts uint userID to a string
	//Generates the jwt token
	token, err := util.GenerateJwt(strconv.Itoa((int(user.ID))))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "jwt error",
		})
	}

	//store the token in a cookie
	//cookie is only supposed to be stored by frontend but not accessed
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,                          //token that i created earlier
		Expires:  time.Now().Add(time.Hour * 24), //no need for unix time
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Login Successful! You may close the popup.",
	})

}
