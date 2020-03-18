package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"go-web-lz/conf"
	"go-web-lz/route"
	"log"
	"syscall"
)

func main() {
	port := fmt.Sprintf("localhost:%v", conf.Sysconfig.Port)
	server := endless.NewServer(port, route.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getegid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err : %v", err)
	}
}
