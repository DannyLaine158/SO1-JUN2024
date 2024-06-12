package Routes

import (
	"Backend/Instance"
	"Backend/Model"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
)

func Setup(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Hello World")
	})

	app.Get("/insertRam", func(ctx *fiber.Ctx) error {
		nameCol := "ram"
		collection := Instance.Mg.Db.Collection(nameCol)

		err := collection.Drop(context.TODO())
		if err != nil {
			return err
		}

		dataParam := strconv.Itoa(rand.Intn(100))

		// collection = Instance.Mg.Db.Collection(nameCol)
		doc := Model.Data{Percent: dataParam}

		_, err = collection.InsertOne(context.TODO(), doc)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(201).JSON(dataParam)
	})

	app.Get("/insertProcess", func(ctx *fiber.Ctx) error {
		log.Println("Insertando proceso")

		cmd := exec.Command("sleep", "infinity")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"pid":     cmd.Process.Pid,
		})
	})

	app.Get("/delProcess", func(ctx *fiber.Ctx) error {
		pid := ctx.Query("pid")
		pidInt, err := strconv.Atoi(pid)
		if err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("kill", "-9", strconv.Itoa(pidInt))
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
		})
	})
}
