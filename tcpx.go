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
	serverConfig := &server.ServerConfig{
		ServerName: config.ServerName,
		Port:       config.Port,
		Pt:         pt,
	}
	server := server.NewServer(serverConfig)

	serveErrWait := make(chan error)
	go func() {
		if err := server.Start(); err != nil {
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
		d.SrvManager.RegisterService(s)
	}

	server.AcceptConnections(d)

	err := <-serveErrWait
	logger.Println(err)
	fmt.Println("server close")
}

func SendMessage(msg *message.Message) {
	err := d.MsgManager.PutMessage(msg)
	if err != nil {
		logger.Println(err)
		fmt.Println(err)
	}
}
