package engine

import (
	"fmt"
)

type Executor struct {
	DryRun bool
}

func NewExecutor(dry bool) *Executor {
	return &Executor{DryRun: dry}
}

func (e *Executor) Execute(actions []Action) {
	for _, a := range actions {

		if e.DryRun {
			fmt.Printf("[DRY-RUN] %s -> %+v\n", a.Type, a)
			continue
		}

		switch a.Type {

		case "power_plan":
			execPowerPlan(a)

		case "startup_remove":
			execStartupRemove(a)

		case "windows_setting":
			execWindowsSetting(a)
		}
	}
}
