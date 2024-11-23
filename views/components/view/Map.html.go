// Code generated by qtc from "Map.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/Map.html:1
package view

//line views/components/view/Map.html:1
import (
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
)

//line views/components/view/Map.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Map.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Map.html:6
func StreamMap(qw422016 *qt422016.Writer, preserveWhitespace bool, m util.ValueMap, ps *cutil.PageState) {
//line views/components/view/Map.html:7
	if m == nil {
//line views/components/view/Map.html:7
		qw422016.N().S(`<em>no result</em>`)
//line views/components/view/Map.html:9
	} else if len(m) == 0 {
//line views/components/view/Map.html:9
		qw422016.N().S(`<em>empty result</em>`)
//line views/components/view/Map.html:11
	} else {
//line views/components/view/Map.html:11
		qw422016.N().S(`<div class="overflow full-width bl"><table class="expanded"><tbody>`)
//line views/components/view/Map.html:15
		for _, k := range m.Keys() {
//line views/components/view/Map.html:15
			qw422016.N().S(`<tr><th class="shrink">`)
//line views/components/view/Map.html:17
			qw422016.E().S(k)
//line views/components/view/Map.html:17
			qw422016.N().S(`</th>`)
//line views/components/view/Map.html:18
			if preserveWhitespace {
//line views/components/view/Map.html:18
				qw422016.N().S(`<td class="prews">`)
//line views/components/view/Map.html:19
				StreamAny(qw422016, m[k], ps)
//line views/components/view/Map.html:19
				qw422016.N().S(`</td>`)
//line views/components/view/Map.html:20
			} else {
//line views/components/view/Map.html:20
				qw422016.N().S(`<td>`)
//line views/components/view/Map.html:21
				StreamAny(qw422016, m[k], ps)
//line views/components/view/Map.html:21
				qw422016.N().S(`</td>`)
//line views/components/view/Map.html:22
			}
//line views/components/view/Map.html:22
			qw422016.N().S(`</tr>`)
//line views/components/view/Map.html:24
		}
//line views/components/view/Map.html:24
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Map.html:28
	}
//line views/components/view/Map.html:29
}

//line views/components/view/Map.html:29
func WriteMap(qq422016 qtio422016.Writer, preserveWhitespace bool, m util.ValueMap, ps *cutil.PageState) {
//line views/components/view/Map.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Map.html:29
	StreamMap(qw422016, preserveWhitespace, m, ps)
//line views/components/view/Map.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Map.html:29
}

//line views/components/view/Map.html:29
func Map(preserveWhitespace bool, m util.ValueMap, ps *cutil.PageState) string {
//line views/components/view/Map.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Map.html:29
	WriteMap(qb422016, preserveWhitespace, m, ps)
//line views/components/view/Map.html:29
	qs422016 := string(qb422016.B)
//line views/components/view/Map.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Map.html:29
	return qs422016
//line views/components/view/Map.html:29
}

//line views/components/view/Map.html:31
func StreamMapKeys(qw422016 *qt422016.Writer, m util.ValueMap) {
//line views/components/view/Map.html:32
	if m == nil || len(m) == 0 {
//line views/components/view/Map.html:32
		qw422016.N().S(`<em>no keys</em>`)
//line views/components/view/Map.html:34
	} else {
//line views/components/view/Map.html:35
		StreamTags(qw422016, m.Keys(), nil)
//line views/components/view/Map.html:36
	}
//line views/components/view/Map.html:37
}

//line views/components/view/Map.html:37
func WriteMapKeys(qq422016 qtio422016.Writer, m util.ValueMap) {
//line views/components/view/Map.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Map.html:37
	StreamMapKeys(qw422016, m)
//line views/components/view/Map.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Map.html:37
}

//line views/components/view/Map.html:37
func MapKeys(m util.ValueMap) string {
//line views/components/view/Map.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Map.html:37
	WriteMapKeys(qb422016, m)
//line views/components/view/Map.html:37
	qs422016 := string(qb422016.B)
//line views/components/view/Map.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Map.html:37
	return qs422016
//line views/components/view/Map.html:37
}

