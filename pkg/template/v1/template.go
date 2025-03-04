package v1

import (
	bytes "bytes"
	v1schema "github.com/YanyChoi/manifesto/pkg/schema/v1"
	yaml "github.com/YanyChoi/manifesto/pkg/yaml"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

type Template struct {
	schema *v1schema.V1Schema
}

func NewTemplate(schema *v1schema.V1Schema) *Template {
	return &Template{schema: schema}
}

func (t *Template) Hydrate() ([]byte, error) {
	objects := writeObjects(t.schema)

	var b bytes.Buffer
	encoder := yaml.NewEncoder()
	if err := encoder.Encode(objects, &b); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func writeObjects(schema *v1schema.V1Schema) []runtime.Object {
	deployment := getDeployment(schema)
	service := getService(schema)
	return []runtime.Object{deployment, service}
}
