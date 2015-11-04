package protocol

import (
	"strings"
)

type SimpleProtocol struct {
	// chatroom.list|key:value;key:value
}

func (self *SimpleProtocol) Marshal(str string) map[string]string {
	m := make(map[string]string)

	firstPart := strings.Split(str[0:strings.Index(str, "|")], ".")
	secondPart := strings.Split(strings.TrimPrefix(str[strings.Index(str, "|"):], "|"), ";")
	m["feature"] = firstPart[0]
	m["command"] = firstPart[1]

	if secondPart[0] != "" {
		for _, param := range secondPart {
			p := strings.Split(param, ":")
			m[p[0]] = p[1]
		}
	}

	return m
}

func (self *SimpleProtocol) UnMarshal(params map[string]string) string {
	return ""
}
