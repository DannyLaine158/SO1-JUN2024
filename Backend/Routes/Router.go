package Routes

import (
	"Backend/Instance"
	"Backend/Model"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"os/exec"
	"strconv"
)

func Setup(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Hello World")
	})

	app.Get("/getRam", func(ctx *fiber.Ctx) error {
		nameCol := "ram"
		collection := Instance.Mg.Db.Collection(nameCol)

		err := collection.Drop(context.TODO())
		if err != nil {
			return err
		}

		cmd := exec.Command("sh", "-c", "cat /proc/ram_201800722")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalln(err)
		}
		out := string(output[:])

		doc := Model.Data{Percent: out}

		_, err = collection.InsertOne(context.TODO(), doc)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(fiber.Map{
			"status":  200,
			"percent": out,
		})
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
