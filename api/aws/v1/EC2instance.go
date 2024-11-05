package v1

import (
	"mgpanel/internal/tasks/k3s/ec2"

	"github.com/gogf/gf/v2/frame/g"
)

// @ 创建EC2实例

type CreateEc2Req struct {
	g.Meta `path:"/create/ec2"  method:"post" summary:"Create an EC2 instance" tags:"aws"`
	ec2.AwsEc2Config
}

type CreateEc2Res struct{}

//@ 删除EC2

type DeleteEc2Req struct {
	g.Meta `path:"/delete/ec2"  method:"post" summary:"Delete an EC2 instance" tags:"aws"`
	ec2.AwsEc2Delete
}

type DeleteEc2Res struct{}
