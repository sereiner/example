package order

import (
	"fmt"

	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
)

type BindHandler struct {
	container component.IContainer
}

func NewBindHandler(container component.IContainer) (u *BindHandler) {
	return &BindHandler{container: container}
}
func (u *BindHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Response.SetStatus(508)
	return fmt.Errorf("发生异常")
}
