package znet

import "github.com/alfmunny/willego/zinterface"

// BaseRouter ...
// Embedded to router, so that PreHandle and PostHandle can be empty
type BaseRouter struct{}

// PreHandle ...
func (br *BaseRouter) PreHandle(request zinterface.IRequest) {

}

// Handle ...
func (br *BaseRouter) Handle(request zinterface.IRequest) {

}

// PostHandle ...
func (br *BaseRouter) PostHandle(request zinterface.IRequest) {

}
