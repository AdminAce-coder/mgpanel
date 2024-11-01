package v1

import (
	"mgpanel/internal/tasks/k8s"

	"github.com/gogf/gf/v2/frame/g"
)

// @ 创建EC2实例

type CreateEc2Req struct {
	g.Meta `path:"/create/ec2"  method:"post" summary:"Create an EC2 instance" tags:"aws"`
	k8s.Config
}

type CreateEc2Res struct{}
