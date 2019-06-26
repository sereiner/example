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
func (u *QueryHandler) Handle(ctx *context.Context) (r interface{}) {
	sdi, sdc := ctx.Request.GetSharding()
	ctx.Log.Infof("%s %d/%d", ctx.Service, sdi, sdc)
	return "success"
}
