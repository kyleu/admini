package cutil

import (
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"net/http"
)

func ParseForm(req *http.Request) (util.ValueMap, error) {
	ct := GetContentType(req)
	if IsContentTypeJSON(ct) {
		return parseJSONForm(req)
	}
	return parseHTTPForm(req)
}

func parseJSONForm(req *http.Request) (util.ValueMap, error) {
	m := map[string]interface{}{}
	err := util.FromJSONReader(req.Body, &m)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse JSON body")
	}
	ret := make(util.ValueMap, 0, len(m))
	for k, v := range m {
		ret = append(ret, &util.ValueMapEntry{Key: k, Value: v})
	}
	ret.Sort()
	return ret, nil
}

func parseHTTPForm(req *http.Request) (util.ValueMap, error) {
	if err := req.ParseForm(); err != nil {
		return nil, errors.Wrap(err, "can't parse form")
	}

	ret := make(util.ValueMap, 0, len(req.Form))
	for k, v := range req.Form {
		ret = append(ret, &util.ValueMapEntry{Key: k, Value: v})
	}
	ret.Sort()
	return ret, nil
}
