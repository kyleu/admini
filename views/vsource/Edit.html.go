// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsource/Edit.html:1
package vsource

//line views/vsource/Edit.html:1
import (
	"admini.dev/app"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/lib/schema"
	"admini.dev/app/loader/lmysql"
	"admini.dev/app/loader/lpostgres"
	"admini.dev/app/loader/lsqlite"
	"admini.dev/app/source"
	"admini.dev/views/components"
	"admini.dev/views/layout"
)

//line views/vsource/Edit.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsource/Edit.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsource/Edit.html:13
type Edit struct {
	layout.Basic
	Source *source.Source
}

//line views/vsource/Edit.html:18
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/Edit.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a class="link-confirm" data-message="Are you sure you want to delete source [`)
//line views/vsource/Edit.html:21
	qw422016.E().S(p.Source.Key)
//line views/vsource/Edit.html:21
	qw422016.N().S(`]?" href="/source/`)
//line views/vsource/Edit.html:21
	qw422016.E().S(p.Source.Key)
//line views/vsource/Edit.html:21
	qw422016.N().S(`/delete">Delete Source</a>
    </div>
    <h3>Edit [`)
//line views/vsource/Edit.html:23
	qw422016.E().S(p.Source.Key)
//line views/vsource/Edit.html:23
	qw422016.N().S(`]</h3>
    <form class="mt" action="/source/`)
//line views/vsource/Edit.html:24
	qw422016.E().S(p.Source.Key)
//line views/vsource/Edit.html:24
	qw422016.N().S(`" method="post" enctype="application/x-www-form-urlencoded">
      <table class="expanded">
        <tbody>
          `)
//line views/vsource/Edit.html:27
	components.StreamTableInput(qw422016, "title", "Title", p.Source.Title, 5)
//line views/vsource/Edit.html:27
	qw422016.N().S(`
          `)
//line views/vsource/Edit.html:28
	components.StreamTableIcons(qw422016, "icon", "Icon", p.Source.Icon, ps, 5)
//line views/vsource/Edit.html:28
	qw422016.N().S(`
          `)
//line views/vsource/Edit.html:29
	components.StreamTableInput(qw422016, "description", "Description", p.Source.Description, 5)
//line views/vsource/Edit.html:29
	qw422016.N().S(`
`)
//line views/vsource/Edit.html:30
	switch p.Source.Type {
//line views/vsource/Edit.html:31
	case schema.OriginMySQL:
//line views/vsource/Edit.html:31
		qw422016.N().S(`          `)
//line views/vsource/Edit.html:32
		streammySQLFields(qw422016, p.Source.Config, 5)
//line views/vsource/Edit.html:32
		qw422016.N().S(`
`)
//line views/vsource/Edit.html:33
	case schema.OriginPostgres:
//line views/vsource/Edit.html:33
		qw422016.N().S(`          `)
//line views/vsource/Edit.html:34
		streampostgresFields(qw422016, p.Source.Config, 5)
//line views/vsource/Edit.html:34
		qw422016.N().S(`
`)
//line views/vsource/Edit.html:35
	case schema.OriginSQLite:
//line views/vsource/Edit.html:35
		qw422016.N().S(`          `)
//line views/vsource/Edit.html:36
		streamsqliteFields(qw422016, p.Source.Config, 5)
//line views/vsource/Edit.html:36
		qw422016.N().S(`
`)
//line views/vsource/Edit.html:37
	}
//line views/vsource/Edit.html:37
	qw422016.N().S(`        </tbody>
      </table>
      <div class="mt">
        <button type="submit">Save Changes</button>
        <button type="reset">Reset</button>
      </div>
    </form>
  </div>
`)
//line views/vsource/Edit.html:46
}

//line views/vsource/Edit.html:46
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsource/Edit.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/Edit.html:46
	p.StreamBody(qw422016, as, ps)
//line views/vsource/Edit.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/Edit.html:46
}

//line views/vsource/Edit.html:46
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsource/Edit.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/Edit.html:46
	p.WriteBody(qb422016, as, ps)
//line views/vsource/Edit.html:46
	qs422016 := string(qb422016.B)
//line views/vsource/Edit.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/Edit.html:46
	return qs422016
//line views/vsource/Edit.html:46
}

