package daemon

import (
	"log"

	"github.com/sctlee/tcpx/protocol"
)

type DaemonConfig struct {
	ServerName string
	Logger     *log.Logger
	Pt         protocol.Protocol
}
