package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sharir/api/routes"
	"sharir/pkg/configuration"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	def := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Request-With",
		AllowCredentials: true,
	}

	app.Use(cors.New(def))
	godotenv.Load()
	config := configuration.FromEnv()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Panic(err)
	}
	db := client.Database("sharir")
	fmt.Println(db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"ping": "pong",
		})
	})
	routes.CreatePhoneOtpRoutes(app)
	log.Panic(app.Listen(":" + os.Getenv("PORT")))
}
