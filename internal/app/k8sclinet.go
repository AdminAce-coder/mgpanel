package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/os/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// NewK8sClient 创建并返回 Kubernetes 客户端
func NewK8sClient() *kubernetes.Clientset {
	path, err := Getconfig()
	if err != nil {
		glog.Error(context.TODO(), fmt.Sprintf("Failed to get config path: %v", err))
		return nil
	}

	glog.Info(context.TODO(), fmt.Sprintf("Config path is: %s", path))
	client, err := GetClientSet(path)
	if err != nil {
		glog.Error(context.TODO(), fmt.Sprintf("Failed to create Kubernetes client: %v", err))
		return nil
	}

	glog.Info(context.TODO(), "Kubernetes client created successfully")
	return client
}

// GetClientSet 创建K8s客户端
func GetClientSet(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build config from flags: %w", err)
	}

	glog.Info(context.TODO(), "Client configuration loaded successfully")
	return kubernetes.NewForConfig(config)
}

// homeDir 获取用户主目录路径
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows 平台
}

// Getconfig 从系统获取配置文件路径
func Getconfig() (string, error) {
	// 获取用户主目录路径的 kubeconfig
	h := homeDir()
	kubeconfig := filepath.Join(h, ".kube", "config")

	if _, err := os.Stat(kubeconfig); os.IsNotExist(err) {
		glog.Info(context.TODO(), "kubeconfig not found in ~/.kube/config, trying backup config path")

		// 获取 manifest/config/k3sconfig.ec2workspace 的绝对路径
		kubeconfig, _ = filepath.Abs(filepath.Join("..", "..", "..", "manifest", "config", "k3sconfig.ec2workspace"))

		if _, err := os.Stat(kubeconfig); os.IsNotExist(err) {
			glog.Error(context.TODO(), fmt.Sprintf("Backup kubeconfig also not found at %s", kubeconfig))
			return "", err // 返回错误
		}
	}

	return kubeconfig, nil
}
