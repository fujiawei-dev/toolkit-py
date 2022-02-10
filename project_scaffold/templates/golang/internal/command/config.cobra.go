{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/spf13/cobra"

	"{{GOLANG_MODULE}}/internal/config"
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Show settings from different sources",
	RunE:  configAction,
}

func configAction(cmd *cobra.Command, args []string) error {
	return config.Conf().Print()
}
