package ec2

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestTemplateGeneration(t *testing.T) {

	in := Instance{
		InstanceType: "t2.micro",
		AmiID:        "ami-0866a3c8686eaeeba",
		VolumeSize:   8,
		VolumeType:   "gp3",
		InstanceName: "MyEC2Instance",
		Istanceid:    "i-088b668cc6ed86f45",
	}
	sg := SecurityGroup{
		KeyName:         "zxb",
		SecurityGroupID: "sg-0d51cce9c33fa3221",
	}
	vpc := Vpc{
		Vpcid:    "vpc-0464b5baebfcaa821",
		SubnetID: "subnet-0711e346568ae125b",
		Sg:       sg,
	}
	t.Logf("test,testtesttesttest")
	var newcrete = NewAwsEc2Config(&AwsEc2Create{}, withInstance(in), WithRegion("us-east-1"), WithVpc(vpc))
	// 调用 TemplateGeneration，并检查是否出现错误
	if newcrete == nil {
		t.Fatalf("配置文件是空")
	}

	if err := newcrete.Operation.TemplateGeneration(newcrete); err != nil {
		t.Fatalf("Template generation failed: %v", err)
	}
	t.Logf("正在创建实例")
	err := newcrete.Operation.CreateEC2(newcrete)
	if err != nil {
		t.Fatalf("CreateEC2 失败: %v", err)
	}
	t.Logf("实例创建成功")

}

func TestDeleteEC2(t *testing.T) {
	ctx := gctx.New()
	Instanceids := []string{"i-088b668cc6ed86f45"}
	delete := AwsEc2Delete{
		Region:     "us-east-1",
		Instanceid: Instanceids,
		Ctx:        ctx,
	}
	err := delete.DeleteEc2()
	if err != nil {
		t.Fatalf("Deleteec2 failed: %v", err)
	}

}
