package main

import (
	"grpc-db-poc/server/api"
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from Ping: %s", response.Greeting)

	h := api.NewHealthClient(conn)

	fmt.Println(h)
	healthCheck, err := h.HealthCheck(context.Background(), &api.HealthMessage{})

	fmt.Println(err)
	log.Printf("Response from health : %s \n", healthCheck.Health)

}
