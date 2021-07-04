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
	return p.Name
}

func (p *Profile) Clone() *Profile {
	return &Profile{Name:  p.Name, Mode:  p.Mode, Theme: p.Theme}
}

func (p *Profile) AuthString(a auth.Sessions) string {
	msg := fmt.Sprintf("signed in as %s", p.String())
	if len(a) == 0 {
		if p.Name == DefaultProfile.Name {
			return "click to sign in"
		}
		return msg
	}
	return fmt.Sprintf("%s using [%s]", msg, a.String())
}
