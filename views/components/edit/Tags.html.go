// Code generated by qtc from "Tags.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/Tags.html:2
package edit

//line views/components/edit/Tags.html:2
import (
	"strings"

	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/components"
)

//line views/components/edit/Tags.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Tags.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Tags.html:9
func StreamTags(qw422016 *qt422016.Writer, key string, id string, values []string, ps *cutil.PageState, placeholder ...string) {
//line views/components/edit/Tags.html:11
	ps.AddIcon("times")
	ps.AddIcon("plus")

//line views/components/edit/Tags.html:13
	qw422016.N().S(`<div class="tag-editor">`)
//line views/components/edit/Tags.html:15
	if id == "" {
//line views/components/edit/Tags.html:15
		qw422016.N().S(`<input class="result" name="`)
//line views/components/edit/Tags.html:16
		qw422016.E().S(key)
//line views/components/edit/Tags.html:16
		qw422016.N().S(`" value="`)
//line views/components/edit/Tags.html:16
		qw422016.E().S(strings.Join(values, `, `))
//line views/components/edit/Tags.html:16
		qw422016.N().S(`"`)
//line views/components/edit/Tags.html:16
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Tags.html:16
		qw422016.N().S(`/>`)
//line views/components/edit/Tags.html:17
	} else {
//line views/components/edit/Tags.html:17
		qw422016.N().S(`<input class="result" id="`)
//line views/components/edit/Tags.html:18
		qw422016.E().S(id)
//line views/components/edit/Tags.html:18
		qw422016.N().S(`" name="`)
//line views/components/edit/Tags.html:18
		qw422016.E().S(key)
//line views/components/edit/Tags.html:18
		qw422016.N().S(`" value="`)
//line views/components/edit/Tags.html:18
		qw422016.E().S(strings.Join(values, `, `))
//line views/components/edit/Tags.html:18
		qw422016.N().S(`"`)
//line views/components/edit/Tags.html:18
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Tags.html:18
		qw422016.N().S(`/>`)
//line views/components/edit/Tags.html:19
	}
//line views/components/edit/Tags.html:19
	qw422016.N().S(`<div class="tags"></div><div class="clear"></div></div>`)
//line views/components/edit/Tags.html:23
}

//line views/components/edit/Tags.html:23
func WriteTags(qq422016 qtio422016.Writer, key string, id string, values []string, ps *cutil.PageState, placeholder ...string) {
//line views/components/edit/Tags.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Tags.html:23
	StreamTags(qw422016, key, id, values, ps, placeholder...)
//line views/components/edit/Tags.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Tags.html:23
}

//line views/components/edit/Tags.html:23
func Tags(key string, id string, values []string, ps *cutil.PageState, placeholder ...string) string {
//line views/components/edit/Tags.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Tags.html:23
	WriteTags(qb422016, key, id, values, ps, placeholder...)
//line views/components/edit/Tags.html:23
	qs422016 := string(qb422016.B)
//line views/components/edit/Tags.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Tags.html:23
	return qs422016
//line views/components/edit/Tags.html:23
}

//line views/components/edit/Tags.html:25
func StreamTagsVertical(qw422016 *qt422016.Writer, key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/edit/Tags.html:26
	id = cutil.CleanID(key, id)

//line views/components/edit/Tags.html:26
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Tags.html:28
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Tags.html:28
	qw422016.N().S(`<label for="`)
//line views/components/edit/Tags.html:29
	qw422016.E().S(id)
//line views/components/edit/Tags.html:29
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/Tags.html:29
	qw422016.E().S(title)
//line views/components/edit/Tags.html:29
	qw422016.N().S(`</em></label>`)
//line views/components/edit/Tags.html:30
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Tags.html:30
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Tags.html:31
	StreamTags(qw422016, key, id, values, ps, help...)
//line views/components/edit/Tags.html:31
	qw422016.N().S(`</div>`)
//line views/components/edit/Tags.html:32
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Tags.html:32
	qw422016.N().S(`</div>`)
//line views/components/edit/Tags.html:34
}

//line views/components/edit/Tags.html:34
func WriteTagsVertical(qq422016 qtio422016.Writer, key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/edit/Tags.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Tags.html:34
	StreamTagsVertical(qw422016, key, id, title, values, ps, indent, help...)
//line views/components/edit/Tags.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Tags.html:34
}

//line views/components/edit/Tags.html:34
func TagsVertical(key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/edit/Tags.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Tags.html:34
	WriteTagsVertical(qb422016, key, id, title, values, ps, indent, help...)
//line views/components/edit/Tags.html:34
	qs422016 := string(qb422016.B)
//line views/components/edit/Tags.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Tags.html:34
	return qs422016
//line views/components/edit/Tags.html:34
}

//line views/components/edit/Tags.html:36
func StreamTagsTable(qw422016 *qt422016.Writer, key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/edit/Tags.html:37
	id = cutil.CleanID(key, id)

//line views/components/edit/Tags.html:37
	qw422016.N().S(`<tr>`)
//line views/components/edit/Tags.html:39
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Tags.html:39
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/Tags.html:40
	qw422016.E().S(id)
//line views/components/edit/Tags.html:40
	qw422016.N().S(`"`)
//line views/components/edit/Tags.html:40
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Tags.html:40
	qw422016.N().S(`>`)
//line views/components/edit/Tags.html:40
	qw422016.E().S(title)
//line views/components/edit/Tags.html:40
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Tags.html:41
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Tags.html:41
	qw422016.N().S(`<td>`)
//line views/components/edit/Tags.html:42
	StreamTags(qw422016, key, id, values, ps, help...)
//line views/components/edit/Tags.html:42
	qw422016.N().S(`</td>`)
//line views/components/edit/Tags.html:43
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Tags.html:43
	qw422016.N().S(`</tr>`)
//line views/components/edit/Tags.html:45
}

//line views/components/edit/Tags.html:45
func WriteTagsTable(qq422016 qtio422016.Writer, key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/edit/Tags.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Tags.html:45
	StreamTagsTable(qw422016, key, id, title, values, ps, indent, help...)
//line views/components/edit/Tags.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Tags.html:45
}

//line views/components/edit/Tags.html:45
func TagsTable(key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/edit/Tags.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Tags.html:45
	WriteTagsTable(qb422016, key, id, title, values, ps, indent, help...)
//line views/components/edit/Tags.html:45
	qs422016 := string(qb422016.B)
//line views/components/edit/Tags.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Tags.html:45
	return qs422016
//line views/components/edit/Tags.html:45
}
