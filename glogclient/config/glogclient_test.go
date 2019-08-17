package config

import (
	"fmt"
	"testing"
)

func TestGetElasticConfig(t *testing.T) {
	config:=GetGLogServer()
	fmt.Println(config)
}
