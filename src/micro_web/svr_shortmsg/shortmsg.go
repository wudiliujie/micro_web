package main

import (
	"log"
	"time"

	shortmsg "micro_web/svr_shortmsg/proto/shortmsg"
	"micro_web/svr_shortmsg/db"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"micro_web/core"
	db2 "micro_web/db"
	"math/rand"
	"fmt"
	"micro_web/consts"
)

type ShortMsgSvr struct{}

func (s *ShortMsgSvr) GetRegisterCode(ctx context.Context, req *shortmsg.GetRegisterCodeReq, rsp *shortmsg.GetRegisterCodeRep) error {
	m:=db.GetRegisterShortMsgInfo(req.Phone);
	if(m.Phone==0){
		rsp.Code=0
		rsp.Tag=1
		return  nil
	}
	time10 := m.SendTime.Add(time.Duration(10) * time.Minute)
	if time10.Before(time.Now()){
		rsp.Code=0
		rsp.Tag=2 //验证码时间过期
		return  nil
	}
	rsp.Code= int32(m.Code)
	return nil
}

func (s *ShortMsgSvr) SendRegisterCode(ctx context.Context, req *shortmsg.SendRegisterCodeReq, rsp *shortmsg.SendRegisterCodeRep) error {
	m:=db.GetRegisterShortMsgInfo(req.Phone);
	if (m.Phone!=0) { //已经发送过验证码
		time10 := m.SendTime.Add(time.Duration(1) * time.Minute)
		if time10.Before(time.Now()){
			rsp.Tag=3 //请1分钟后在发送
			return  nil
		}
	}
	m.Phone= req.Phone
	m.SendTime=time.Now()
	code:=int(rand.Int31n(999999-100000)+100000)
	//发送短信

	//保存
	_,rows:= db.UpdateRegisterShortMsgInfo(code,time.Now(),0,req.Phone)
	if(rows==0){
		db.InsertRegisterShortMsgInfo(req.Phone,code,time.Now(),0)
	}
	return nil
}



func main() {
	fmt.Println("启动")
	core.Config(core.GetCurrentPath()+"\\svr_shortmsg.log",core.DebugLevel)
	core.Infof("启动成功")
	fmt.Println("配置文件")
	db2.G_dbmanage.Init();

	afterStart := func() error {
		core.Infof(consts.START_SUCCESS)
		return nil
	}
	service := micro.NewService(
		micro.Name(consts.API_NAME_SHORTMSG),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.AfterStart(afterStart),
	)
	// optionally setup command line usage
	service.Init()

	// Register Handlers
	shortmsg.RegisterShortmsgHandler(service.Server(), new(ShortMsgSvr))
	core.Infof("shortmsg:success")
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
