package ec2

import (
	"context"
	"fmt"
	"mgpanel/internal/consts"
	w "mgpanel/internal/tasks/k3s/ec2/ec2workspace"
	"mgpanel/internal/tasks/k3s/ec2/tmp"
	"path/filepath"
	"text/template"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

// 定义配置结构体
type AwsEc2Create struct {
}

// TemplateGeneration 封装一个解析函数，传入目标的 tmpl 和 config 生成一个配置文件
func (c *AwsEc2Create) TemplateGeneration(awsec2config *AwsEc2Config) error {
	// 创建输出文件
	w := w.CreateTf{}
	outputFile := w.GetWorkspaceAbs().Creatfile(consts.AwsEC2createOut)
	if outputFile == nil {
		return fmt.Errorf("failed to create output file")
	}
	defer outputFile.Close() // 确保文件在使用后正确关闭

	// 获取渲染模板路径
	fl := tmp.GetTmpDir()

	AwsEc2creaTetmplPath := filepath.Join(fl, consts.AwsEc2creaTetmplPath)

	// 解析模板文件
	glog.New().Info(gctx.New(), "解析文件是：", AwsEc2creaTetmplPath)
	tmpl, err := template.ParseFiles(AwsEc2creaTetmplPath)
	if err != nil {
		glog.New().Error(gctx.New(), "解析渲染模板失败")
		return err
	}

	err = tmpl.Execute(outputFile, awsec2config)
	if err != nil {
		glog.New().Error(gctx.New(), "渲染模板失败")
		return err
	}
	glog.New().Info(gctx.New(), "已成功渲染")

	return nil // 正常完成时返回 nil
}

//  拿到模板直接创建实例

func (c *AwsEc2Create) CreateEC2(awsec2config *AwsEc2Config) error {
	// 解析模板
	c.TemplateGeneration(awsec2config)
	// 执行 apply 操作
	wl := w.CreateTf{}
	tfclinet, err := GetTerraformInstance(wl.GetWorkspaceAbs().CreatefilePath)
	if err != nil {
		glog.New().Error(gctx.New(), "tf客户端创建失败")
	}
	if tfclinet == nil {
		glog.New().Error(gctx.New(), "tfclinet 为空")
	}
	//if err := tf.Apply(context.Background(), tfexec.Target("aws_instance.web_server")); err != nil {
	//	fmt.Printf("Failed to apply Terraform configuration for target resource: %v\n", err)
	//	return err
	//}
	//return nil
	// 执行 apply 操作（去掉 Target 参数）
	glog.New().Info(gctx.New(), "正在创建EC2实例...")
	if err := tfclinet.Apply(context.Background()); err != nil {
		fmt.Printf("Failed to apply Terraform configuration: %v\n", err)
		return err
	}
	return nil
}
