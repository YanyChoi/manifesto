package v1

import (
	v1schema "github.com/YanyChoi/manifesto/pkg/schema/v1"
	"github.com/YanyChoi/manifesto/pkg/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getDeployment(s *v1schema.V1Schema) *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: s.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: s.Replicas,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  s.ContainerName,
							Image: s.Image,
						},
					},
					TopologySpreadConstraints: []corev1.TopologySpreadConstraint{
						{
							MaxSkew:           1,
							TopologyKey:       "node.kubernetes.io/hostname",
							MinDomains:        util.GetPointerOfInt32(1),
							WhenUnsatisfiable: "DoNotSchedule",
						},
						{
							MaxSkew:           1,
							TopologyKey:       "topology.kubernetes.io/zone",
							MinDomains:        util.GetPointerOfInt32(3),
							WhenUnsatisfiable: "DoNotSchedule",
						},
					},
				},
			},
		},
	}
}
