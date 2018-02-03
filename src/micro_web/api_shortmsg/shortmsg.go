package main

import (
	"encoding/json"
	"log"
	shortmsg "micro_web/svr_shortmsg/proto/shortmsg"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"

	"golang.org/x/net/context"
	"strconv"
	"micro_web/consts"
	"time"
)

type ShortMsgApi struct {
	Client shortmsg.ShortmsgClient
}

func (s *ShortMsgApi) SendRegisterCode(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	phone, ok := req.Post["phone"]
	if !ok || len(phone.Values) == 0 {
		return errors.BadRequest(consts.API_NAME_SHORTMSG, "手机号码为空")
	}
	p, err := strconv.ParseInt(phone.Values[0], 10, 64)
	if err!=nil {
		return errors.BadRequest(consts.API_NAME_SHORTMSG, "手机号码不正确")
	}
	response, err := s.Client.SendRegisterCode(ctx, &shortmsg.SendRegisterCodeReq{
		Phone:p,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": string(response.Tag),
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name(consts.API_NAME_SHORTMSG),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&ShortMsgApi{Client: shortmsg.NewShortmsgClient(consts.SVR_NAME_SHORTMSG, service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
