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

func UpdateComment(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	//We need the token to verify that the user is indeed the creator of the comment
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

	var data map[string]interface{}

	err = c.BodyParser(&data)

	if err != nil {
		fmt.Println("Unable to parse body")
	}

	//Retrieves comment. Checks if comment is created by user
	var targetComment models.Comment
	var intCommentID = data["id"].(float64)
	if err := connect.DB.Find(&targetComment, "id = ?", intCommentID); err != nil {
		fmt.Println("Error retrieving comment")
	}

	intUserID, _ := strconv.Atoi(claims.Issuer)

	if int(targetComment.UserID) != intUserID {
		c.JSON(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "You are not the creator of this comment.",
		})
	}

	targetComment.Content = data["content"].(string)

	//Updates comment to MySQL
	if err := connect.DB.Save(&targetComment); err != nil {
		fmt.Println("Error saving updated comment")
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Comment updated successfully.",
	})
}
