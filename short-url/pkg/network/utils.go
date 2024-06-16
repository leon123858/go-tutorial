package network

import (
	"context"
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const ns = "default" // 替換為實際的命名空間
const sn = "app"     // 替換為實際的服務名稱

func GetServerIP() (string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", err
	}

	namespace := ns
	serviceName := sn
	service, err := clientset.CoreV1().Services(namespace).Get(context.Background(), serviceName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	var externalURL string
	if service.Spec.Type == "LoadBalancer" {
		ingress := service.Status.LoadBalancer.Ingress
		if len(ingress) > 0 {
			externalURL = ingress[0].IP
			if ingress[0].Hostname != "" {
				externalURL = ingress[0].Hostname
			}
		}
	}

	if externalURL == "" {
		return "", errors.New("no external IP found")
	}
	fmt.Println("External URL:", externalURL)

	return externalURL, nil
}
