package engine

import "fmt"

func Validator(actions []Action) []Action {
	var valid []Action

	for _, a := range actions {
		switch a.Type {

		case "power_plan", "startup_remove", "windows_setting":
			valid = append(valid, a)

		default:
			fmt.Println("⚠ skipped unknown action:", a.Type)
		}
	}

	return valid
}
