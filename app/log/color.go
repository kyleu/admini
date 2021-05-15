package log

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Color represents a text color.
type Color uint8

// Add adds the coloring to the given string.
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

var levelToColor = map[zapcore.Level]Color{
	zapcore.DebugLevel:  Magenta,
	zapcore.InfoLevel:   Blue,
	zapcore.WarnLevel:   Yellow,
	zapcore.ErrorLevel:  Red,
	zapcore.DPanicLevel: Red,
	zapcore.PanicLevel:  Red,
	zapcore.FatalLevel:  Red,
}
