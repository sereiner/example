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
	input := make(map[string]interface{})
	if err := ctx.Request.GetJWT(&input); err != nil {
		return err
	}
	return input
}
