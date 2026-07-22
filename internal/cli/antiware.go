package cli

import (
	"context"
	"fmt"

	"MWF/internal/engine"
	"MWF/internal/snapshot"
	"MWF/internal/system/antiware"

	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(newAntiwareCmd()) }

func newAntiwareCmd() *cobra.Command {
	var dryRun bool
	cmd := &cobra.Command{Use: "antiware", Short: "Manage automatic OEM companion-software suggestions"}
	enable := &cobra.Command{
		Use: "enable", Short: "Prevent Windows device metadata from being retrieved online",
		Run: func(_ *cobra.Command, _ []string) { runAntiware(true, dryRun) },
	}
	disable := &cobra.Command{
		Use: "disable", Short: "Restore the default device-metadata behavior",
		Run: func(_ *cobra.Command, _ []string) { runAntiware(false, dryRun) },
	}
	status := &cobra.Command{
		Use: "status", Short: "Show whether Antiware is enabled",
		Run: func(_ *cobra.Command, _ []string) {
			status, err := newAntiwareManager(false).Status(context.Background())
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			if status.Enabled {
				fmt.Println("Antiware is enabled.")
				return
			}
			if status.Configured {
				fmt.Println("Antiware is disabled (policy value is not enabled).")
				return
			}
			fmt.Println("Antiware is disabled (Windows default behavior).")
		},
	}
	enable.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without applying")
	disable.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without applying")
	cmd.AddCommand(enable, disable, status)
	return cmd
}

func runAntiware(enable, dryRun bool) {
	manager := newAntiwareManager(dryRun)
	var err error
	if enable {
		_, err = manager.Enable(context.Background())
	} else {
		_, err = manager.Disable(context.Background())
	}
	if err != nil {
		fmt.Println("error:", err)
	}
}

func newAntiwareManager(dryRun bool) *antiware.Manager {
	return antiware.New(engine.NewExecutor(dryRun), snapshot.New(""))
}
