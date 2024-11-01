package k8s

import (
	"testing"
)

func TestTemplateGeneration(t *testing.T) {
	config := Config{
		Region:          "us-east-1",
		VpcID:           "vpc-0464b5baebfcaa821",
		SubnetID:        "subnet-0711e346568ae125b",
		SecurityGroupID: "sg-0d51cce9c33fa3221",
		AmiID:           "ami-0866a3c8686eaeeba",
		InstanceType:    "t2.micro",
		KeyName:         "zxb",
		VolumeSize:      20,
		VolumeType:      "gp3",
		InstanceName:    "MyEC2Instance",
	}
	path := "awcec2.tmpl"
	// 调用 TemplateGeneration，并检查是否出现错误
	if err := config.TemplateGeneration(path); err != nil {
		t.Fatalf("Template generation failed: %v", err)
	}
}
