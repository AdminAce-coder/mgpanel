package ec2

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hashicorp/terraform-exec/tfexec"
)

// 初始化terr
func init() {
	// 设置 Terraform 路径和工作目录
	tf, err := tfexec.NewTerraform(WorkDir, "terraform")
	if err != nil {
		fmt.Printf("Error initializing Terraform: %v\n", err)
		return
	}

	// 初始化 Terraform
	if err := tf.Init(context.Background()); err != nil {
		fmt.Printf("Failed to initialize Terraform: %v\n", err)
		return
	}
	glog.Info(gctx.New(), "Terraform initialized,初始化成功....")
}
