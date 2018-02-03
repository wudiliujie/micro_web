package main

import (
	"log"
	"time"

	account "micro_web/svr_account/proto/account"
	 db "micro_web/svr_account/db"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	db2 "micro_web/db"
)

type Account struct{}

func (s *Account) Login(ctx context.Context, req *account.LoginRequest, rsp *account.LoginResponse) error {
	user := db.GetUserInfo(1);
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + user.UserName+":"+user.UserPass
	return nil
}

func (s *Account) GameLogin(ctx context.Context, req *account.GameLoginRequest, rsp *account.GameLoginResponse) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}



func main() {
	db2.G_dbmanage.Init();
	service := micro.NewService(
		micro.Name("yxl.micro.srv.account"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	account.RegisterAccountHandler(service.Server(), new(Account))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
