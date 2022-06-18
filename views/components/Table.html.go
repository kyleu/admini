// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Table.html:2
package components

//line views/components/Table.html:2
import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/vutil"
)

//line views/components/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Table.html:13
func StreamTableInput(qw422016 *qt422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:13
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:15
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:15
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:16
	qw422016.E().S(key)
//line views/components/Table.html:16
	qw422016.N().S(`"`)
//line views/components/Table.html:16
	streamtitleFor(qw422016, help)
//line views/components/Table.html:16
	qw422016.N().S(`>`)
//line views/components/Table.html:16
	qw422016.E().S(title)
//line views/components/Table.html:16
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:17
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:17
	qw422016.N().S(`<td>`)
//line views/components/Table.html:18
	StreamFormInput(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:18
	qw422016.N().S(`</td>`)
//line views/components/Table.html:19
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:19
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:21
}

//line views/components/Table.html:21
func WriteTableInput(qq422016 qtio422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:21
	StreamTableInput(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:21
}

//line views/components/Table.html:21
func TableInput(key string, title string, value string, indent int, help ...string) string {
//line views/components/Table.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:21
	WriteTableInput(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:21
	qs422016 := string(qb422016.B)
//line views/components/Table.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:21
	return qs422016
//line views/components/Table.html:21
}

//line views/components/Table.html:23
func StreamTableInputPassword(qw422016 *qt422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:23
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:25
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:25
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:26
	qw422016.E().S(key)
//line views/components/Table.html:26
	qw422016.N().S(`"`)
//line views/components/Table.html:26
	streamtitleFor(qw422016, help)
//line views/components/Table.html:26
	qw422016.N().S(`>`)
//line views/components/Table.html:26
	qw422016.E().S(title)
//line views/components/Table.html:26
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:27
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:27
	qw422016.N().S(`<td>`)
//line views/components/Table.html:28
	StreamFormInputPassword(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:28
	qw422016.N().S(`</td>`)
//line views/components/Table.html:29
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:29
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:31
}

//line views/components/Table.html:31
func WriteTableInputPassword(qq422016 qtio422016.Writer, key string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:31
	StreamTableInputPassword(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:31
}

//line views/components/Table.html:31
func TableInputPassword(key string, title string, value string, indent int, help ...string) string {
//line views/components/Table.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:31
	WriteTableInputPassword(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:31
	qs422016 := string(qb422016.B)
//line views/components/Table.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:31
	return qs422016
//line views/components/Table.html:31
}

//line views/components/Table.html:33
func StreamTableInputNumber(qw422016 *qt422016.Writer, key string, title string, value int, indent int, help ...string) {
//line views/components/Table.html:33
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:35
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:35
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:36
	qw422016.E().S(key)
//line views/components/Table.html:36
	qw422016.N().S(`"`)
//line views/components/Table.html:36
	streamtitleFor(qw422016, help)
//line views/components/Table.html:36
	qw422016.N().S(`>`)
//line views/components/Table.html:36
	qw422016.E().S(title)
//line views/components/Table.html:36
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:37
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:37
	qw422016.N().S(`<td>`)
//line views/components/Table.html:38
	StreamFormInputNumber(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:38
	qw422016.N().S(`</td>`)
//line views/components/Table.html:39
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:39
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:41
}

//line views/components/Table.html:41
func WriteTableInputNumber(qq422016 qtio422016.Writer, key string, title string, value int, indent int, help ...string) {
//line views/components/Table.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:41
	StreamTableInputNumber(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:41
}

//line views/components/Table.html:41
func TableInputNumber(key string, title string, value int, indent int, help ...string) string {
//line views/components/Table.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:41
	WriteTableInputNumber(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:41
	qs422016 := string(qb422016.B)
//line views/components/Table.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:41
	return qs422016
//line views/components/Table.html:41
}

//line views/components/Table.html:43
func StreamTableInputTimestamp(qw422016 *qt422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:43
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:45
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:45
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:46
	qw422016.E().S(key)
//line views/components/Table.html:46
	qw422016.N().S(`"`)
//line views/components/Table.html:46
	streamtitleFor(qw422016, help)
//line views/components/Table.html:46
	qw422016.N().S(`>`)
//line views/components/Table.html:46
	qw422016.E().S(title)
//line views/components/Table.html:46
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:47
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:47
	qw422016.N().S(`<td>`)
//line views/components/Table.html:48
	StreamFormInputTimestamp(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:48
	qw422016.N().S(`</td>`)
//line views/components/Table.html:49
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:49
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:51
}

//line views/components/Table.html:51
func WriteTableInputTimestamp(qq422016 qtio422016.Writer, key string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:51
	StreamTableInputTimestamp(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:51
}

//line views/components/Table.html:51
func TableInputTimestamp(key string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/Table.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:51
	WriteTableInputTimestamp(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:51
	qs422016 := string(qb422016.B)
//line views/components/Table.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:51
	return qs422016
//line views/components/Table.html:51
}

//line views/components/Table.html:53
func StreamTableInputUUID(qw422016 *qt422016.Writer, key string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/Table.html:53
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:55
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:55
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:56
	qw422016.E().S(key)
//line views/components/Table.html:56
	qw422016.N().S(`"`)
//line views/components/Table.html:56
	streamtitleFor(qw422016, help)
//line views/components/Table.html:56
	qw422016.N().S(`>`)
//line views/components/Table.html:56
	qw422016.E().S(title)
//line views/components/Table.html:56
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:57
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:57
	qw422016.N().S(`<td>`)
//line views/components/Table.html:58
	StreamFormInputUUID(qw422016, key, "input-"+key, value, help...)
//line views/components/Table.html:58
	qw422016.N().S(`</td>`)
//line views/components/Table.html:59
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:59
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:61
}

//line views/components/Table.html:61
func WriteTableInputUUID(qq422016 qtio422016.Writer, key string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/Table.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:61
	StreamTableInputUUID(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:61
}

//line views/components/Table.html:61
func TableInputUUID(key string, title string, value *uuid.UUID, indent int, help ...string) string {
//line views/components/Table.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:61
	WriteTableInputUUID(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:61
	qs422016 := string(qb422016.B)
//line views/components/Table.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:61
	return qs422016
//line views/components/Table.html:61
}

//line views/components/Table.html:63
func StreamTableTextarea(qw422016 *qt422016.Writer, key string, title string, rows int, value string, indent int, help ...string) {
//line views/components/Table.html:63
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:65
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:65
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:66
	qw422016.E().S(key)
//line views/components/Table.html:66
	qw422016.N().S(`"`)
//line views/components/Table.html:66
	streamtitleFor(qw422016, help)
//line views/components/Table.html:66
	qw422016.N().S(`>`)
//line views/components/Table.html:66
	qw422016.E().S(title)
//line views/components/Table.html:66
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:67
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:67
	qw422016.N().S(`<td>`)
//line views/components/Table.html:68
	StreamFormTextarea(qw422016, key, "input-"+key, rows, value, help...)
//line views/components/Table.html:68
	qw422016.N().S(`</td>`)
//line views/components/Table.html:69
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:69
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:71
}

//line views/components/Table.html:71
func WriteTableTextarea(qq422016 qtio422016.Writer, key string, title string, rows int, value string, indent int, help ...string) {
//line views/components/Table.html:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:71
	StreamTableTextarea(qw422016, key, title, rows, value, indent, help...)
//line views/components/Table.html:71
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:71
}

//line views/components/Table.html:71
func TableTextarea(key string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/Table.html:71
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:71
	WriteTableTextarea(qb422016, key, title, rows, value, indent, help...)
//line views/components/Table.html:71
	qs422016 := string(qb422016.B)
//line views/components/Table.html:71
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:71
	return qs422016
//line views/components/Table.html:71
}

//line views/components/Table.html:73
func StreamTableSelect(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:73
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:75
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:75
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:76
	qw422016.E().S(key)
//line views/components/Table.html:76
	qw422016.N().S(`"`)
//line views/components/Table.html:76
	streamtitleFor(qw422016, help)
//line views/components/Table.html:76
	qw422016.N().S(`>`)
//line views/components/Table.html:76
	qw422016.E().S(title)
//line views/components/Table.html:76
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:77
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:77
	qw422016.N().S(`<td>`)
//line views/components/Table.html:78
	StreamFormSelect(qw422016, key, "input-"+key, value, opts, titles, indent)
//line views/components/Table.html:78
	qw422016.N().S(`</td>`)
//line views/components/Table.html:79
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:79
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:81
}

//line views/components/Table.html:81
func WriteTableSelect(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:81
	StreamTableSelect(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:81
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:81
}

//line views/components/Table.html:81
func TableSelect(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:81
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:81
	WriteTableSelect(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:81
	qs422016 := string(qb422016.B)
//line views/components/Table.html:81
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:81
	return qs422016
//line views/components/Table.html:81
}

//line views/components/Table.html:83
func StreamTableDatalist(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:83
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:85
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:85
	qw422016.N().S(`<th class="shrink"><label for="input-`)
//line views/components/Table.html:86
	qw422016.E().S(key)
//line views/components/Table.html:86
	qw422016.N().S(`"`)
//line views/components/Table.html:86
	streamtitleFor(qw422016, help)
//line views/components/Table.html:86
	qw422016.N().S(`>`)
//line views/components/Table.html:86
	qw422016.E().S(title)
//line views/components/Table.html:86
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:87
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:87
	qw422016.N().S(`<td>`)
//line views/components/Table.html:88
	StreamFormDatalist(qw422016, key, "input-"+key, value, opts, titles, indent)
//line views/components/Table.html:88
	qw422016.N().S(`</td>`)
//line views/components/Table.html:89
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:89
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:91
}

//line views/components/Table.html:91
func WriteTableDatalist(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:91
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:91
	StreamTableDatalist(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:91
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:91
}

//line views/components/Table.html:91
func TableDatalist(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:91
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:91
	WriteTableDatalist(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:91
	qs422016 := string(qb422016.B)
//line views/components/Table.html:91
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:91
	return qs422016
//line views/components/Table.html:91
}

//line views/components/Table.html:93
func StreamTableRadio(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:93
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:95
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:95
	qw422016.N().S(`<th class="shrink"><label`)
//line views/components/Table.html:96
	streamtitleFor(qw422016, help)
//line views/components/Table.html:96
	qw422016.N().S(`>`)
//line views/components/Table.html:96
	qw422016.E().S(title)
//line views/components/Table.html:96
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:97
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:97
	qw422016.N().S(`<td>`)
//line views/components/Table.html:99
	StreamFormRadio(qw422016, key, value, opts, titles, indent+2)
//line views/components/Table.html:100
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:100
	qw422016.N().S(`</td>`)
//line views/components/Table.html:102
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:102
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:104
}

//line views/components/Table.html:104
func WriteTableRadio(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:104
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:104
	StreamTableRadio(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:104
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:104
}

//line views/components/Table.html:104
func TableRadio(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:104
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:104
	WriteTableRadio(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:104
	qs422016 := string(qb422016.B)
//line views/components/Table.html:104
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:104
	return qs422016
//line views/components/Table.html:104
}

//line views/components/Table.html:106
func StreamTableBoolean(qw422016 *qt422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/Table.html:107
	StreamTableRadio(qw422016, key, title, fmt.Sprint(value), []string{"true", "false"}, []string{"True", "False"}, indent)
//line views/components/Table.html:108
}

//line views/components/Table.html:108
func WriteTableBoolean(qq422016 qtio422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/Table.html:108
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:108
	StreamTableBoolean(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:108
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:108
}

//line views/components/Table.html:108
func TableBoolean(key string, title string, value bool, indent int, help ...string) string {
//line views/components/Table.html:108
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:108
	WriteTableBoolean(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:108
	qs422016 := string(qb422016.B)
//line views/components/Table.html:108
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:108
	return qs422016
//line views/components/Table.html:108
}

//line views/components/Table.html:110
func StreamTableCheckbox(qw422016 *qt422016.Writer, key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) {
//line views/components/Table.html:110
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:112
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:112
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:113
	qw422016.E().S(title)
//line views/components/Table.html:113
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:114
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:114
	qw422016.N().S(`<td>`)
//line views/components/Table.html:116
	StreamFormCheckbox(qw422016, key, values, opts, titles, linebreaks, indent+2)
//line views/components/Table.html:117
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:117
	qw422016.N().S(`</td>`)
//line views/components/Table.html:119
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:119
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:121
}

//line views/components/Table.html:121
func WriteTableCheckbox(qq422016 qtio422016.Writer, key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) {
//line views/components/Table.html:121
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:121
	StreamTableCheckbox(qw422016, key, title, values, opts, titles, linebreaks, indent, help...)
//line views/components/Table.html:121
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:121
}

//line views/components/Table.html:121
func TableCheckbox(key string, title string, values []string, opts []string, titles []string, linebreaks bool, indent int, help ...string) string {
//line views/components/Table.html:121
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:121
	WriteTableCheckbox(qb422016, key, title, values, opts, titles, linebreaks, indent, help...)
//line views/components/Table.html:121
	qs422016 := string(qb422016.B)
//line views/components/Table.html:121
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:121
	return qs422016
//line views/components/Table.html:121
}

//line views/components/Table.html:123
func StreamTableIcons(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/Table.html:123
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:125
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:125
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:126
	qw422016.E().S(title)
//line views/components/Table.html:126
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:127
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:127
	qw422016.N().S(`<td>`)
//line views/components/Table.html:128
	StreamIconPicker(qw422016, value, key, ps, indent+2)
//line views/components/Table.html:128
	qw422016.N().S(`</td>`)
//line views/components/Table.html:130
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:130
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:132
}

//line views/components/Table.html:132
func WriteTableIcons(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/Table.html:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:132
	StreamTableIcons(qw422016, key, title, value, ps, indent, help...)
//line views/components/Table.html:132
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:132
}

//line views/components/Table.html:132
func TableIcons(key string, title string, value string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/Table.html:132
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:132
	WriteTableIcons(qb422016, key, title, value, ps, indent, help...)
//line views/components/Table.html:132
	qs422016 := string(qb422016.B)
//line views/components/Table.html:132
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:132
	return qs422016
//line views/components/Table.html:132
}

//line views/components/Table.html:134
func streamtitleFor(qw422016 *qt422016.Writer, help []string) {
//line views/components/Table.html:135
	if len(help) > 0 {
//line views/components/Table.html:135
		qw422016.N().S(` `)
//line views/components/Table.html:135
		qw422016.N().S(`title="`)
//line views/components/Table.html:135
		qw422016.E().S(strings.Join(help, "; "))
//line views/components/Table.html:135
		qw422016.N().S(`"`)
//line views/components/Table.html:135
	}
//line views/components/Table.html:136
}

//line views/components/Table.html:136
func writetitleFor(qq422016 qtio422016.Writer, help []string) {
//line views/components/Table.html:136
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:136
	streamtitleFor(qw422016, help)
//line views/components/Table.html:136
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:136
}

//line views/components/Table.html:136
func titleFor(help []string) string {
//line views/components/Table.html:136
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:136
	writetitleFor(qb422016, help)
//line views/components/Table.html:136
	qs422016 := string(qb422016.B)
//line views/components/Table.html:136
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:136
	return qs422016
//line views/components/Table.html:136
}
