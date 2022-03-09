// Code generated by qtc from "Display.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Display.html:2
package components

//line views/components/Display.html:2
import (
	"time"

	"github.com/google/uuid"

	"admini.dev/admini/app/util"
)

//line views/components/Display.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Display.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Display.html:10
func StreamDisplayTimestamp(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:11
	qw422016.E().S(util.TimeToFull(value))
//line views/components/Display.html:12
}

//line views/components/Display.html:12
func WriteDisplayTimestamp(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:12
	StreamDisplayTimestamp(qw422016, value)
//line views/components/Display.html:12
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:12
}

//line views/components/Display.html:12
func DisplayTimestamp(value *time.Time) string {
//line views/components/Display.html:12
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:12
	WriteDisplayTimestamp(qb422016, value)
//line views/components/Display.html:12
	qs422016 := string(qb422016.B)
//line views/components/Display.html:12
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:12
	return qs422016
//line views/components/Display.html:12
}

//line views/components/Display.html:14
func StreamDisplayUUID(qw422016 *qt422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:15
	if value != nil {
//line views/components/Display.html:16
		qw422016.E().S(value.String())
//line views/components/Display.html:17
	}
//line views/components/Display.html:18
}

//line views/components/Display.html:18
func WriteDisplayUUID(qq422016 qtio422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:18
	StreamDisplayUUID(qw422016, value)
//line views/components/Display.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:18
}

//line views/components/Display.html:18
func DisplayUUID(value *uuid.UUID) string {
//line views/components/Display.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:18
	WriteDisplayUUID(qb422016, value)
//line views/components/Display.html:18
	qs422016 := string(qb422016.B)
//line views/components/Display.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:18
	return qs422016
//line views/components/Display.html:18
}

//line views/components/Display.html:20
func StreamDisplayDiffs(qw422016 *qt422016.Writer, value util.Diffs) {
//line views/components/Display.html:21
	if len(value) == 0 {
//line views/components/Display.html:21
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/Display.html:23
	} else {
//line views/components/Display.html:23
		qw422016.N().S(`<ul>`)
//line views/components/Display.html:25
		for _, d := range value {
//line views/components/Display.html:25
			qw422016.N().S(`<li><code>`)
//line views/components/Display.html:27
			qw422016.E().S(d.Path)
//line views/components/Display.html:27
			qw422016.N().S(`</code><ul><li>Old: <code>`)
//line views/components/Display.html:29
			qw422016.E().S(d.Old)
//line views/components/Display.html:29
			qw422016.N().S(`</code></li><li>New: <code>`)
//line views/components/Display.html:30
			qw422016.E().S(d.New)
//line views/components/Display.html:30
			qw422016.N().S(`</code></li></ul>`)
//line views/components/Display.html:32
		}
//line views/components/Display.html:32
		qw422016.N().S(`</ul>`)
//line views/components/Display.html:34
	}
//line views/components/Display.html:35
}

//line views/components/Display.html:35
func WriteDisplayDiffs(qq422016 qtio422016.Writer, value util.Diffs) {
//line views/components/Display.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:35
	StreamDisplayDiffs(qw422016, value)
//line views/components/Display.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:35
}

//line views/components/Display.html:35
func DisplayDiffs(value util.Diffs) string {
//line views/components/Display.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:35
	WriteDisplayDiffs(qb422016, value)
//line views/components/Display.html:35
	qs422016 := string(qb422016.B)
//line views/components/Display.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:35
	return qs422016
//line views/components/Display.html:35
}
