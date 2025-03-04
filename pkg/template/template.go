package template

import (
	appsv1 "k8s.io/api/apps/v1"
)

type Template interface {
	Hydrate(deployment *appsv1.Deployment)
}
