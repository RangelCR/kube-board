package k8s

import (
	"context"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// KubeClient is a Kubernetes client.
type KubeClient struct {
	client *kubernetes.Clientset
}

// NewClient creates a new Kubernetes client.
func NewClient() (*KubeClient, error) {

	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil { // Provavelmente está sendo executado in cluster

		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}

	client, err := kubernetes.NewForConfig(config)

	return &KubeClient{client: client}, err
}

func (client *KubeClient) ListNamespaces() (*v1.NamespaceList, error) {
	return client.client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
}

// ListDeployments lists all deployments in the given namespace.
func (client *KubeClient) ListDeployments(namespace string) (*appsv1.DeploymentList, error) {
	if namespace == "" {
		namespace = "default"
	}
	return client.client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
}

// ListPods lists all pods in the given namespace.
func (client *KubeClient) ListPods(namespace string) (*v1.PodList, error) {

	if namespace == "" {
		namespace = "default"
	}
	return client.client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (client *KubeClient) ListReplicaSets(namespace string) (*appsv1.ReplicaSetList, error) {

	if namespace == "" {
		namespace = "default"
	}

	return client.client.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (client *KubeClient) ListDaemonSets(namespace string) (*appsv1.DaemonSetList, error) {
	if namespace == "" {
		namespace = "default"
	}

	return client.client.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (client *KubeClient) ListStatefulSets(namespace string) (*appsv1.StatefulSetList, error) {
	if namespace == "" {
		namespace = "default"
	}

	return client.client.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (client *KubeClient) ListJobs(namespace string) (*batchv1.JobList, error) {
	if namespace == "" {
		namespace = "default"
	}

	return client.client.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (client *KubeClient) ListCronJobs(namespace string) (*batchv1.CronJobList, error) {
	if namespace == "" {
		namespace = "default"
	}

	return client.client.BatchV1().CronJobs(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (client *KubeClient) ListServices(namespace string) (*v1.ServiceList, error) {
	if namespace == "" {
		namespace = "default"
	}

	return client.client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
}
