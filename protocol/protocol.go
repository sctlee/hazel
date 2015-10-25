package protocol

type Protocol interface {
	Marshal(str string) map[string]string
}
