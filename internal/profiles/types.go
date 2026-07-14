package profiles

import "MWF/internal/actions"

type Profile struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description"`
	Actions     []actions.Action `yaml:"actions"`
}
