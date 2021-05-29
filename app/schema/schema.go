package schema

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"
)

type Schema struct {
	Paths           Paths        `json:"paths,omitempty"`
	Scalars         Scalars      `json:"scalars,omitempty"`
	Models          model.Models `json:"models,omitempty"`
	Metadata        *Metadata    `json:"metadata,omitempty"`
	modelsByPackage *model.Package
}

type Schemata map[string]*Schema

func (s Schemata) Get(key string) *Schema {
	return s[key]
}

func (s Schemata) GetWithError(key string) (*Schema, error) {
	if ret := s.Get(key); ret != nil {
		return ret, nil
	}

	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return nil, errors.Errorf("no schema [%v] available among candidates [%v]", key, strings.Join(keys, ", "))
}

func NewSchema(paths []string, md *Metadata) *Schema {
	return &Schema{Paths: paths, Metadata: md}
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
		return errors.New("nil scalar")
	}
	if s.Scalars.Get(sc.Pkg, sc.Key) != nil {
		return errors.New("scalar [" + sc.Key + "] already exists")
	}
	s.Scalars = append(s.Scalars, sc)
	return nil
}

func (s *Schema) AddModel(m *model.Model) error {
	if m == nil {
		return errors.New("nil model")
	}
	if s.Models.Get(m.Pkg, m.Key) != nil {
		return errors.New("model [" + m.Key + "] already exists")
	}
	s.Models = append(s.Models, m)
	return nil
}

func (s *Schema) Validate() *ValidationResult {
	return validateSchema(s)
}

func (s *Schema) ValidateModel(m *model.Model) *ValidationResult {
	r := &ValidationResult{Schema: "TODO"}
	return validateModel(r, s, m)
}

func (s *Schema) ModelsByPackage() *model.Package {
	if s.modelsByPackage == nil {
		s.modelsByPackage = model.ToModelPackage(s.Models)
	}
	return s.modelsByPackage
}
