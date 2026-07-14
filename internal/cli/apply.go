package cli

import (
	"context"
	"fmt"

	"MWF/internal/engine"
	"MWF/internal/profiles"
	"MWF/internal/snapshot"

	"github.com/spf13/cobra"
)

var dryRun bool

var applyCmd = &cobra.Command{
	Use:   "apply [profile]",
	Short: "Apply a system profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]

		p, err := profiles.Load(profileName)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		executor := engine.NewExecutor(dryRun)

		if _, err := engine.RunPipeline(context.Background(), p, executor, snapshot.New("")); err != nil {
			fmt.Println("error:", err)
		}
	},
}

func init() {
	applyCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without applying")
}
