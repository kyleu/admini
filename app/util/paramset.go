package util

import (
	"strings"

	"go.uber.org/zap"
)

// A map of string keys to Params
type ParamSet map[string]*Params

// Gets the Params matching the provided key
func (s ParamSet) Get(key string, allowed []string, logger *zap.SugaredLogger) *Params {
	x, ok := s[key]
	if !ok {
		return &Params{Key: key}
	}

	return x.Filtered(allowed, logger)
}

func (s ParamSet) String() string {
	ret := make([]string, 0, len(s))
	for _, p := range s {
		ret = append(ret, p.String())
	}

	return strings.Join(ret, ", ")
}
