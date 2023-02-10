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

func DeleteComment(c *fiber.Ctx) error {
	commentID, _ := strconv.Atoi(c.Params("id"))

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

	var testComment models.Comment
	//Find the comment with given comment id
	if err := connect.DB.Find(&testComment, "id = ?", commentID); err != nil {
		fmt.Println("Error retrieving comment")
	}

	intUserID, _ := strconv.Atoi(claims.Issuer)

	if int(testComment.UserID) != intUserID {
		c.JSON(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "You are not the creator of this comment.",
		})
	}

	//Deletes the comment
	var commentFormat models.Comment

	connect.DB.Delete(&commentFormat, commentID)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Comment deleted successfully",
	})

}