//line views/vsource/Edit.html:48
func streampostgresFields(qw422016 *qt422016.Writer, b []byte, indent int) {
//line views/vsource/Edit.html:50
	cfg, _ := lpostgres.LoadConfig(b)

//line views/vsource/Edit.html:52
	components.StreamTableInput(qw422016, "host", "Host", cfg.Host, indent)
//line views/vsource/Edit.html:53
	components.StreamTableInputNumber(qw422016, "port", "Port", cfg.Port, indent)
//line views/vsource/Edit.html:54
	components.StreamTableInput(qw422016, "username", "Username", cfg.Username, indent)
//line views/vsource/Edit.html:55
	components.StreamTableInput(qw422016, "password", "Password", cfg.Password, indent)
//line views/vsource/Edit.html:56
	components.StreamTableInput(qw422016, "database", "Database", cfg.Database, indent)
//line views/vsource/Edit.html:57
	components.StreamTableInput(qw422016, "schema", "Schema", cfg.Schema, indent)
//line views/vsource/Edit.html:58
	components.StreamTableBoolean(qw422016, "debug", "Debug", cfg.Debug, indent)
//line views/vsource/Edit.html:59
}

//line views/vsource/Edit.html:59
func writepostgresFields(qq422016 qtio422016.Writer, b []byte, indent int) {
//line views/vsource/Edit.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/Edit.html:59
	streampostgresFields(qw422016, b, indent)
//line views/vsource/Edit.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/Edit.html:59
}

//line views/vsource/Edit.html:59
func postgresFields(b []byte, indent int) string {
//line views/vsource/Edit.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/Edit.html:59
	writepostgresFields(qb422016, b, indent)
//line views/vsource/Edit.html:59
	qs422016 := string(qb422016.B)
//line views/vsource/Edit.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/Edit.html:59
	return qs422016
//line views/vsource/Edit.html:59
}

//line views/vsource/Edit.html:61
func streammySQLFields(qw422016 *qt422016.Writer, b []byte, indent int) {
//line views/vsource/Edit.html:63
	cfg, _ := lmysql.LoadConfig(b)

//line views/vsource/Edit.html:65
	components.StreamTableInput(qw422016, "host", "Host", cfg.Host, indent)
//line views/vsource/Edit.html:66
	components.StreamTableInputNumber(qw422016, "port", "Port", cfg.Port, indent)
//line views/vsource/Edit.html:67
	components.StreamTableInput(qw422016, "username", "Username", cfg.Username, indent)
//line views/vsource/Edit.html:68
	components.StreamTableInput(qw422016, "password", "Password", cfg.Password, indent)
//line views/vsource/Edit.html:69
	components.StreamTableInput(qw422016, "database", "Database", cfg.Database, indent)
//line views/vsource/Edit.html:70
	components.StreamTableInput(qw422016, "schema", "Schema", cfg.Schema, indent)
//line views/vsource/Edit.html:71
	components.StreamTableBoolean(qw422016, "debug", "Debug", cfg.Debug, indent)
//line views/vsource/Edit.html:72
}

//line views/vsource/Edit.html:72
func writemySQLFields(qq422016 qtio422016.Writer, b []byte, indent int) {
//line views/vsource/Edit.html:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/Edit.html:72
	streammySQLFields(qw422016, b, indent)
//line views/vsource/Edit.html:72
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/Edit.html:72
}

//line views/vsource/Edit.html:72
func mySQLFields(b []byte, indent int) string {
//line views/vsource/Edit.html:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/Edit.html:72
	writemySQLFields(qb422016, b, indent)
//line views/vsource/Edit.html:72
	qs422016 := string(qb422016.B)
//line views/vsource/Edit.html:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/Edit.html:72
	return qs422016
//line views/vsource/Edit.html:72
}

//line views/vsource/Edit.html:74
func streamsqliteFields(qw422016 *qt422016.Writer, b []byte, indent int) {
//line views/vsource/Edit.html:76
	println(len(b))
	cfg, err := lsqlite.LoadConfig(b)
	if err != nil {
		panic(err)
	}

//line views/vsource/Edit.html:82
	components.StreamTableInput(qw422016, "file", "File", cfg.File, indent)
//line views/vsource/Edit.html:83
	components.StreamTableInput(qw422016, "schema", "Schema", cfg.Schema, indent)
//line views/vsource/Edit.html:84
	components.StreamTableBoolean(qw422016, "debug", "Debug", cfg.Debug, indent)
//line views/vsource/Edit.html:85
}

//line views/vsource/Edit.html:85
func writesqliteFields(qq422016 qtio422016.Writer, b []byte, indent int) {
//line views/vsource/Edit.html:85
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsource/Edit.html:85
	streamsqliteFields(qw422016, b, indent)
//line views/vsource/Edit.html:85
	qt422016.ReleaseWriter(qw422016)
//line views/vsource/Edit.html:85
}

//line views/vsource/Edit.html:85
func sqliteFields(b []byte, indent int) string {
//line views/vsource/Edit.html:85
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsource/Edit.html:85
	writesqliteFields(qb422016, b, indent)
//line views/vsource/Edit.html:85
	qs422016 := string(qb422016.B)
//line views/vsource/Edit.html:85
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsource/Edit.html:85
	return qs422016
//line views/vsource/Edit.html:85
}