//line views/components/view/Map.html:39
func StreamMapArray(qw422016 *qt422016.Writer, preserveWhitespace bool, ps *cutil.PageState, maps ...util.ValueMap) {
//line views/components/view/Map.html:40
	if len(maps) == 0 {
//line views/components/view/Map.html:40
		qw422016.N().S(`<em>no results</em>`)
//line views/components/view/Map.html:42
	} else {
//line views/components/view/Map.html:42
		qw422016.N().S(`<div class="overflow full-width"><table><thead><tr>`)
//line views/components/view/Map.html:47
		for _, k := range maps[0].Keys() {
//line views/components/view/Map.html:47
			qw422016.N().S(`<th>`)
//line views/components/view/Map.html:48
			qw422016.E().S(k)
//line views/components/view/Map.html:48
			qw422016.N().S(`</th>`)
//line views/components/view/Map.html:49
		}
//line views/components/view/Map.html:49
		qw422016.N().S(`</tr></thead><tbody>`)
//line views/components/view/Map.html:53
		for _, m := range maps {
//line views/components/view/Map.html:53
			qw422016.N().S(`<tr>`)
//line views/components/view/Map.html:55
			for _, k := range m.Keys() {
//line views/components/view/Map.html:56
				if preserveWhitespace {
//line views/components/view/Map.html:56
					qw422016.N().S(`<td class="prews">`)
//line views/components/view/Map.html:57
					StreamAny(qw422016, m[k], ps)
//line views/components/view/Map.html:57
					qw422016.N().S(`</td>`)
//line views/components/view/Map.html:58
				} else {
//line views/components/view/Map.html:58
					qw422016.N().S(`<td>`)
//line views/components/view/Map.html:59
					StreamAny(qw422016, m[k], ps)
//line views/components/view/Map.html:59
					qw422016.N().S(`</td>`)
//line views/components/view/Map.html:60
				}
//line views/components/view/Map.html:61
			}
//line views/components/view/Map.html:61
			qw422016.N().S(`</tr>`)
//line views/components/view/Map.html:63
		}
//line views/components/view/Map.html:63
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Map.html:67
	}
//line views/components/view/Map.html:68
}

//line views/components/view/Map.html:68
func WriteMapArray(qq422016 qtio422016.Writer, preserveWhitespace bool, ps *cutil.PageState, maps ...util.ValueMap) {
//line views/components/view/Map.html:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Map.html:68
	StreamMapArray(qw422016, preserveWhitespace, ps, maps...)
//line views/components/view/Map.html:68
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Map.html:68
}

//line views/components/view/Map.html:68
func MapArray(preserveWhitespace bool, ps *cutil.PageState, maps ...util.ValueMap) string {
//line views/components/view/Map.html:68
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Map.html:68
	WriteMapArray(qb422016, preserveWhitespace, ps, maps...)
//line views/components/view/Map.html:68
	qs422016 := string(qb422016.B)
//line views/components/view/Map.html:68
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Map.html:68
	return qs422016
//line views/components/view/Map.html:68
}

//line views/components/view/Map.html:70
func StreamOrderedMap(qw422016 *qt422016.Writer, preserveWhitespace bool, m *util.OrderedMap[any], ps *cutil.PageState) {
//line views/components/view/Map.html:71
	if m == nil {
//line views/components/view/Map.html:71
		qw422016.N().S(`<em>no result</em>`)
//line views/components/view/Map.html:73
	} else if len(m.Map) == 0 {
//line views/components/view/Map.html:73
		qw422016.N().S(`<em>empty result</em>`)
//line views/components/view/Map.html:75
	} else {
//line views/components/view/Map.html:75
		qw422016.N().S(`<div class="overflow full-width bl"><table class="expanded"><tbody>`)
//line views/components/view/Map.html:79
		for _, k := range m.Order {
//line views/components/view/Map.html:79
			qw422016.N().S(`<tr><th class="shrink">`)
//line views/components/view/Map.html:81
			qw422016.E().S(k)
//line views/components/view/Map.html:81
			qw422016.N().S(`</th>`)
//line views/components/view/Map.html:82
			if preserveWhitespace {
//line views/components/view/Map.html:82
				qw422016.N().S(`<td class="prews">`)
//line views/components/view/Map.html:83
				StreamAny(qw422016, m.GetSimple(k), ps)
//line views/components/view/Map.html:83
				qw422016.N().S(`</td>`)
//line views/components/view/Map.html:84
			} else {
//line views/components/view/Map.html:84
				qw422016.N().S(`<td>`)
//line views/components/view/Map.html:85
				StreamAny(qw422016, m.GetSimple(k), ps)
//line views/components/view/Map.html:85
				qw422016.N().S(`</td>`)
//line views/components/view/Map.html:86
			}
//line views/components/view/Map.html:86
			qw422016.N().S(`</tr>`)
//line views/components/view/Map.html:88
		}
//line views/components/view/Map.html:88
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Map.html:92
	}
//line views/components/view/Map.html:93
}

