package main

import (
	"Backend/Controller"
	"Backend/Database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	app := fiber.New()

	if err := Database.Connect(); err != nil {
		log.Fatal("Error en", err)
	}

	getMem()

	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 500)
}

func getMem() {
	for range time.Tick(time.Second * 1) {
		percentMem := strconv.Itoa(rand.Intn(100))
		fmt.Println(percentMem)
		Controller.InsertData("ram", percentMem)
	}
}
