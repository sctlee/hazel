package daemon

const (
	EVENT_CLIENT_QUIT = "event of client quiting"
)

type ClientQuitEventer interface {
	OnClientQuit(cid string)
}

func (self *ServiceManager) TriggerEvent(eventType string, params ...string) {
	for _, s := range self.Services {
		switch eventType {
		case EVENT_CLIENT_QUIT:
			if e, ok := interface{}(s.GetOriginalService()).(ClientQuitEventer); ok {
				e.OnClientQuit(params[0])
			}
		}
	}
}
