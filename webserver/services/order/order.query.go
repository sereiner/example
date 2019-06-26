package order

import (
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
)

type QueryHandler struct {
	container component.IContainer
}

func NewQueryHandler(container component.IContainer) (u *QueryHandler) {
	return &QueryHandler{container: container}
}
func (u *QueryHandler) GetHandle(ctx *context.Context) (r interface{}) {
	return "get.success" + ctx.Request.Setting.GetString("db")
}
func (u *QueryHandler) Handle(ctx *context.Context) (r interface{}) {
	return "success"
}
