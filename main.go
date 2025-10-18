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

	languageRepo := repository.NewLanguageRepository(config.DB)
	languageService := service.NewLanguageService(languageRepo)
	languageHandler := handler.NewLanguageHandler(languageService)

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	personRepo := repository.NewPersonRepository(config.DB)
	personService := service.NewPersonService(personRepo)
	personHandler := handler.NewPersonHandler(personService)

	positionRepo := repository.NewPositionRepository(config.DB)
	positionService := service.NewPositionService(positionRepo)
	positionHandler := handler.NewPositionHandler(positionService)

	referenceRepo := repository.NewReferenceRepository(config.DB)
	referenceService := service.NewReferenceService(referenceRepo)
	referenceHandler := handler.NewReferenceHandler(referenceService)

	SetupRouter(app, userHandler, personHandler, positionHandler, referenceHandler, languageHandler)

	port := "8081"
	log.Printf("🚀 Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
