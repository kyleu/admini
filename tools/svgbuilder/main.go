package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

var re = regexp.MustCompile(`\n[ \\t]*`)

type SVG struct {
	Key    string
	Markup string
}

func main() {
	src := os.Args[1]
	if len(os.Args) != 3 {
		panic("pass two arguments, source directory and target directory")
	}

	svgs, err := loadSVGs(src)
	if err != nil {
		panic(err)
	}

	out := template(src, svgs)

	err = writeFile(os.Args[2], out)
	if err != nil {
		panic(err)
	}
}

func markup(key string, bytes []byte) string {
	orig := strings.TrimSpace(string(bytes))
	if !strings.Contains(orig, "id=\"svg-") {
		panic(fmt.Sprintf("no id for SVG [%s]", key))
	}
	replaced := re.ReplaceAllLiteralString(orig, "")
	return replaced
}

func loadSVGs(src string) ([]*SVG, error) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot list path [%s]", src)
	}
	var svgs []*SVG
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".svg") {
			b, err := ioutil.ReadFile(path.Join(src, f.Name()))
			if err != nil {
				panic(err)
			}
			key := strings.TrimSuffix(f.Name(), ".svg")
			svgs = append(svgs, &SVG{
				Key:    key,
				Markup: markup(key, b),
			})
		}
	}

	sort.Slice(svgs, func(i int, j int) bool {
		return svgs[i].Key < svgs[j].Key
	})

	return svgs, nil
}

func template(src string, svgs []*SVG) string {
	out := strings.Builder{}
	w := func(s string) {
		out.WriteString(s)
		out.WriteString("\n")
	}

	maxKeyLength := 0
	var keys []string
	for _, svg := range svgs {
		if len(svg.Key) > maxKeyLength {
			maxKeyLength = len(svg.Key)
		}
		switch svg.Key {
		case "search", "up", "down", "left", "right":
			// noop
		default:
			keys = append(keys, fmt.Sprintf(`"%s"`, svg.Key))
		}
	}

	w("// Code generated from files in [" + src + "]. DO NOT EDIT.")
	w("package util")
	w("")
	w("var SVGLibrary = map[string]string{")
	msg := "\t%-" + fmt.Sprintf("%d", maxKeyLength+3) + "s `%s`,"
	for _, fn := range svgs {
		w(fmt.Sprintf(msg, `"`+fn.Key+`":`, fn.Markup))
	}
	w("}")
	w("")
	w("var SVGIconKeys = []string{" + strings.Join(keys, ", ") + "}")

	return out.String()
}

func writeFile(fn string, out string) error {
	info, err := os.Stat(fn)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fn, []byte(out), info.Mode())
}
