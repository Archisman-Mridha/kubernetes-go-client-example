package main

import (
	"context"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	podReplicas int32= 2
)

func main( ) {
	var kubeconfigFileLocation string = "C:/Users/archi/.kube/config"

	kubeConfig, error := clientcmd.BuildConfigFromFlags("", kubeconfigFileLocation)
	panicFilter(error)

	// client set consists of the clients prepared for all the kubernetes api group
	clientSet, error := kubernetes.NewForConfig(kubeConfig)
	panicFilter(error)

	appsV1APIClient := clientSet.AppsV1( )

	//* deployment definition
	deploymentDefinition := &appsV1.Deployment {

		ObjectMeta: metaV1.ObjectMeta {

			Name: "demo-deployment",
			Namespace: coreV1.NamespaceDefault,
		},

		Spec: appsV1.DeploymentSpec {
			Replicas: &podReplicas,

			Selector: &metaV1.LabelSelector {
				MatchLabels: map[string]string {
					"app": "demo",
				},
			},

			Template: coreV1.PodTemplateSpec {

				ObjectMeta: metaV1.ObjectMeta {
					Labels: map[string]string {
						"app": "demo",
					},
				},

				Spec: coreV1.PodSpec {
					Containers: [ ]coreV1.Container {
						{
							Name:  "web",
							Image: "nginx:1.12",
							ImagePullPolicy: coreV1.PullIfNotPresent,
							Ports: [ ]coreV1.ContainerPort{
								{
									Name: "http",
									Protocol: coreV1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// create the deployment
	createdDeployment, error := appsV1APIClient.Deployments(coreV1.NamespaceDefault).Create(context.Background( ), deploymentDefinition, metaV1.CreateOptions{ })
	panicFilter(error)

	log.Printf("created deployment %s \n", createdDeployment.ObjectMeta.Name)
}

func panicFilter(providedError error) {

	if providedError != nil {
		panic(providedError.Error( ))}
}