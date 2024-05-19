package start

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	startFlagSet = &pflag.FlagSet{}
)

func init() {
	startFlagSet.Uint16P("port", "p", 0, "port to run isotope on")
	// startFlagSet.String("externalDomain", "", "domain isotope will be exposed on")
	// startFlagSet.String("externalPort", "", "port isotope will be exposed on")
}

func startFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlagSet(startFlagSet)
	// logging.OnError(
	// 	viper.BindPFlags(startFlagSet),
	// ).Fatal("start flags")

	// tls.AddTLSModeFlag(cmd)
	// key.AddMasterKeyFlag(cmd)
}
