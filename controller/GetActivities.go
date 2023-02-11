package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/h4g-backend/connect"
	"github.com/oeggy03/h4g-backend/models"
)

func GetActivities(c *fiber.Ctx) error {
	var disaAct []models.Activity
	var enaAct []models.Activity

	//retrieving all activities by "disabled" people
	if err := connect.DB.Find(&disaAct, "creator_type = ?", 0); err != nil {
		fmt.Println("Error retrieving activities by special friends")
	}

	//retrieving all activities by "not disabled" people
	if err := connect.DB.Find(&enaAct, "creator_type = ?", 1); err != nil {
		fmt.Println("Error retrieving activities by best buddies")
	}

	return (c.JSON(fiber.Map{
		"disabled": disaAct,
		"enabled":  enaAct}))
}
