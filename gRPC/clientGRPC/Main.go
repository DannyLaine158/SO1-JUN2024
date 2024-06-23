package main

import (
	pb "clientGRPC/client"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var ctx = context.Background()

type Data struct {
	Texto string
	Pais  string
}

func sendData(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	tweet := Data{
		Texto: data["texto"],
		Pais:  data["pais"],
	}

	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Texto: tweet.Texto,
		Pais:  tweet.Pais,
	})
	if err != nil {
		return err
	}

	fmt.Println("Respuesta del servidor ", ret)

	return nil
}

func main() {
	app := fiber.New()

	app.Post("/insert", sendData)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
