package aws

import (
	"context"
	"mgpanel/api/aws/v1"
	"mgpanel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) GetALL(ctx context.Context, req *v1.GetALLReq) (res *v1.GetALLRes, err error) {
	r := g.RequestFromCtx(ctx)

	result := service.Getinfo().GetAll(ctx)
	if result == " " {
		glog.New().Info(ctx, "数据为空...")
	}
	r.Response.WriteJson(result)
	return
}
