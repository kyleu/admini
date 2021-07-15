package cmd

import (
	"fmt"

	"github.com/kyleu/admini/app/util"
	"github.com/spf13/cobra"
)

func rootF(*cobra.Command, []string) error {
	return startServer()
}

func rootCmd() *cobra.Command {
	ret := &cobra.Command{Use: util.AppKey, Short: util.AppSummary, RunE: rootF}
	ret.AddCommand(serverCmd(), siteCmd(), allCmd())

	ret.PersistentFlags().StringVarP(&_flags.ConfigDir, "dir", "d", "", "directory for configuration, defaults to system config dir")
	ret.PersistentFlags().BoolVarP(&_flags.Debug, "verbose", "v", false, "enables verbose logging and additional checks")
	ret.PersistentFlags().BoolVarP(&_flags.JSON, "json", "j", false, "enables json logging")
	ret.PersistentFlags().StringVarP(&_flags.Address, "addr", "a", "127.0.0.1", "address to listen on, defaults to [127.0.0.1]")
	ret.PersistentFlags().Uint16VarP(&_flags.Port, "port", "p", util.AppPort, fmt.Sprintf("port to listen on, defaults to [%d]", util.AppPort))

	return ret
}
