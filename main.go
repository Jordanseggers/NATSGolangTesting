package main

import (
	"runtime"
)

var appId int = 1
var flagName string = `App 1 Flag 1`
var server string = `nats://127.0.0.1:4222`
var userContext string = `375d39e6-9c3f-4f58-80bd-e5960b710295`
var sdkKey string = `myToken`

func main() {
	//natsClient.ConnectAndListen()

	runtime.Goexit()
}
