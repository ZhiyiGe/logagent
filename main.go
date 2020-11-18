package main

import (
	"fmt"
	"github.com/996-to-ICU/MystudyCode/log/logagent/conf"
	"github.com/996-to-ICU/MystudyCode/log/logagent/kafka"
	"github.com/996-to-ICU/MystudyCode/log/logagent/taillog"
	"gopkg.in/ini.v1"
	"time"
)

var (
	cfg =new(conf.AppConf)
)

func run() {
	for {
		select {
		case line := <-taillog.Readlog():
			_ = kafka.Sendmessage(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
func main() {
	//cfg,err:=ini.Load("./conf/conf.ini")
	err:=ini.MapTo(cfg,"C:/GOPATH/src/github.com/996-to-ICU/MystudyCode/log/logagent/conf/conf.ini")
	if err != nil {
		fmt.Println("load ini failed,err",err)
		return
	}
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Println("init kafka failed,err:", err)
		return
	}
	err = taillog.Init(cfg.TailLog.FileName)
	if err != nil {
		fmt.Println("init taillog failed,err:", err)
		return
	}
	run()
}
