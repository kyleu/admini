package auth

import (
	"sort"
	"strings"

	"github.com/kyleu/admini/app/util"
)

type Session struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
}

func sessionFromString(s string) *Session {
	p, e := util.SplitString(s, ':', true)
	return &Session{Provider: p, Email: e}
}

func (r Session) String() string {
	return r.Provider + ":" + r.Email
}

type Sessions []*Session

func (s Sessions) String() string {
	ret := make([]string, 0, len(s))
	for _, x := range s {
		ret = append(ret, x.String())
	}
	return strings.Join(ret, ",")
}

func (s Sessions) Sort() {
	sort.Slice(s, func(i, j int) bool {
		l := s[i]
		r := s[j]
		if l.Provider == r.Provider {
			return l.Email < r.Email
		}
		return l.Provider < r.Provider
	})
}

func (s Sessions) GetByProvider(p string) Sessions {
	var ret Sessions
	for _, x := range s {
		if x.Provider == p {
			ret = append(ret, x)
		}
	}
	return ret
}

func SessionsFromString(s string) Sessions {
	split := util.SplitAndTrim(s, ",")
	ret := make(Sessions, 0, len(split))
	for _, x := range split {
		ret = append(ret, sessionFromString(x))
	}
	return ret
}
