package controller

import "app/infra"

type BaseCtl struct{}

func NewBaseCtl() *BaseCtl {
	return &BaseCtl{}
}

// Middlewares implements infra.Controller.
func (b *BaseCtl) Middlewares() {
	panic("unimplemented")
}

// Routes implements infra.Controller.
func (b *BaseCtl) Routes() {
	panic("unimplemented")
}

var _ infra.Controller = (*BaseCtl)(nil)
