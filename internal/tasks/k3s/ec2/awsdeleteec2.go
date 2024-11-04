package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

// 定义配置结构体
type AwsEc2Delete struct {
	Region     string
	Instanceid []string
	Ctx        context.Context
}

func (d *AwsEc2Delete) Deleteec2() error {
	clinet, _ := NewClient(d.Ctx, d.Region)
	input := &ec2.TerminateInstancesInput{
		InstanceIds: d.Instanceid,
	}
	result, err := clinet.EC2Client.TerminateInstances(d.Ctx, input)
	if err != nil {
		glog.New().Error(d.Ctx, "无法终止 EC2 实例: ", err.Error())
		return err
	}
	for _, info := range result.TerminatingInstances {
		glog.New().Info(d.Ctx,
			gconv.String(*info.InstanceId),
			info.PreviousState,
			info.CurrentState,
		)
	}
	glog.New().Info(d.Ctx, "实例终止成功")
	return nil
}
