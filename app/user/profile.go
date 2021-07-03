package user

import (
	"fmt"

	"github.com/kyleu/admini/app/auth"
)

type Profile struct {
	Name  string `json:"name"`
	Mode  string `json:"mode,omitempty"`
	Theme string `json:"theme,omitempty"`
}

var DefaultProfile = &Profile{Name: "Guest"}

func (p *Profile) String() string {
	if p.Theme == "" {
		return p.Name
	}
	return fmt.Sprintf("%s (%s)", p.Name, p.Theme)
}

func (p *Profile) Clone() *Profile {
	return &Profile{Name:  p.Name, Mode:  p.Mode, Theme: p.Theme}
}

func (p *Profile) AuthString(a auth.Sessions) string {
	if len(a) == 0 {
		return fmt.Sprintf("Signed in as %s", p.String())
	}
	return fmt.Sprintf("Signed in as %s using [%s]", p.String(), a.String())
}
