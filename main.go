package krakend_eureka

import (
	"fmt"
	"time"
)

func initPlugin() {
	fmt.Println("init eureka plugin")
	time.Now()
	for true {
		time.Sleep(time.Minute)
	}
}
