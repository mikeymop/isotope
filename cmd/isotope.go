package cmd

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
)

func New(out io.Writer, in io.Reader, args []string, server chan<- *start.Server) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "isotope",
		Short: "The isotope CLI lets you interact with isotope",
		Long:  `The isotope CLI lets you interact with isotope`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no additional command provided")
		},
		Version: build.Version(),
	}

	cmd.AddCommand(
		// admin.New(), //is now deprecated, remove later on
		// initialise.New(),
		// setup.New(),
		start.New(server),
		start.NewStartFromInit(server),
		start.NewStartFromSetup(server),
		// key.New(),
		// ready.New(),
	)

	cmd.InitDefaultVersionFlag()

	return cmd
}
