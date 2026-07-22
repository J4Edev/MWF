package engine

import (
	"MWF/internal/actions"
	"MWF/internal/system/power"
	"MWF/internal/system/registry"
	"MWF/internal/system/startup"
	windowssettings "MWF/internal/system/windows"
	"context"
	"fmt"
)

// Systems is injected so the executor orchestrates actions without knowing how
// Windows APIs or commands are implemented.
type Systems struct {
	Registry registry.Manager
	Power    power.Manager
	Startup  startup.Manager
	Windows  windowssettings.Manager
}

func DefaultSystems() Systems {
	r := registry.New()
	return Systems{Registry: r, Power: power.New(), Startup: startup.New(r), Windows: windowssettings.New()}
}

type Executor struct {
	DryRun  bool
	Systems Systems
}

func NewExecutor(dry bool) *Executor { return &Executor{DryRun: dry, Systems: DefaultSystems()} }
func (e *Executor) Execute(ctx context.Context, items []actions.Action) error {
	for _, a := range items {
		if e.DryRun {
			fmt.Printf("[DRY-RUN] [%s] %s\n", a.EffectiveRisk(), a.String())
			continue
		}
		var err error
		switch a.Type {
		case actions.RegistrySet:
			err = e.Systems.Registry.Set(ctx, a.Key, a.Name, a.Value)
		case actions.RegistrySetDWORD:
			err = e.Systems.Registry.SetDWORD(ctx, a.Key, a.Name, a.Value)
		case actions.RegistryDelete:
			err = e.Systems.Registry.Delete(ctx, a.Key, a.Name)
		case actions.PowerPlan:
			err = e.Systems.Power.SetPlan(ctx, a.Value)
		case actions.StartupRemove:
			for _, name := range a.Target {
				if err = e.Systems.Startup.Remove(ctx, name); err != nil {
					break
				}
			}
		case actions.WindowsSetting:
			err = e.Systems.Windows.Apply(ctx, a.Value, a.Value)
		}
		if err != nil {
			return fmt.Errorf("execute %s: %w", a.String(), err)
		}
	}
	return nil
}
