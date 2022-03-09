// Code generated by qtc from "String.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/fieldedit/String.html:2
package fieldedit

//line views/components/fieldedit/String.html:2
import (
	"fmt"

	"admini.dev/admini/views/components"
)

//line views/components/fieldedit/String.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldedit/String.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldedit/String.html:8
func StreamString(qw422016 *qt422016.Writer, x interface{}, k string) {
//line views/components/fieldedit/String.html:9
	components.StreamFormInput(qw422016, k, "", fmt.Sprintf("%v", x))
//line views/components/fieldedit/String.html:10
}

//line views/components/fieldedit/String.html:10
func WriteString(qq422016 qtio422016.Writer, x interface{}, k string) {
//line views/components/fieldedit/String.html:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldedit/String.html:10
	StreamString(qw422016, x, k)
//line views/components/fieldedit/String.html:10
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldedit/String.html:10
}

//line views/components/fieldedit/String.html:10
func String(x interface{}, k string) string {
//line views/components/fieldedit/String.html:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldedit/String.html:10
	WriteString(qb422016, x, k)
//line views/components/fieldedit/String.html:10
	qs422016 := string(qb422016.B)
//line views/components/fieldedit/String.html:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldedit/String.html:10
	return qs422016
//line views/components/fieldedit/String.html:10
}
