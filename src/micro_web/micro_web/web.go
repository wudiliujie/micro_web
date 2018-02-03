package main

import (
	"log"
	"time"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"micro_web/greeter"
	"strconv"
)

type Say struct{}

func (s *Say) Hello11(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	log.Print("Received Say.Hello request")
	count:=0
	for i:=0;i<10;i++{
		count+=i
	}
	rsp.Greeting = "Hello " + req.Name+ ":"+strconv.Itoa(count)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	greeter.RegisterGreeterHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}