// Code generated by qtc from "test.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/test.sql:1
package queries

//line queries/test.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/test.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/test.sql:1
func StreamTest(qw422016 *qt422016.Writer) {
//line queries/test.sql:1
	qw422016.N().S(`
select 1;
-- `)
//line queries/test.sql:3
}

//line queries/test.sql:3
func WriteTest(qq422016 qtio422016.Writer) {
//line queries/test.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/test.sql:3
	StreamTest(qw422016)
//line queries/test.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/test.sql:3
}

//line queries/test.sql:3
func Test() string {
//line queries/test.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/test.sql:3
	WriteTest(qb422016)
//line queries/test.sql:3
	qs422016 := string(qb422016.B)
//line queries/test.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/test.sql:3
	return qs422016
//line queries/test.sql:3
}
