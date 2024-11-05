package aws

import (
	"context"
	"mgpanel/api/aws/v1"
	"mgpanel/internal/tasks/k3s/ec2"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) DeleteEc2(ctx context.Context, req *v1.DeleteEc2Req) (res *v1.DeleteEc2Res, err error) {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "无法获取请求实例")
	}

	delete := ec2.AwsEc2Delete{
		Ctx: ctx,
	}
	// 将请求中的 JSON 数据解析到 config 中
	if err := r.Parse(&delete); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInvalidParameter, err, "请求参数解析失败")
	}
	glog.New().Info(ctx, "解析后的删除结构体为", delete)
	if delete.DeleteEc2(); err != nil {
		glog.New().Error(ctx, "实例删除失败:", err)
		r.Response.WriteJson(g.Map{
			"message": "EC2实例删除失败",
		})
	}
	r.Response.WriteJson(g.Map{
		"message": "EC2删除成功",
	})
	return
}
