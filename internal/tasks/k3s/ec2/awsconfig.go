package ec2

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type AwsEc2Config struct {
	Region        string
	Instance      Instance       // EC2实例
	Vpc           Vpc            // VPC网络
	AwsAccessesCg *AwsAccessesCg //访问秘钥
	Operation     operation      //  所有操作
}
type option func(a *AwsEc2Config)

func NewAwsEc2Config(op operation, option ...option) *AwsEc2Config {
	awsec2config := &AwsEc2Config{
		AwsAccessesCg: newAwsAccess(),
		Operation:     op,
	}
	for _, option := range option {
		option(awsec2config)
	}
	return awsec2config
}

func WithRegion(region string) option {
	return func(a *AwsEc2Config) {
		a.Region = region
	}
}

func withInstance(i Instance) option {
	return func(a *AwsEc2Config) {
		a.Instance = i
	}
}
func WithVpc(v Vpc) option {
	return func(a *AwsEc2Config) {
		a.Vpc = v
	}
}

type Instance struct {
	Istanceid    string
	AmiID        string
	InstanceType string
	VolumeSize   int    `v:"required|min:8|max:1024" dc:"存储卷大小（GB）"`
	VolumeType   string `v:"required|in:gp2,gp3,io1,io2,st1,sc1" dc:"存储卷类型"`
	InstanceName string `v:"required" dc:"实例名称"`
}

type Vpc struct {
	Vpcid    string `dc:"Vpcid"`
	SubnetID string `dc:"SubnetID"` // 子网ID
	Sg       SecurityGroup
}

type SecurityGroup struct {
	SecurityGroupID string `dc:"SecurityGroupID"`
	KeyName         string `dc:"KeyName"`
}

type AwsAccessesCg struct {
	AccessedId string `v:"required" dc:"访问ID"`
	SecretKey  string `v:"required" dc:"访问秘钥"`
}

type operation interface {
	TemplateGeneration(aws *AwsEc2Config) error
	CreateEC2() error
}

func newAwsAccess() *AwsAccessesCg {
	var ctx = gctx.New()
	acessid, _ := g.Cfg(Configpath).Get(ctx, "aws.acessId")
	accessKey, _ := g.Cfg(Configpath).Get(ctx, "aws.accessKey")
	fmt.Printf("id:%s,\n", acessid.Strings())
	fmt.Printf("accessKey:%s,\n", accessKey.Strings())
	return &AwsAccessesCg{
		AccessedId: acessid.Strings()[0],
		SecretKey:  accessKey.Strings()[0],
	}
}
