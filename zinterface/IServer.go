package zinterface

// IServer
type IServer interface {
	// Start
	Start()
	// Stop
	Stop()
	// Serve
	Serve()
	// Router
	AddRouter(router IRouter)
}
