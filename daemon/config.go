package daemon

import (
	"log"

	"github.com/sctlee/hazel/protocol"
)

type DaemonConfig struct {
	ServerName string
	Logger     *log.Logger
	Pt         protocol.Protocol
}
