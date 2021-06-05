package cutil

import (
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/util"
)

func ParamSetFromRequest(ctx *fasthttp.RequestCtx) util.ParamSet {
	ret := util.ParamSet{}
	args := ctx.URI().QueryArgs()
	args.VisitAll(func(key []byte, value []byte) {
		qk := string(key)
		if strings.Contains(qk, ".") {
			ret = apply(ret, qk, string(args.Peek(qk)))
		}
	})
	return ret
}

func apply(ps util.ParamSet, qk string, qv string) util.ParamSet {
	switch {
	case strings.HasSuffix(qk, ".o"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".o"))
		asc := true
		if strings.HasSuffix(qv, ".d") {
			asc = false
			qv = qv[0 : len(qv)-2]
		}
		curr.Orderings = append(curr.Orderings, &util.Ordering{Column: qv, Asc: asc})
	case strings.HasSuffix(qk, ".l"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".l"))
		li, err := strconv.ParseInt(qv, 10, 64)
		if err == nil {
			curr.Limit = int(li)
			max := 10000
			if curr.Limit > max {
				curr.Limit = max
			}
		}
	case strings.HasSuffix(qk, ".x"):
		curr := getCurr(ps, strings.TrimSuffix(qk, ".x"))
		xi, err := strconv.ParseInt(qv, 10, 64)
		if err == nil {
			curr.Offset = int(xi)
		}
	}
	return ps
}

func getCurr(q util.ParamSet, key string) *util.Params {
	curr, ok := q[key]
	if !ok {
		curr = &util.Params{Key: key}
		q[key] = curr
	}
	return curr
}
