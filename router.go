package main

import (
	"cv_backend/app/handler"
	"cv_backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(
	app *fiber.App,
	userHandler *handler.UserHandler,
	personHandler *handler.PersonHandler,
	positionHandler *handler.PositionHandler,
	referenceHandler *handler.ReferenceHandler,
	languageHandler *handler.LanguageHandler,
) {
	api := app.Group("/api")

	// ✅ Health check (DB & server)
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Pong! ✅ Server & DB are working."})
	})

	// ✅ Public (auth yok)
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	// ✅ Protected (JWT zorunlu)
	auth := api.Group("/auth", middleware.JWTMiddleware())
	auth.Get("/me", userHandler.GetProfile)

	// ✅ Person işlemleri (başvurular)
	auth.Get("/persons", personHandler.GetAllPersons)
	auth.Get("/persons/:id", personHandler.GetPersonByID)
	auth.Put("/persons/:id/status", personHandler.UpdatePersonStatus)
	auth.Delete("/persons/:id", personHandler.DeletePerson)

	// ✅ Language işlemleri
	auth.Get("/languages", languageHandler.GetAllLanguages)
	auth.Get("/languages/:id", languageHandler.GetLanguageByID)
	// Eğer ileride eklenecekse:
	// auth.Delete("/languages/:id", languageHandler.DeleteLanguage)

	// ✅ Position işlemleri
	auth.Get("/positions", positionHandler.GetAllPositions)
	auth.Get("/positions/:id", positionHandler.GetPositionByID)
	auth.Delete("/positions/:id", positionHandler.DeletePosition)

	// ✅ Reference işlemleri
	auth.Get("/references", referenceHandler.GetAllReferences)
	auth.Get("/references/:id", referenceHandler.GetReferenceByID)
	auth.Delete("/references/:id", referenceHandler.DeleteReference)
}
