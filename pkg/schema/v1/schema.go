package v1

import (
	yaml "gopkg.in/yaml.v3"
	log "log"
	os "os"
)

type V1Schema struct {
	Name          string `yaml:"name"`
	Replicas      *int32 `yaml:"replicas"`
	Image         string `yaml:"image"`
	ContainerName string `yaml:"containerName"`
}

func NewV1Schema(path string) *V1Schema {
	schema := &V1Schema{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	yaml.NewDecoder(file).Decode(&schema)
	return schema
}

func (s *V1Schema) Validate(obj any) error {
	return nil
}
