package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

// AwsClient 包含EC2和Lightsail客户端实例
type AwsClient struct {
	EC2Client       *ec2.Client
	LightsailClient *lightsail.Client
}

// NewClient 初始化一个包含EC2和Lightsail客户端的新AwsClient
func NewClient(ctx context.Context, region string) (*AwsClient, error) {
	// 加载默认配置
	accessid, _ := g.Cfg().Get(ctx, "aws.acessId")
	accessKey, _ := g.Cfg().Get(ctx, "aws.accessKey")
	id := accessid.Strings()[0]
	key := accessKey.Strings()[0]
	glog.New().Info(ctx, "aws client id:", id, "key:", key)
	// 创建一个自定义的凭证提供程序
	creds := credentials.NewStaticCredentialsProvider(id, key, "")

	// 使用自定义凭证和区域配置加载 AWS 配置
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(creds),
	)
	if err != nil {
		glog.New().Error(ctx, "无法加载AWS配置: ", err.Error())
		return nil, err
	}

	// 创建EC2客户端
	ec2Client := ec2.NewFromConfig(cfg)

	// 创建Lightsail客户端
	lightsailClient := lightsail.NewFromConfig(cfg)

	// 返回包含EC2和Lightsail客户端的AwsClient实例
	return &AwsClient{
		EC2Client:       ec2Client,
		LightsailClient: lightsailClient,
	}, nil
}

// 示例：列出EC2实例
func (client *AwsClient) ListEC2Instances(ctx context.Context) error {
	input := &ec2.DescribeInstancesInput{}

	// 调用 EC2 API 获取实例列表
	result, err := client.EC2Client.DescribeInstances(ctx, input)
	if err != nil {
		glog.New().Error(ctx, "无法获取EC2实例列表: ", err.Error())
		return err
	}

	// 打印实例信息
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			glog.New().Infof(ctx, "EC2实例ID: %s, 状态: %s", gconv.String(instance.InstanceId), gconv.String(instance.State.Name))
		}
	}

	return nil
}

// 示例：列出Lightsail实例
func (client *AwsClient) ListLightsailInstances(ctx context.Context) error {
	input := &lightsail.GetInstancesInput{}

	// 调用 Lightsail API 获取实例列表
	result, err := client.LightsailClient.GetInstances(ctx, input)
	if err != nil {
		glog.New().Error(ctx, "无法获取Lightsail实例列表: ", err.Error())
		return err
	}

	// 打印Lightsail实例信息
	for _, instance := range result.Instances {
		glog.New().Infof(ctx, "Lightsail实例名称: %s, 状态: %s", gconv.String(*instance.Name), gconv.String(instance.State.Name))
	}

	return nil
}
