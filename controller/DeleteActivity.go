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

func DeleteActivity(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	activityID1, _ := strconv.Atoi(c.Params("id"))
	activityID := uint(activityID1)

	var activityFormat models.Activity
	var creatorID int64
	var userID int

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

	//Checks if user is indeed the activity creator
	if err := connect.DB.Table("activities").Select("user_id").Where("id = ?", activityID).Find(&creatorID); err != nil {
		fmt.Println("Error retrieving activity's creator: delete activity")
	}

	userID, _ = strconv.Atoi(claims.Issuer)

	if creatorID != int64(userID) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "You are not the creator of this activity.",
		})
	}

	//To prevent the foreign key constraint - cannot delete parent row error, we need to delete the activity' joiners first
	var joinerFormat models.Joiner
	connect.DB.Where("activity_id = ?", activityID).Delete(&joinerFormat)

	//Now we delete from activities
	connect.DB.Delete(&activityFormat, activityID)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Activity deleted.",
	})
}
