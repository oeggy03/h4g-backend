package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oeggy03/h4g-backend/controller"
	// "github.com/oeggy03/h4g-backend/middleware" for the middleware if i have time
)

func Setup(app *fiber.App) {
	app.Post("/api/SignUp", controller.SignUp)
	app.Post("/api/SignIn", controller.SignIn)
	app.Get("/api/GetUser", controller.GetUser)
	app.Get("/api/GetProfile/:id", controller.GetProfile)
	app.Post("/api/SignOut", controller.SignOut)
	app.Get("/api/GetActivities", controller.GetActivities)
	app.Post("/api/CreateActivity", controller.CreateActivity)
	app.Get("/api/GetActivity/:id", controller.GetActivity)
	app.Post("/api/JoinActivity", controller.JoinActivity)
	app.Delete("/api/LeaveActivity/:id", controller.LeaveActivity)
	app.Delete("/api/DeleteActivity/:id", controller.DeleteActivity)
	app.Put("/api/UpdateActivity", controller.UpdateActivity)
	app.Post("/api/CreateComment", controller.CreateComment)
	app.Delete("/api/DeleteComment/:id", controller.DeleteComment)
	app.Put("/api/UpdateComment", controller.UpdateComment)

	// app.Get("/api/RetrievePost/:id", controller.RetrievePost)

	// app.Post("/api/CreateComment", controller.CreateComment)
	// app.Get("/api/RetrieveComments/:id", controller.RetrieveComments)

}
