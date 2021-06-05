package cutil

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/util"
)

const (
	mimeJSON = "application/json"
	mimeXML  = "text/xml"
)

func WriteCORS(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Method", "GET,POST,DELETE,PUT,PATCH,OPTIONS,HEAD")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
}

func RespondJSON(ctx *fasthttp.RequestCtx, filename string, body interface{}) (string, error) {
	b := util.ToJSONBytes(body, true)
	return RespondMIME(filename, mimeJSON, "json", b, ctx)
}

type XMLResponse struct {
	Result interface{} `xml:"result"`
}

func RespondXML(ctx *fasthttp.RequestCtx, filename string, body interface{}) (string, error) {
	body = XMLResponse{Result: body}
	b, err := xml.Marshal(body)
	if err != nil {
		return "", errors.Wrapf(err, "can't serialize response of type [%T] to XML", body)
	}
	return RespondMIME(filename, mimeXML, "xml", b, ctx)
}

func RespondMIME(filename string, mime string, ext string, ba []byte, ctx *fasthttp.RequestCtx) (string, error) {
	ctx.Response.Header.SetContentType(mime + "; charset=UTF-8")
	if len(filename) > 0 {
		if !strings.HasSuffix(filename, "."+ext) {
			filename = filename + "." + ext
		}
		ctx.Response.Header.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	}
	WriteCORS(ctx)
	if len(ba) == 0 {
		return "", errors.New("no bytes available to write")
	}
	if _, err := ctx.Write(ba); err != nil {
		return "", errors.Wrap(err, "cannot write to response")
	}

	return "", nil
}

func GetContentType(ctx *fasthttp.RequestCtx) string {
	ret := string(ctx.Request.Header.ContentType())
	if idx := strings.Index(ret, ";"); idx > -1 {
		ret = ret[0:idx]
	}
	t := string(ctx.URI().QueryArgs().Peek("t"))
	switch t {
	case "json":
		return mimeJSON
	case "xml":
		return mimeXML
	default:
		return strings.TrimSpace(ret)
	}
}

func IsContentTypeJSON(c string) bool {
	return c == "application/json" || c == "text/json"
}

func IsContentTypeXML(c string) bool {
	return c == "application/xml" || c == mimeXML
}
