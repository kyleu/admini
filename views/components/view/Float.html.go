// Code generated by qtc from "Float.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/Float.html:1
package view

//line views/components/view/Float.html:1
import "admini.dev/admini/app/util"

//line views/components/view/Float.html:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Float.html:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Float.html:3
func StreamFloat(qw422016 *qt422016.Writer, f any) {
//line views/components/view/Float.html:4
	qw422016.E().V(f)
//line views/components/view/Float.html:5
}

//line views/components/view/Float.html:5
func WriteFloat(qq422016 qtio422016.Writer, f any) {
//line views/components/view/Float.html:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Float.html:5
	StreamFloat(qw422016, f)
//line views/components/view/Float.html:5
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Float.html:5
}

//line views/components/view/Float.html:5
func Float(f any) string {
//line views/components/view/Float.html:5
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Float.html:5
	WriteFloat(qb422016, f)
//line views/components/view/Float.html:5
	qs422016 := string(qb422016.B)
//line views/components/view/Float.html:5
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Float.html:5
	return qs422016
//line views/components/view/Float.html:5
}

//line views/components/view/Float.html:7
func StreamFloatArray(qw422016 *qt422016.Writer, value []any) {
//line views/components/view/Float.html:8
	StreamStringArray(qw422016, util.ArrayToStringArray(value))
//line views/components/view/Float.html:9
}

//line views/components/view/Float.html:9
func WriteFloatArray(qq422016 qtio422016.Writer, value []any) {
//line views/components/view/Float.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Float.html:9
	StreamFloatArray(qw422016, value)
//line views/components/view/Float.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Float.html:9
}

//line views/components/view/Float.html:9
func FloatArray(value []any) string {
//line views/components/view/Float.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Float.html:9
	WriteFloatArray(qb422016, value)
//line views/components/view/Float.html:9
	qs422016 := string(qb422016.B)
//line views/components/view/Float.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Float.html:9
	return qs422016
//line views/components/view/Float.html:9
}
