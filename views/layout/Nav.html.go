// Code generated by qtc from "Nav.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Nav.html:2
package layout

//line views/layout/Nav.html:2
import (
	"strings"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cmenu"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/menu"
	"admini.dev/admini/views/components"
	"admini.dev/admini/views/vutil"
)

//line views/layout/Nav.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Nav.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Nav.html:13
func StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:13
	qw422016.N().S(`
<nav id="navbar">
  <a class="logo" href="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:15
	qw422016.N().S(`" title="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:15
	qw422016.N().S(` `)
//line views/layout/Nav.html:15
	qw422016.E().S(as.BuildInfo.String())
//line views/layout/Nav.html:15
	qw422016.N().S(`">`)
//line views/layout/Nav.html:15
	components.StreamSVGRef(qw422016, ps.RootIcon, 32, 32, ``, ps)
//line views/layout/Nav.html:15
	qw422016.N().S(`</a>
  <div class="breadcrumbs">
    <a class="link" href="`)
//line views/layout/Nav.html:17
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:17
	qw422016.N().S(`">`)
//line views/layout/Nav.html:17
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:17
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:17
	StreamNavItems(qw422016, ps.Menu, ps.Breadcrumbs)
//line views/layout/Nav.html:17
	qw422016.N().S(`
  </div>
  <form action="`)
//line views/layout/Nav.html:19
	qw422016.E().S(ps.SearchPath)
//line views/layout/Nav.html:19
	qw422016.N().S(`" class="search" title="search">
    <input type="search" name="q" placeholder=" " />
    <div class="search-image" style="display: none;"><svg><use xlink:href="#svg-searchbox" /></svg></div>
  </form>
  <a class="profile" title="`)
//line views/layout/Nav.html:23
	qw422016.E().S(ps.Profile.AuthString(ps.Accounts))
//line views/layout/Nav.html:23
	qw422016.N().S(`" href="`)
//line views/layout/Nav.html:23
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:23
	qw422016.N().S(`">`)
//line views/layout/Nav.html:23
	components.StreamSVGRef(qw422016, `profile`, 24, 24, ``, ps)
//line views/layout/Nav.html:23
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:23
	if !ps.HideMenu {
//line views/layout/Nav.html:23
		qw422016.N().S(`

  <input type="checkbox" id="menu-toggle-input" style="display: none;" />
  <label class="menu-toggle" for="menu-toggle-input"><div class="spinner diagonal part-1"></div><div class="spinner horizontal"></div><div class="spinner diagonal part-2"></div></label>
  `)
//line views/layout/Nav.html:27
		StreamMenu(qw422016, ps)
//line views/layout/Nav.html:27
	}
//line views/layout/Nav.html:27
	qw422016.N().S(`
</nav>`)
//line views/layout/Nav.html:28
}

//line views/layout/Nav.html:28
func WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:28
	StreamNav(qw422016, as, ps)
//line views/layout/Nav.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:28
}

//line views/layout/Nav.html:28
func Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:28
	WriteNav(qb422016, as, ps)
//line views/layout/Nav.html:28
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:28
	return qs422016
//line views/layout/Nav.html:28
}

//line views/layout/Nav.html:30
func StreamNavItems(qw422016 *qt422016.Writer, m menu.Items, breadcrumbs cmenu.Breadcrumbs) {
//line views/layout/Nav.html:31
	for idx, bc := range breadcrumbs {
//line views/layout/Nav.html:33
		i := m.GetByPath(breadcrumbs[:idx+1])

//line views/layout/Nav.html:35
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:35
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:37
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:38
		if i == nil {
//line views/layout/Nav.html:40
			bcLink := ""
			if strings.Contains(bc, "||") {
				bci := strings.Index(bc, "||")
				bcLink = bc[bci+2:]
				bc = bc[:bci]
			}

//line views/layout/Nav.html:46
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:47
			qw422016.E().S(bcLink)
//line views/layout/Nav.html:47
			qw422016.N().S(`">`)
//line views/layout/Nav.html:47
			qw422016.E().S(bc)
//line views/layout/Nav.html:47
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:48
		} else {
//line views/layout/Nav.html:48
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:49
			qw422016.E().S(i.Route)
//line views/layout/Nav.html:49
			qw422016.N().S(`">`)
//line views/layout/Nav.html:49
			qw422016.E().S(i.Title)
//line views/layout/Nav.html:49
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:50
		}
//line views/layout/Nav.html:51
	}
//line views/layout/Nav.html:52
}

//line views/layout/Nav.html:52
func WriteNavItems(qq422016 qtio422016.Writer, m menu.Items, breadcrumbs cmenu.Breadcrumbs) {
//line views/layout/Nav.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:52
	StreamNavItems(qw422016, m, breadcrumbs)
//line views/layout/Nav.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:52
}

//line views/layout/Nav.html:52
func NavItems(m menu.Items, breadcrumbs cmenu.Breadcrumbs) string {
//line views/layout/Nav.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:52
	WriteNavItems(qb422016, m, breadcrumbs)
//line views/layout/Nav.html:52
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:52
	return qs422016
//line views/layout/Nav.html:52
}
