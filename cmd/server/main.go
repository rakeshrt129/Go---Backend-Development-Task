package main

import (
	"log"
	"os"

	"user-age-api/internal/handler"
	"user-age-api/internal/logger"
	"user-age-api/internal/middleware"
	"user-age-api/internal/repository"
	"user-age-api/internal/routes"
	"user-age-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)


func main() {

	// -------------------- INIT LOGGER --------------------
	logger.InitLogger()
	defer logger.Log.Sync()

	logger.Log.Info("starting user-age-api server")

	// -------------------- DB CONNECTION --------------------
	// Example DB URL:
	// postgres://username:password@localhost:5432/user_age_db?sslmode=disable

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	db, err := repository.NewDB(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	queries := repository.NewQueries(db)

	// -------------------- SERVICE & HANDLER --------------------
	userService := service.NewUserService(queries)
	userHandler := handler.NewUserHandler(userService)

	// -------------------- FIBER APP --------------------
	app := fiber.New()

	// Middleware
	app.Use(middleware.RequestLogger())

	// Routes
	routes.RegisterUserRoutes(app, userHandler)

	// -------------------- START SERVER --------------------
	port := ":3000"
	logger.Log.Info("server running", zap.String("port", port))

	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
