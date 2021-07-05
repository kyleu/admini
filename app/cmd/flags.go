package cmd

import (
	"fmt"

	"github.com/kyleu/admini/app/util"
	"github.com/spf13/pflag"
)

type Flags struct {
	Address   string
	Port      int
	ConfigDir string
	Debug     bool
	JSON      bool
	Mode      string
}

func (f *Flags) Addr() string {
	if f.Port == 0 {
		f.Port = util.AppPort
	}
	return fmt.Sprintf("%s:%d", f.Address, f.Port)

}

func parseFlags() *Flags {
	ret := &Flags{}
	pflag.StringVarP(&ret.Address, "addr", "a", "127.0.0.1", "address to listen on, defaults to [127.0.0.1]")
	pflag.IntVarP(&ret.Port, "port", "p", util.AppPort, fmt.Sprintf("port to listen on, defaults to [%d]", util.AppPort))
	pflag.StringVarP(&ret.ConfigDir, "dir", "d", "", "directory for configuration, defaults to system config dir")
	pflag.BoolVarP(&ret.Debug, "verbose", "v", false, "enables verbose logging and additional checks")
	pflag.BoolVarP(&ret.JSON, "json", "j", false, "enables json logging")
	pflag.StringVarP(&ret.Mode, "mode", "m", "app", "determines startup behavior, you probably don't want this")
	pflag.Parse()
	return ret
}
