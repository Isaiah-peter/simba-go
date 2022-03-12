package routes

import (
	"simba-clone/pkg/controllers"
	utils "simba-clone/pkg/util"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route *fiber.App) {
	route.Post("/login", controllers.Login)
	route.Post("/register", controllers.Register)
	route.Get("/user/all", utils.Authuser(), controllers.GetAllUser)
	route.Get("/user/:id", utils.Authuser(), controllers.GetSingleUser)
	route.Put("/user/update/:id", utils.Authuser(), controllers.UpdateUser)
}

func Transaction(route *fiber.App) {
	route.Post("/transfer", utils.Authuser(), controllers.CreateTransaction)
	route.Get("/alltransaction", utils.Authuser(), controllers.GetAllTransaction)
	route.Put("/update/transfer", utils.Authuser(), controllers.UpdateTransaction)
}
