package server

import (
	"github.com/sctlee/hazel/protocol"
)

type ServerConfig struct {
	ServerName string
	Port       string
	Pt         protocol.Protocol
}
