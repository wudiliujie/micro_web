package main

import (
	"encoding/json"
	"log"
	"strings"

	account "micro_web/svr_account/proto/account"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"

	"golang.org/x/net/context"
	"time"
)

type Accountapi struct {

	Client account.AccountClient
}

func (s *Accountapi) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.account", "Name cannot be blank")
	}
	pass, ok := req.Get["Pass"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.account", "Name cannot be blank")
	}

	response, err := s.Client.Login(ctx, &account.LoginRequest{
		Name: strings.Join(name.Values, " "),
		Pass: strings.Join(pass.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.account"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Accountapi{Client: account.NewAccountClient("yxl.micro.srv.account", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
