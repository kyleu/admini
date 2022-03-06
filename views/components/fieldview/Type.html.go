// Code generated by qtc from "Type.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/fieldview/Type.html:2
package fieldview

//line views/components/fieldview/Type.html:2
import (
	"admini.dev/app/lib/types"
)

//line views/components/fieldview/Type.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldview/Type.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldview/Type.html:6
func StreamType(qw422016 *qt422016.Writer, v types.Type) {
//line views/components/fieldview/Type.html:7
	qw422016.E().S(v.String())
//line views/components/fieldview/Type.html:8
}

//line views/components/fieldview/Type.html:8
func WriteType(qq422016 qtio422016.Writer, v types.Type) {
//line views/components/fieldview/Type.html:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/Type.html:8
	StreamType(qw422016, v)
//line views/components/fieldview/Type.html:8
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/Type.html:8
}

//line views/components/fieldview/Type.html:8
func Type(v types.Type) string {
//line views/components/fieldview/Type.html:8
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/Type.html:8
	WriteType(qb422016, v)
//line views/components/fieldview/Type.html:8
	qs422016 := string(qb422016.B)
//line views/components/fieldview/Type.html:8
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/Type.html:8
	return qs422016
//line views/components/fieldview/Type.html:8
}
