package v1

import (
	"mgpanel/internal/app/nginx"

	"github.com/gogf/gf/v2/frame/g"
)

// @ 创建EC2实例

type CreateEc2Req struct {
	g.Meta `path:"/create/ec2"  method:"post" summary:"Create an EC2 instance" tags:"aws"`
	nginx.NginxApp
}

type CreateEc2Res struct{}
