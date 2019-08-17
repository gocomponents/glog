package glogclient

import (
	"context"
	"fmt"
	"glog/glogclient/config"
	"glog/proto"
	"google.golang.org/grpc"
)

var logCh = make(chan *proto.Log, 500)

func init()  {
	go consume()
}

func consume()  {
	conn, err := grpc.Dial(config.GetGLogServer(), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := proto.NewLogStashClient(conn)
	for {
		log, ok := <-logCh
		if ok {
			go func(log *proto.Log) {
				_,err:=client.Send(context.Background(), log)
				if nil!=err {
					fmt.Println(err)
				}
			}(log)
		}
	}
}

func Send(log *proto.Log)  {
	logCh<-log
}
