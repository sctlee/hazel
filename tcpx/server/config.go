package server

import (
	"github.com/sctlee/tcpx/protocol"
)

type ServerConfig struct {
	ServerName string
	Port       string
	Pt         protocol.Protocol
}
