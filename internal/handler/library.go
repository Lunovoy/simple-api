package handler

import (
	"context"

	"github.com/Lunovoy/simple-api/internal/models"
	"github.com/Lunovoy/simple-api/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type libraryDTO struct {
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

func CreateLibrary(c *fiber.Ctx) error {

	nLibrary := new(libraryDTO)
	if err := c.BodyParser(nLibrary); err != nil {
		return err
	}

	libraryCollection := repository.GetCollection("libraries")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), nLibrary)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}

func GetLibraries(c *fiber.Ctx) error {
	libraryCollection := repository.GetCollection("libraries")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	var libraries []models.Library
	if err := cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}

	return c.JSON(libraries)
}
