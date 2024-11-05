package ec2

import (
	"context"
	"fmt"
	"sync"

	"github.com/hashicorp/terraform-exec/tfexec"
)

var (
	// 全局变量，存储已初始化的 Terraform 实例
	tfInstance *tfexec.Terraform
	// 用于确保初始化过程的线程安全
	once sync.Once
	// 错误变量用于存储初始化过程中遇到的错误
	initErr error
)

// 初始化或获取 Terraform 实例
func GetTerraformInstance(workDir string) (*tfexec.Terraform, error) {
	// 使用 sync.Once 确保只执行一次初始化
	once.Do(func() {
		// 初始化 Terraform 实例
		tf, err := tfexec.NewTerraform(workDir, "terraform")
		if err != nil {
			initErr = fmt.Errorf("Error initializing Terraform: %v", err)
			return
		}

		// 设置初始化选项
		initOpts := []tfexec.InitOption{
			tfexec.PluginDir("C:\\Users\\ThinkPad\\AppData\\Roaming\\terraform.d\\plugins"),
			//tfexec.GetPlugins(false), // 如果希望禁用插件下载
		}

		// 初始化 Terraform，并传递选项
		if err := tf.Init(context.Background(), initOpts...); err != nil {
			fmt.Printf("Failed to initialize Terraform: %v\n", err)
			return
		}

		tfInstance = tf // 初始化成功，赋值给全局变量
	})

	return tfInstance, initErr
}
