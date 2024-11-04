package nginx

import (
	"mgpanel/internal/app"
	"testing"
)

func TestNginx(t *testing.T) {
	cl := app.NewK8sClient()
	ng := NginxApp{
		Client:    cl,
		Namespace: "default",
		Name:      "testnginx",
		Replicas:  1,
		Image:     "registry.cn-chengdu.aliyuncs.com/aliyunzzzxb/nginx:latest",
	}
	err := ng.createDeployment() // 检查 createDeployment 的错误
	if err != nil {
		t.Fatalf("Failed to create Nginx deployment: %v", err)
	}
}

func TestGetconfig(t *testing.T) {
	path, err := app.Getconfig()
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
	}
	t.Logf("Config path: %s", path)

	client := app.NewK8sClient()
	if client == nil {
		t.Fatalf("Failed to create Kubernetes client")
	}
}
