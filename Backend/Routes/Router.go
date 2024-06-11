package Routes

import (
	"Backend/Instance"
	"Backend/Model"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"strconv"
)

func Setup(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Hello World")
	})

	app.Get("/insertRam", func(ctx *fiber.Ctx) error {

		nameCol := "ram"
		dataParam := strconv.Itoa(rand.Intn(100))

		collection := Instance.Mg.Db.Collection(nameCol)
		doc := Model.Data{Percent: dataParam}

		_, err := collection.InsertOne(context.TODO(), doc)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(201).JSON(dataParam)
	})
}
