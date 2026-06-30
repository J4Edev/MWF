package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run system diagnostics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MWF Doctor Mode")
		fmt.Println("- CPU: OK (stub)")
		fmt.Println("- Memory: OK (stub)")
		fmt.Println("- Services: OK (stub)")
	},
}
