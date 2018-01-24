package server

var server3_file = `

package server

import (
	"common/proto"
	"controller"

	"github.com/17bixin/gobase/grpcserver"
)

var (
	s           = grpcserver.New()
	rpcRegister = func() {
		//proto.Register{}Server(s.RPCSvr, controller.New{}())
	}
	httpRegisters = []grpcserver.HTTPRegister{
		//proto.Register{}HandlerFromEndpoint,
	}
)

func Run() {
	s.SetPort(8001)
	s.SetRPCRegister(rpcRegister)
	s.SetHTTPRegisters(httpRegisters)

	s.Run()
}

func Stop() {
	s.Stop()
}
`
