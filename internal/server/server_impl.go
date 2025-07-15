package server

import (
	"github.com/sawdustofmind/eth-balance-proxy/internal/service"
)

var _ ServerInterface = (*ServerImpl)(nil)

type ServerImpl struct {
	bg *service.BalanceGetter
}

func NewServerImpl(bg *service.BalanceGetter) *ServerImpl {
	return &ServerImpl{
		bg: bg,
	}
}