//line views/components/view/Map.html:93
func WriteOrderedMap(qq422016 qtio422016.Writer, preserveWhitespace bool, m *util.OrderedMap[any], ps *cutil.PageState) {
//line views/components/view/Map.html:93
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Map.html:93
	StreamOrderedMap(qw422016, preserveWhitespace, m, ps)
//line views/components/view/Map.html:93
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Map.html:93
}

//line views/components/view/Map.html:93
func OrderedMap(preserveWhitespace bool, m *util.OrderedMap[any], ps *cutil.PageState) string {
//line views/components/view/Map.html:93
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Map.html:93
	WriteOrderedMap(qb422016, preserveWhitespace, m, ps)
//line views/components/view/Map.html:93
	qs422016 := string(qb422016.B)
//line views/components/view/Map.html:93
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Map.html:93
	return qs422016
//line views/components/view/Map.html:93
}

//line views/components/view/Map.html:95
func StreamOrderedMapArray(qw422016 *qt422016.Writer, preserveWhitespace bool, ps *cutil.PageState, maps ...*util.OrderedMap[any]) {
//line views/components/view/Map.html:96
	if len(maps) == 0 {
//line views/components/view/Map.html:96
		qw422016.N().S(`<em>no results</em>`)
//line views/components/view/Map.html:98
	} else {
//line views/components/view/Map.html:98
		qw422016.N().S(`<div class="overflow full-width"><table><thead><tr>`)
//line views/components/view/Map.html:103
		for _, k := range maps[0].Order {
//line views/components/view/Map.html:103
			qw422016.N().S(`<th>`)
//line views/components/view/Map.html:104
			qw422016.E().S(k)
//line views/components/view/Map.html:104
			qw422016.N().S(`</th>`)
//line views/components/view/Map.html:105
		}
//line views/components/view/Map.html:105
		qw422016.N().S(`</tr></thead><tbody>`)
//line views/components/view/Map.html:109
		for _, m := range maps {
//line views/components/view/Map.html:109
			qw422016.N().S(`<tr>`)
//line views/components/view/Map.html:111
			for _, k := range m.Order {
//line views/components/view/Map.html:112
				if preserveWhitespace {
//line views/components/view/Map.html:112
					qw422016.N().S(`<td class="prews">`)
//line views/components/view/Map.html:113
					StreamAny(qw422016, m.GetSimple(k), ps)
//line views/components/view/Map.html:113
					qw422016.N().S(`</td>`)
//line views/components/view/Map.html:114
				} else {
//line views/components/view/Map.html:114
					qw422016.N().S(`<td>`)
//line views/components/view/Map.html:115
					StreamAny(qw422016, m.GetSimple(k), ps)
//line views/components/view/Map.html:115
					qw422016.N().S(`</td>`)
//line views/components/view/Map.html:116
				}
//line views/components/view/Map.html:117
			}
//line views/components/view/Map.html:117
			qw422016.N().S(`</tr>`)
//line views/components/view/Map.html:119
		}
//line views/components/view/Map.html:119
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Map.html:123
	}
//line views/components/view/Map.html:124
}

//line views/components/view/Map.html:124
func WriteOrderedMapArray(qq422016 qtio422016.Writer, preserveWhitespace bool, ps *cutil.PageState, maps ...*util.OrderedMap[any]) {
//line views/components/view/Map.html:124
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Map.html:124
	StreamOrderedMapArray(qw422016, preserveWhitespace, ps, maps...)
//line views/components/view/Map.html:124
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Map.html:124
}

//line views/components/view/Map.html:124
func OrderedMapArray(preserveWhitespace bool, ps *cutil.PageState, maps ...*util.OrderedMap[any]) string {
//line views/components/view/Map.html:124
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Map.html:124
	WriteOrderedMapArray(qb422016, preserveWhitespace, ps, maps...)
//line views/components/view/Map.html:124
	qs422016 := string(qb422016.B)
//line views/components/view/Map.html:124
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Map.html:124
	return qs422016
//line views/components/view/Map.html:124
}
