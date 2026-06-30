package engine

import "fmt"

// NOTE: These are SAFE STUBS for v0.2

func execPowerPlan(a Action) {
	fmt.Println("🔋 Setting power plan:", a.Value)
}

func execStartupRemove(a Action) {
	fmt.Println("🚀 Removing startup apps:", a.Target)
}

func execWindowsSetting(a Action) {
	fmt.Println("🪟 Applying windows setting:", a.Value)
}
