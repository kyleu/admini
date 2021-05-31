package cutil

import (
	"net/http"

	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func ParseForm(req *http.Request) (util.ValueMap, error) {
	if ct := GetContentType(req); IsContentTypeJSON(ct) {
		return parseJSONForm(req)
	}
	return parseHTTPForm(req)
}

func parseJSONForm(req *http.Request) (util.ValueMap, error) {
	ret := util.ValueMap{}
	err := util.FromJSONReader(req.Body, &ret)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse JSON body")
	}
	return ret, nil
}

func parseHTTPForm(req *http.Request) (util.ValueMap, error) {
	if err := req.ParseForm(); err != nil {
		return nil, errors.Wrap(err, "can't parse form")
	}

	ret := make(util.ValueMap, len(req.Form))
	for k, v := range req.Form {
		ret[k] = v
	}
	return ret, nil
}
