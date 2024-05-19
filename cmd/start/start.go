package start

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

type Server struct {
	// Config     *Config
	// DB         *database.DB
	// KeyStorage crypto.KeyStorage
	// Keys       *encryption.EncryptionKeys
	// Eventstore *eventstore.Eventstore
	// Queries    *query.Queries
	// AuthzRepo  authz_repo.Repository
	// Storage    static.Storage
	// Commands   *command.Commands
	// Router     *mux.Router
	// TLSConfig  *tls.Config
	Shutdown chan<- os.Signal
}

func New(server chan<- *Server) *cobra.Command {
	start := &cobra.Command{
		Use:   "start",
		Short: "starts isotope",
		Long:  "starts isotope instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			// err := cmd_tls.ModeFromFlag(cmd)
			// if err != nil {
			// 	return err
			// }
			// config := MustNewConfig(viper.GetViper())
			// masterKey, err := key.MasterKey(cmd)
			// if err != nil {
			// 	return err
			// }
			return startIsotope(cmd.Context(), server)
		},
	}

	startFlags(start)

	return start
}

func startIsotope(ctx context.Context, server chan<- *Server) error {

}
