package glogclient

import (
	"glog/proto"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	log:=proto.Log{
		App:        "test",
		Module:     "consume",
		Level:      proto.Log_Info,
		TraceId:    "123",
		Message:    []byte("456"),
		Exception:  nil,
		UserIp:     "192.168.11.11",
		ExecTime:   12,
		CreateTime: time.Now().Add(time.Duration(10) * time.Millisecond).Format("2006-01-02 15:04:05"),
	}
	Send(&log)
	<-make(chan bool)
}
