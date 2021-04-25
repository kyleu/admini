package schema

import (
	"fmt"
)

type Schema struct {
	Paths           Paths     `json:"paths,omitempty"`
	Scalars         Scalars   `json:"scalars,omitempty"`
	Models          Models    `json:"models,omitempty"`
	Errors          []string  `json:"errors,omitempty"`
	Metadata        *Metadata `json:"metadata,omitempty"`
	modelsByPackage *ModelPackage
}

type Schemata []*Schema

func NewSchema(paths []string, md *Metadata) *Schema {
	return &Schema{Paths: paths, Metadata: md}
}

func alreadyExists(t string, key string) string {
	return t + " [" + key + "] already exists"
}

func (s *Schema) AddPath(path string) bool {
	if path == "" {
		return false
	}
	if s.Paths.Exists(path) {
		return false
	}
	s.Paths = append(s.Paths, path)
	return true
}

func (s *Schema) AddScalar(sc *Scalar) error {
	if sc == nil {
		return fmt.Errorf("nil scalar")
	}
	if s.Scalars.Get(sc.Pkg, sc.Key) != nil {
		return fmt.Errorf(alreadyExists("scalar", sc.Key))
	}
	s.Scalars = append(s.Scalars, sc)
	return nil
}

func (s *Schema) AddModel(m *Model) error {
	if m == nil {
		return fmt.Errorf("nil model")
	}
	if s.Models.Get(m.Pkg, m.Key) != nil {
		return fmt.Errorf(alreadyExists("model", m.Key))
	}
	s.Models = append(s.Models, m)
	return nil
}

func (s *Schema) Validate() *ValidationResult {
	return validateSchema(s)
}

func (s *Schema) ValidateModel(model *Model) *ValidationResult {
	r := &ValidationResult{Schema: "TODO"}
	return validateModel(r, s, model)
}

func (s *Schema) ModelsByPackage() *ModelPackage {
	if s.modelsByPackage == nil {
		s.modelsByPackage = ToModelPackage(s.Models)
	}
	return s.modelsByPackage
}
