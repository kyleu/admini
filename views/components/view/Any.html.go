// Code generated by qtc from "Any.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Any.html:2
package view

//line views/components/view/Any.html:2
import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"admini.dev/admini/app/util"
)

//line views/components/view/Any.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Any.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Any.html:12
func StreamAny(qw422016 *qt422016.Writer, x any) {
//line views/components/view/Any.html:13
	if x == nil {
//line views/components/view/Any.html:13
		qw422016.N().S(`<em>nil</em>`)
//line views/components/view/Any.html:15
	} else {
//line views/components/view/Any.html:16
		switch t := x.(type) {
//line views/components/view/Any.html:17
		case bool:
//line views/components/view/Any.html:18
			StreamBool(qw422016, t)
//line views/components/view/Any.html:19
		case util.Diffs:
//line views/components/view/Any.html:20
			StreamDiffs(qw422016, t)
//line views/components/view/Any.html:21
		case float32:
//line views/components/view/Any.html:22
			StreamFloat(qw422016, t)
//line views/components/view/Any.html:23
		case float64:
//line views/components/view/Any.html:24
			StreamFloat(qw422016, t)
//line views/components/view/Any.html:25
		case int:
//line views/components/view/Any.html:26
			StreamInt(qw422016, t)
//line views/components/view/Any.html:27
		case int32:
//line views/components/view/Any.html:28
			StreamInt(qw422016, t)
//line views/components/view/Any.html:29
		case int64:
//line views/components/view/Any.html:30
			StreamInt(qw422016, t)
//line views/components/view/Any.html:31
		case util.ValueMap:
//line views/components/view/Any.html:32
			StreamMap(qw422016, false, t)
//line views/components/view/Any.html:33
		case []util.ValueMap:
//line views/components/view/Any.html:34
			StreamMapArray(qw422016, false, t...)
//line views/components/view/Any.html:35
		case util.Pkg:
//line views/components/view/Any.html:36
			StreamPackage(qw422016, t)
//line views/components/view/Any.html:37
		case string:
//line views/components/view/Any.html:38
			StreamString(qw422016, t)
//line views/components/view/Any.html:39
		case time.Time:
//line views/components/view/Any.html:40
			StreamTimestamp(qw422016, &t)
//line views/components/view/Any.html:41
		case *time.Time:
//line views/components/view/Any.html:42
			StreamTimestamp(qw422016, t)
//line views/components/view/Any.html:43
		case url.URL:
//line views/components/view/Any.html:44
			StreamURL(qw422016, t)
//line views/components/view/Any.html:45
		case uuid.UUID:
//line views/components/view/Any.html:46
			StreamUUID(qw422016, &t)
//line views/components/view/Any.html:47
		case *uuid.UUID:
//line views/components/view/Any.html:48
			StreamUUID(qw422016, t)
//line views/components/view/Any.html:49
		default:
//line views/components/view/Any.html:49
			qw422016.N().S(`unhandled type [`)
//line views/components/view/Any.html:50
			qw422016.E().S(fmt.Sprintf("%T", x))
//line views/components/view/Any.html:50
			qw422016.N().S(`]`)
//line views/components/view/Any.html:51
		}
//line views/components/view/Any.html:52
	}
//line views/components/view/Any.html:53
}

//line views/components/view/Any.html:53
func WriteAny(qq422016 qtio422016.Writer, x any) {
//line views/components/view/Any.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Any.html:53
	StreamAny(qw422016, x)
//line views/components/view/Any.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Any.html:53
}

//line views/components/view/Any.html:53
func Any(x any) string {
//line views/components/view/Any.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Any.html:53
	WriteAny(qb422016, x)
//line views/components/view/Any.html:53
	qs422016 := string(qb422016.B)
//line views/components/view/Any.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Any.html:53
	return qs422016
//line views/components/view/Any.html:53
}
