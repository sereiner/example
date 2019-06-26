package order

import (
	"github.com/sereiner/parrot/component"
	"github.com/sereiner/parrot/context"
)

type Input struct {
	ID   string `form:"id" json:"id" valid:"int,required"` //绑定输入参数，并验证类型否是否必须输入
	Name string `form:"name" json:"name"`
}
type BindHandler struct {
	container component.IContainer
}

func NewBindHandler(container component.IContainer) (u *BindHandler) {
	return &BindHandler{container: container}
}
func (u *BindHandler) Handle(ctx *context.Context) (r interface{}) {
	var input Input
	// 请求参数绑定
	if err := ctx.Request.Bind(&input); err != nil {
		return err
	}
	return input
}
