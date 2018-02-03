package main

import (
	"encoding/json"
	"log"
	"strings"

	shortmsg "svr_shortmsg/proto/shortmsg"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"

	"golang.org/x/net/context"
	"core"
	"strconv"
)

type ShortMsgApi struct {
	Client shortmsg.ShortmsgClient
}

func (s *ShortMsgApi) SendRegisterCode(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	phone, ok := req.Post["phone"]
	if !ok || len(phone.Values) == 0 {
		return errors.BadRequest(core.API_NAME_SHORTMSG, "手机号码为空")
	}
	p, err := strconv.ParseInt(phone.Values[0], 10, 64)
	if err!=nil {
		return errors.BadRequest(core.API_NAME_SHORTMSG, "手机号码不正确")
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
		micro.Name(core.API_NAME_SHORTMSG),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&ShortMsgApi{Client: shortmsg.NewShortmsgClient(core.SVR_NAME_SHORTMSG, service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
