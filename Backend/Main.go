package main

import (
	"Backend/Database"
	"Backend/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()

	if err := Database.Connect(); err != nil {
		log.Fatal("Error en", err)
	}

	app.Use(cors.New())

	Routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}

//func getMem() {
//	for range time.Tick(time.Second * 1) {
//		percentMem := strconv.Itoa(rand.Intn(100))
//		fmt.Println(percentMem)
//		Controller.InsertData("ram", percentMem)
//	}
//}
