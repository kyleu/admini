package util

import (
	"strings"
)

type Pkg []string

func (p Pkg) StringWith(extra ...string) string {
	return strings.Join(append(p, extra...), "::")
}

func (p Pkg) String() string {
	return p.StringWith()
}

func (p Pkg) ToPath(extra ...string) string {
	return strings.Join(append(p, extra...), "/")
}

func (p Pkg) Equals(other Pkg) bool {
	if len(p) != len(other) {
		return false
	}
	for idx, v := range p {
		if v != other[idx] {
			return false
		}
	}
	return true
}

func (p Pkg) Trim(src Pkg) Pkg {
	ret := make(Pkg, 0, len(p))
	for idx, v := range p {
		if len(src) >= idx && src[idx] == v {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}

func (p Pkg) Last() string {
	return p[len(p)-1]
}

func (p Pkg) Shift() Pkg {
	ret := make(Pkg, 0, len(p)-1)
	for i, s := range p {
		if i == len(p)-1 {
			break
		}
		ret = append(ret, s)
	}
	return ret
}

func (p Pkg) Push(name string) Pkg {
	if strings.Contains(name, "/") {
		return Pkg{"ERROR:contains-slash"}
	}
	return append(p, name)
}

func SplitPackage(s string) (Pkg, string) {
	sp := strings.Split(s, "::")
	pkg := sp[0 : len(sp)-1]
	n := sp[len(sp)-1]
	return pkg, n
}

func SplitPackageSlash(s string) (Pkg, string) {
	sp := strings.Split(s, "/")
	pkg := sp[0 : len(sp)-1]
	n := sp[len(sp)-1]
	return pkg, n
}
