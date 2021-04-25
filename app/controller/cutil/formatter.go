package cutil

import (
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/kyleu/admini/app/util"
)

func Format(v interface{}) (string, error) {
	s := styles.MonokaiLight
	l := lexers.Get("json")
	f := html.New(html.WithClasses(true), html.WithLineNumbers(true), html.LineNumbersInTable(true))
	j := util.ToJSON(v)
	i, err := l.Tokenise(nil, j)
	if err != nil {
		return "", fmt.Errorf("can't tokenize: %w", err)
	}
	x := &strings.Builder{}
	err = f.Format(x, s, i)
	if err != nil {
		return "", fmt.Errorf("can't format: %w", err)
	}

	return x.String(), nil
}
