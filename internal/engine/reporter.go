package engine
import ("fmt"; "MWF/internal/actions")
type Reporter struct{}
func(Reporter)Plan(items []actions.Action,dry bool){for _,a:=range items{prefix:="PLAN";if dry{prefix="DRY-RUN"};fmt.Printf("[%s] [%s] %s\n",prefix,a.EffectiveRisk(),a.String())}}
func(Reporter)Complete(id string,dry bool){if dry{fmt.Println("Dry run complete; no system changes were made.");return};fmt.Println("Done. Snapshot:",id)}
