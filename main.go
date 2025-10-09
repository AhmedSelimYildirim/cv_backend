package main

import (
	"cv_backend/app/handler"
	"cv_backend/app/repository"
	"cv_backend/app/service"
	"cv_backend/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Setup()

	app := fiber.New()

	// Language
	languageRepo := repository.NewLanguageRepository(config.DB)
	languageService := service.NewLanguageService(languageRepo)
	languageHandler := handler.NewLanguageHandler(languageService)

	// User
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Person
	personRepo := repository.NewPersonRepository(config.DB)
	personService := service.NewPersonService(personRepo)
	personHandler := handler.NewPersonHandler(personService)

	// Position
	positionRepo := repository.NewPositionRepository(config.DB)
	positionService := service.NewPositionService(positionRepo)
	positionHandler := handler.NewPositionHandler(positionService)

	// Reference
	referenceRepo := repository.NewReferenceRepository(config.DB)
	referenceService := service.NewReferenceService(referenceRepo)
	referenceHandler := handler.NewReferenceHandler(referenceService)

	// Router
	SetupRouter(app, userHandler, personHandler, positionHandler, referenceHandler, languageHandler)

	port := "8081"
	log.Printf("🚀 Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
