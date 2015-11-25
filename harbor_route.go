package tcpx

var harborAction = &HarborAction{}

var HarborRouter = map[string]RouteFun{
	"pop": harborAction.Pop,
}
