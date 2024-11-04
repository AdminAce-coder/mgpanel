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

func (c *ControllerV1) CreateEc2(ctx context.Context, req *v1.CreateEc2Req) (res *v1.CreateEc2Res, err error) {
	// 获取请求实例
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "无法获取请求实例")
	}

	// 创建一个 config 结构体实例
	var config ec2.Config

	// 将请求中的 JSON 数据解析到 config 中
	if err := r.Parse(&config); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInvalidParameter, err, "请求参数解析失败")
	}

	// 记录解析后的参数内容
	glog.New().Info(ctx, "Parsed Config:", config)
	glog.New().Info(ctx, "terrform模板生成成功")
	// 执行其他逻辑，处理配置后的内容
	r.Response.WriteJson(g.Map{
		"message": "terrform模板生成成功",
	})

	return
}
