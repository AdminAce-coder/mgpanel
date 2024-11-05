package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

//@ 查询数据

type GetALLReq struct {
	g.Meta `path:"/get/ec2info"  method:"get" summary:"获取EC2实例信息" tags:"aws"`
}

type GetALLRes struct {
}
