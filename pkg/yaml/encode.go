package yaml

import (
	"bytes"
	"io"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

type Encoder struct {
	encoder *json.Serializer
}

func NewEncoder() *Encoder {
	return &Encoder{json.NewYAMLSerializer(json.DefaultMetaFactory, nil, nil)}
}

func (e *Encoder) Encode(objects []runtime.Object, w io.Writer) error {
	for i, object := range objects {
		if i > 0 {
			// Add YAML document separator before each document except the first one
			if _, err := w.Write([]byte("---\n")); err != nil {
				return err
			}
		}

		// Use a buffer to capture the YAML output
		var buf bytes.Buffer
		if err := e.encoder.Encode(object, &buf); err != nil {
			return err
		}

		// Write the processed YAML to the output
		if _, err := w.Write(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
