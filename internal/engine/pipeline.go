package engine

import (
	"fmt"

	"MWF/internal/profiles"
)

func RunPipeline(p *profiles.Profile, exec *Executor) {
	fmt.Println("▶ Loading profile:", p.Name)

	actions := Planner(p)

	validActions := Validator(actions)

	fmt.Println("▶ Executing actions...")

	exec.Execute(validActions)

	fmt.Println("✔ Done")
}
