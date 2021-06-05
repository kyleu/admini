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
	return nil, errors.Errorf("no schema [%s] available among candidates [%s]", key, strings.Join(keys, ", "))
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
		return errors.Errorf("scalar [%s] already exists", sc.Key)
	}
	s.Scalars = append(s.Scalars, sc)
	return nil
}

func (s *Schema) AddModel(m *model.Model) error {
	if m == nil {
		return errors.New("nil model")
	}
	if s.Models.Get(m.Pkg, m.Key) != nil {
		return errors.Errorf("model [%s] already exists", m.Key)
	}
	s.Models = append(s.Models, m)
	return nil
}

func (s *Schema) Validate(n string) *ValidationResult {
	return validateSchema(n, s)
}

func (s *Schema) ValidateModel(sch string, m *model.Model) *ValidationResult {
	r := &ValidationResult{Schema: sch}
	return validateModel(r, s, m)
}

func (s *Schema) ModelsByPackage() *model.Package {
	if s.modelsByPackage == nil {
		s.modelsByPackage = model.ToModelPackage(s.Models)
	}
	return s.modelsByPackage
}
