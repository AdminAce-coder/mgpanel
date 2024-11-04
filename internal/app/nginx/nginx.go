package nginx

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NginxApp 表示 Nginx 应用的部署配置
type NginxApp struct {
	Client    *kubernetes.Clientset
	Namespace string
	Name      string
	Replicas  int32
	Image     string
}

// 创建 int32 指针的辅助函数
func int32Ptr(i int32) *int32 { return &i }

// createDeployment 创建并部署 Nginx Deployment
func (n *NginxApp) createDeployment() error {
	// 定义 Deployment 规范
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      n.Name,
			Namespace: n.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(n.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": n.Name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": n.Name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  n.Name,
							Image: n.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// 使用 Kubernetes Client 创建 Deployment
	deploymentsClient := n.Client.AppsV1().Deployments(n.Namespace)
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("Deployment created: %q.\n", result.GetObjectMeta().GetName())
	return nil
}
