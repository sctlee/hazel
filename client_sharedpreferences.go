package tcpx

type SharedPreferences interface {
	Get(key string) (string, bool)
	Set(key string, value string)
	Del(key string)
}

type MapSharedPreferences struct {
	dict map[string]string
}

func NewSharePreferences(kind string) SharedPreferences {
	switch kind {
	case "map":
		return &MapSharedPreferences{make(map[string]string)}
	default:
		return &MapSharedPreferences{make(map[string]string)}
	}
}

func (self *MapSharedPreferences) Set(key string, value string) {
	self.dict[key] = value
}

func (self *MapSharedPreferences) Get(key string) (string, bool) {
	value, ok := self.dict[key]
	return value, ok
}

func (self *MapSharedPreferences) Del(key string) {
	delete(self.dict, key)
}
