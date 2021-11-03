package starportcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewPlugins creates plugin commands.
func NewPlugins() *cobra.Command {
	c := &cobra.Command{
		Use:   "plugins",
		Short: "Plugin management command",
		Args:  cobra.ExactArgs(0),
		RunE:  pluginHandler,
	}

	// TODO: Add sub commands.
	// install
	// run...

	return c
}

func pluginHandler(cmd *cobra.Command, args []string) error {
	fmt.Println("Welcome to new plugin system")
	return nil
}
