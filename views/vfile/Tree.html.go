// Code generated by qtc from "Tree.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vfile/Tree.html:2
package vfile

//line views/vfile/Tree.html:2
import (
	"path"
	"strings"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/filesystem"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/components"
)

//line views/vfile/Tree.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/Tree.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/Tree.html:13
func StreamTree(qw422016 *qt422016.Writer, t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Tree.html:14
	components.StreamIndent(qw422016, true, 2)
//line views/vfile/Tree.html:14
	qw422016.N().S(`<ul class="accordion">`)
//line views/vfile/Tree.html:16
	streamtreeNodes(qw422016, t.Nodes, "", urlPrefix, actions, as, ps, 2)
//line views/vfile/Tree.html:17
	components.StreamIndent(qw422016, true, 2)
//line views/vfile/Tree.html:17
	qw422016.N().S(`</ul>`)
//line views/vfile/Tree.html:19
}

//line views/vfile/Tree.html:19
func WriteTree(qq422016 qtio422016.Writer, t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Tree.html:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:19
	StreamTree(qw422016, t, urlPrefix, actions, as, ps)
//line views/vfile/Tree.html:19
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:19
}

//line views/vfile/Tree.html:19
func Tree(t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/Tree.html:19
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:19
	WriteTree(qb422016, t, urlPrefix, actions, as, ps)
//line views/vfile/Tree.html:19
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:19
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:19
	return qs422016
//line views/vfile/Tree.html:19
}

//line views/vfile/Tree.html:21
func streamtreeNode(qw422016 *qt422016.Writer, n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:23
	pathID := n.Name
	if pth != "" {
		pathID = pth + "/" + pathID
	}
	pathID = strings.ReplaceAll(pathID, "/", "__")

//line views/vfile/Tree.html:29
	components.StreamIndent(qw422016, true, indent)
//line views/vfile/Tree.html:29
	qw422016.N().S(`<li>`)
//line views/vfile/Tree.html:31
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:31
	qw422016.N().S(`<input id="accordion-`)
//line views/vfile/Tree.html:32
	qw422016.E().S(pathID)
//line views/vfile/Tree.html:32
	qw422016.N().S(`" type="checkbox" hidden="hidden" />`)
//line views/vfile/Tree.html:33
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:33
	qw422016.N().S(`<label for="accordion-`)
//line views/vfile/Tree.html:34
	qw422016.E().S(pathID)
//line views/vfile/Tree.html:34
	qw422016.N().S(`">`)
//line views/vfile/Tree.html:36
	var acts map[string]string
	if actions != nil {
		acts = actions(pth, n)
	}

//line views/vfile/Tree.html:40
	qw422016.N().S(`<div class="right">`)
//line views/vfile/Tree.html:42
	if n.Size > 0 {
//line views/vfile/Tree.html:42
		qw422016.N().S(`<em>`)
//line views/vfile/Tree.html:43
		qw422016.E().S(util.ByteSizeSI(int64(n.Size)))
//line views/vfile/Tree.html:43
		qw422016.N().S(`</em>`)
//line views/vfile/Tree.html:43
		qw422016.N().S(` `)
//line views/vfile/Tree.html:44
	}
//line views/vfile/Tree.html:45
	for k, v := range acts {
//line views/vfile/Tree.html:45
		qw422016.N().S(`<a href="`)
//line views/vfile/Tree.html:46
		qw422016.E().S(v)
//line views/vfile/Tree.html:46
		qw422016.N().S(`">`)
//line views/vfile/Tree.html:46
		qw422016.E().S(k)
//line views/vfile/Tree.html:46
		qw422016.N().S(`</a>`)
//line views/vfile/Tree.html:47
	}
//line views/vfile/Tree.html:47
	qw422016.N().S(`</div>`)
//line views/vfile/Tree.html:49
	components.StreamExpandCollapse(qw422016, indent+2, ps)
//line views/vfile/Tree.html:50
	if n.Dir {
//line views/vfile/Tree.html:51
		components.StreamSVGRef(qw422016, `folder`, 15, 15, ``, ps)
//line views/vfile/Tree.html:52
	} else {
//line views/vfile/Tree.html:53
		components.StreamSVGRef(qw422016, `file`, 15, 15, ``, ps)
//line views/vfile/Tree.html:54
	}
//line views/vfile/Tree.html:55
	qw422016.N().S(` `)
//line views/vfile/Tree.html:55
	qw422016.E().S(n.Name)
//line views/vfile/Tree.html:56
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:56
	qw422016.N().S(`</label>`)
//line views/vfile/Tree.html:58
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:58
	qw422016.N().S(`<div class="bd"><div><div>`)
//line views/vfile/Tree.html:60
	if len(n.Children) == 0 {
//line views/vfile/Tree.html:61
		components.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:61
		qw422016.N().S(`<div>`)
//line views/vfile/Tree.html:62
		qw422016.E().S(n.Name)
//line views/vfile/Tree.html:62
		qw422016.N().S(`</div>`)
//line views/vfile/Tree.html:63
	} else {
//line views/vfile/Tree.html:64
		components.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:64
		qw422016.N().S(`<ul class="accordion" style="margin-left:`)
//line views/vfile/Tree.html:65
		qw422016.N().D((indent / 3) * 6)
//line views/vfile/Tree.html:65
		qw422016.N().S(`px; margin-bottom: 0;">`)
//line views/vfile/Tree.html:66
		streamtreeNodes(qw422016, n.Children, path.Join(pth, n.Name), urlPrefix, actions, as, ps, indent+3)
//line views/vfile/Tree.html:67
		components.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:67
		qw422016.N().S(`</ul>`)
//line views/vfile/Tree.html:69
	}
//line views/vfile/Tree.html:70
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:70
	qw422016.N().S(`</div></div></div>`)
//line views/vfile/Tree.html:72
	components.StreamIndent(qw422016, true, indent)
//line views/vfile/Tree.html:72
	qw422016.N().S(`</li>`)
//line views/vfile/Tree.html:74
}

//line views/vfile/Tree.html:74
func writetreeNode(qq422016 qtio422016.Writer, n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:74
	streamtreeNode(qw422016, n, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:74
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:74
}

//line views/vfile/Tree.html:74
func treeNode(n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) string {
//line views/vfile/Tree.html:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:74
	writetreeNode(qb422016, n, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:74
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:74
	return qs422016
//line views/vfile/Tree.html:74
}

//line views/vfile/Tree.html:76
func streamtreeNodes(qw422016 *qt422016.Writer, nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:77
	for _, n := range nodes {
//line views/vfile/Tree.html:78
		streamtreeNode(qw422016, n, pth, urlPrefix, actions, as, ps, indent+1)
//line views/vfile/Tree.html:79
	}
//line views/vfile/Tree.html:80
}

//line views/vfile/Tree.html:80
func writetreeNodes(qq422016 qtio422016.Writer, nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:80
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:80
	streamtreeNodes(qw422016, nodes, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:80
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:80
}

//line views/vfile/Tree.html:80
func treeNodes(nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) string {
//line views/vfile/Tree.html:80
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:80
	writetreeNodes(qb422016, nodes, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:80
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:80
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:80
	return qs422016
//line views/vfile/Tree.html:80
}
