package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply [profile]",
	Short: "Apply a system profile",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		profile := args[0]

		if dryRun {
			fmt.Println("[DRY RUN] Applying profile:", profile)
			return
		}

		fmt.Println("Applying profile:", profile)

		if profile == "default" {
			fmt.Println("-> Applying safe baseline configuration")
		} else {
			fmt.Println("-> Profile not fully implemented yet")
		}
	},
}

var dryRun bool

func init() {
	applyCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without applying")
}
