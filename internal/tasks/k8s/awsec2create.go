package k8s

import (
	"fmt"
	"mgpanel/internal/file"
	"os/exec"
	"text/template"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

// 定义配置结构体
type Config struct {
	Region          string `v:"required" dc:"部署区域"`
	VpcID           string `v:"required" dc:"VPC ID"`
	SubnetID        string `v:"required" dc:"子网ID"`
	SecurityGroupID string `v:"required" dc:"安全组ID"`
	AmiID           string `v:"required" dc:"AMI镜像ID"`
	InstanceType    string `v:"required" dc:"实例类型（如t2.micro）"`
	KeyName         string `v:"required" dc:"SSH密钥对名称"`
	VolumeSize      int    `v:"required|min:8|max:1024" dc:"存储卷大小（GB）"`
	VolumeType      string `v:"required|in:gp2,gp3,io1,io2,st1,sc1" dc:"存储卷类型"`
	InstanceName    string `v:"required" dc:"实例名称"`
}

// TemplateGeneration 封装一个解析函数，传入目标的 tmpl 和 config 生成一个配置文件
func (c *Config) TemplateGeneration() error {

	// 创建输出文件
	outputFile := file.Creatfile("yaml/main.tf")
	if outputFile == nil {
		return fmt.Errorf("failed to create output file")
	}
	defer outputFile.Close() // 确保文件在使用后正确关闭

	// 解析模板文件
	tmpl, err := template.ParseFiles(AwsEc2creaTetmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template file: %w", err)
	}

	// 渲染模板到输出文件
	err = tmpl.Execute(outputFile, c)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	glog.New().Info(gctx.New(), "已成功渲染")

	return nil // 正常完成时返回 nil
}

// 创建configmap
func (c *Config) CreateConfigMap() error {
	// 使用 fmt.Sprintf 构造 --from-file 参数
	fromFileArg := fmt.Sprintf("--from-file=%s", maintf)

	// 创建 exec.Command 实例，将参数拆开
	cmd := exec.Command("kubectl", "create", "configmap", "terraform-config", fromFileArg)

	// 运行命令并捕获错误
	if err := cmd.Run(); err != nil {
		return err // 返回错误
	}

	return nil // 成功执行返回 nil
}
