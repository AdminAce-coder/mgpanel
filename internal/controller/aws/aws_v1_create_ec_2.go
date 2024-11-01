package aws

import (
	"context"
	"mgpanel/api/aws/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) CreateEc2(ctx context.Context, req *v1.CreateEc2Req) (res *v1.CreateEc2Res, err error) {
	// 获取请求实例
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "无法获取请求实例")
	}

	// 记录请求中的参数
	glog.New().Info(ctx, "AmiID:", req.AmiID)

	// 获取 URL 查询参数 `someParam`
	paramValue := r.GetQuery("someParam")
	glog.New().Info(ctx, "someParam:", paramValue)

	// 写入 JSON 响应并直接 return
	r.Response.WriteJson(g.Map{
		"message": "paramValueparamValueparamValueparamValue",
	})
	return // 不再返回任何内容
}
