package v1

import (
	"mgpanel/internal/tasks/k3s/ec2"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateDpReq struct {
	g.Meta `path:"/create/dp"  method:"post" summary:"Create an deployment" tags:"k8s"`
	ec2.AwsEc2Create
}

type CreateDpRes struct{}
