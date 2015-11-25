package tcpx

import (
	"fmt"

	"github.com/sctlee/tcpx/daemon"
	"github.com/sctlee/tcpx/daemon/message"
	"github.com/sctlee/tcpx/daemon/service"
	"github.com/sctlee/tcpx/tcpx/server"
)

var serverName string
var d *daemon.Daemon

func MainDaemon(config *Config, services ...*service.Service) {
	server := server.NewServer()

	serveErrWait := make(chan error)
	go func() {
		if err := server.Start(config.Port); err != nil {
			serveErrWait <- err
		}
	}()

	daemonConfig := &daemon.DaemonConfig{
		ServerName: config.ServerName,
		Pt:         pt,
		Logger:     logger,
	}

	d = daemon.NewDaemon(daemonConfig)
	for _, s := range services {
		d.RegisterService(s)
	}

	server.AcceptConnections(d)

	err := <-serveErrWait
	logger.Println(err)
	fmt.Println("server close")
}

func SendMessage(msg *message.Message) {
	d.MsgManager.PutMessage(msg)
}
