package routes

import (
	"fmt"
	"sharir/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func CreateAuthRoutes(app *fiber.App, userRepo *auth.Repo) {
	fmt.Println("Hello")
}
