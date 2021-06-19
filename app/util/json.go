package util

import (
	"bytes"
	"io"

	jsoniter "github.com/json-iterator/go"
)

func ToJSON(x interface{}) string {
	return string(ToJSONBytes(x, true))
}

func ToJSONCompact(x interface{}) string {
	return string(ToJSONBytes(x, false))
}

func ToJSONBytes(x interface{}, indent bool) []byte {
	if indent {
		b, _ := jsoniter.MarshalIndent(x, "", "  ")
		return b
	}
	b, _ := jsoniter.Marshal(x)
	return b
}

func FromJSON(msg jsoniter.RawMessage, tgt interface{}) error {
	return jsoniter.Unmarshal(msg, tgt)
}

func FromJSONReader(r io.Reader, tgt interface{}) error {
	return jsoniter.NewDecoder(r).Decode(tgt)
}

func FromJSONStrict(msg jsoniter.RawMessage, tgt interface{}) error {
	dec := jsoniter.NewDecoder(bytes.NewReader(msg))
	dec.DisallowUnknownFields()
	return dec.Decode(tgt)
}
