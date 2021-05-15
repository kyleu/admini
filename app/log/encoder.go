package log

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const timeFormat = "15:04:05.000000"

type customEncoder struct {
	zapcore.Encoder
	pool buffer.Pool
}

func NewEncoder(cfg zapcore.EncoderConfig) *customEncoder {
	return &customEncoder{Encoder: zapcore.NewJSONEncoder(cfg), pool: buffer.NewPool()}
}

func (e *customEncoder) Clone() zapcore.Encoder {
	return &customEncoder{Encoder: e.Encoder.Clone(), pool: e.pool}
}

func (e *customEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	b, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, errors.Wrap(err, "logging error")
	}
	out := b.Bytes()
	b.Free()

	data := map[string]interface{}{}
	err = util.FromJSON(out, &data)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse logging JSON")
	}

	delete(data, "C")
	delete(data, "L")
	delete(data, "M")
	delete(data, "T")
	delete(data, "func")
	delete(data, "stacktrace")

	ret := e.pool.Get()
	addLine := func(l string) {
		ret.AppendString(l)
		ret.AppendByte('\n')
	}

	lvl := levelToColor[entry.Level].Add(entry.Level.CapitalString())
	tm := entry.Time.Format(timeFormat)
	addLine(fmt.Sprintf("[%v] %v - %v", lvl, tm, Cyan.Add(entry.Message)))
	if len(data) > 0 {
		addLine(util.ToJSONCompact(data))
	}
	caller := entry.Caller.String()
	if entry.Caller.Function != "" {
		caller += " (" + entry.Caller.Function + ")"
	}
	addLine(caller)

	if entry.Stack != "" {
		st := strings.Split(entry.Stack, "\n")
		for _, stl := range st {
			addLine("  " + stl)
		}
	}
	return ret, nil
}
