package main

import (
	"github.com/vox-ai/nacos-go/example"
	"time"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-18 10:52
**/

func main() {
	example.ExampleServiceClient_RegisterServiceInstance()
	example.ExampleServiceClient_StartBeatTask()
	//example.ExampleServiceClient_Subscribe()
	time.Sleep(1000*time.Second)
}
