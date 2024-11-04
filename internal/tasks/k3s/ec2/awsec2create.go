package ec2

import (
	"context"
	"fmt"
	"mgpanel/internal/file"
	"text/template"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hashicorp/terraform-exec/tfexec"
)

// 定义配置结构体
type AwsEc2Create struct {
}

// TemplateGeneration 封装一个解析函数，传入目标的 tmpl 和 config 生成一个配置文件
func (c *AwsEc2Create) TemplateGeneration(awsec2config *AwsEc2Config) error {
	// 创建输出文件
	outputFile := file.Creatfile(AwsEC2createOut)
	if outputFile == nil {
		return fmt.Errorf("failed to create output file")
	}
	defer outputFile.Close() // 确保文件在使用后正确关闭

	// 解析模板文件
	tmpl, err := template.ParseFiles(AwsEc2creaTetmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template file: %w", err)
	}

	err = tmpl.Execute(outputFile, awsec2config)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	glog.New().Info(gctx.New(), "已成功渲染")

	return nil // 正常完成时返回 nil
}

//  拿到模板直接创建实例

func (c *AwsEc2Create) CreateEC2() error {
	// 执行 apply 操作
	tf, err := tfexec.NewTerraform(WorkDir, "terraform")
	if err != nil {
		glog.New().Error(gctx.New(), "tf客户端创建失败")
	}
	if err := tf.Apply(context.Background(), tfexec.Target("aws_instance.web_server")); err != nil {
		fmt.Printf("Failed to apply Terraform configuration for target resource: %v\n", err)
		return err
	}
	return nil
}
