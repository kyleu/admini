package schema

import "errors"

type Summary struct {
	Key         string    `json:"key"`
	Title       string    `json:"title"`
	Paths       Paths     `json:"paths"`
	Description string    `json:"description,omitempty"`
	Metadata    *Metadata `json:"metadata,omitempty"`
}

type Summaries []*Summary

type Schema struct {
	Key         string    `json:"key"`
	Title       string    `json:"title"`
	Paths       Paths     `json:"paths"`
	Scalars     Scalars   `json:"scalars,omitempty"`
	Models      Models    `json:"models,omitempty"`
	Errors      []string  `json:"errors,omitempty"`
	Description string    `json:"description,omitempty"`
	Metadata    *Metadata `json:"metadata,omitempty"`
}

type Schemata []*Schema

func NewSchema(key string, title string, paths []string, md *Metadata) *Schema {
	return &Schema{Key: key, Title: title, Paths: paths, Metadata: md}
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
		return errors.New("nil scalar")
	}
	if s.Scalars.Get(sc.Pkg, sc.Key) != nil {
		return errors.New(alreadyExists("scalar", sc.Key))
	}
	s.Scalars = append(s.Scalars, sc)
	return nil
}

func (s *Schema) AddModel(m *Model) error {
	if m == nil {
		return errors.New("nil model")
	}
	if s.Models.Get(m.Pkg, m.Key) != nil {
		return errors.New(alreadyExists("model", m.Key))
	}
	s.Models = append(s.Models, m)
	return nil
}

func (s *Schema) Validate() *ValidationResult {
	return validateSchema(s)
}

func (s *Schema) ValidateModel(model *Model) *ValidationResult {
	r := &ValidationResult{Schema: s.Key}
	return validateModel(r, s, model)
}
