package daemon

import (
	"github.com/sctlee/hazel/daemon/service"
)

type ServiceManager struct {
	daemon   *Daemon
	Services map[string]*service.Service
}

func (self *ServiceManager) RegisterService(s *service.Service) {
	self.Services[s.Name] = s
	go s.Listen()
}

func (self *ServiceManager) GetService(name string) (s *service.Service, ok bool) {
	s, ok = self.Services[name]
	return s, ok
}

func NewServiceManager(d *Daemon) *ServiceManager {
	return &ServiceManager{
		daemon:   d,
		Services: make(ServiceList),
	}
}
