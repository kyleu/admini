package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
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
		panic("no id for SVG [" + key + "]")
	}
	replaced := re.ReplaceAllLiteralString(orig, "")
	return replaced
}


func loadSVGs(src string) ([]*SVG, error) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return nil, fmt.Errorf("cannot list path [%v]: %w", src, err)
	}
	svgs := make([]*SVG, 0)
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
	var w = func(s string) {
		out.WriteString(s)
		out.WriteString("\n")
	}
	w("// Code generated from files in [" + src + "]. DO NOT EDIT.")
	w("package util")
	w("")
	w("var SVGLibrary = map[string]string{")
	for _, fn := range svgs {
		w(fmt.Sprintf("\t\"%v\": `%v`,", fn.Key, fn.Markup))
	}
	w("}")

	return out.String()
}

func writeFile(fn string, out string) error {
	info, err := os.Stat(fn)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fn, []byte(out), info.Mode())
}
