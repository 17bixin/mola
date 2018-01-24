package front

var Front2_file = `

package main

import (
	"os"
	"os/signal"
	"syscall"

	"common/server"

	"common/config"
	"github.com/17bixin/gobase/log"
	"routers"
)

func main() {

	log.Setlogfile("pay.log")
	log.Info(log.Fields{"app": "exec args", "args": os.Args})

	defer func() {
		if error := recover(); error != nil {
			log.Fatal(log.Fields{"panic": error})
			exit(-1)
		}
	}()

	go func() {
		singals := make(chan os.Signal)
		signal.Notify(singals,
			syscall.SIGTERM,
			syscall.SIGINT,
			syscall.SIGKILL,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)
		<-singals
		exit(0)
	}()
	config.Init("src/apps/front/config.json")

	go func() {
		example.Run()
	}()

	routers.Init()

	server.Run()
}

func exit(status int) {
	example.Stop()
	server.Stop()
	log.Info(log.Fields{"app": status})
	os.Exit(status)
}
`
