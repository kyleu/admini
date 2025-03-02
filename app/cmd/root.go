package cmd

import (
	"fmt"

	"github.com/muesli/coral"

	"admini.dev/admini/app/util"
)

func rootF(*coral.Command, []string) error {
	// $PF_SECTION_START(rootAction)$
	return startServer(_flags)
	// $PF_SECTION_END(rootAction)$
}

func rootCmd() *coral.Command {
	short := fmt.Sprintf("%s %s - %s", util.AppName, _buildInfo.Version, util.AppSummary)
	ret := &coral.Command{Use: util.AppKey, Short: short, RunE: rootF}
	ret.AddCommand(serverCmd(), siteCmd(), allCmd(), upgradeCmd())
	// $PF_SECTION_START(cmds)$
	// $PF_SECTION_END(cmds)$
	ret.AddCommand(versionCmd())

	ret.PersistentFlags().StringVarP(&_flags.WorkingDir, "working_dir", "w", ".", "directory for projects, defaults to current dir")
	ret.PersistentFlags().StringVarP(&_flags.ConfigDir, "config_dir", "c", "", "directory for configuration, defaults to system config dir")
	ret.PersistentFlags().BoolVarP(&_flags.Debug, "verbose", "v", false, "enables verbose logging and additional checks")
	ret.PersistentFlags().StringVarP(&_flags.Address, "addr", "a", "127.0.0.1", "address to listen on, defaults to [127.0.0.1]")
	ret.PersistentFlags().Uint16VarP(&_flags.Port, "port", "p", util.AppPort, fmt.Sprintf("port to listen on, defaults to [%d]", util.AppPort))

	return ret
}
