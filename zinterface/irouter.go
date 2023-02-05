package zinterface

type IRouter interface {
	// Pre Hook
	PreHandle(request IRequest)
	// Main Hook
	Handle(request IRequest)
	// Post Hook
	PostHandle(request IRequest)
}
