package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/usermanager"
)

func main() {
	sigChannel := make(chan os.Signal)
	exitChannel := make(chan bool)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go usermanager.StartGrpcTransport(uManagerService)
	go usermanager.StartHTTPTransport(uManagerService)
	go func(){
		sig := <- sigChannel
		log.Println(sig)
		exitChannel <- true
	}()
	
	<-exitChannel
	close(exitChannel)
	log.Println("exit")
}
