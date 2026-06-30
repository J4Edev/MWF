package engine

import "MWF/internal/profiles"

type Action struct {
	Type   string
	Target []string
	Value  string
}

func Planner(p *profiles.Profile) []Action {
	var actions []Action

	for _, a := range p.Actions {
		actions = append(actions, Action{
			Type:   a.Type,
			Target: a.Target,
			Value:  a.Value,
		})
	}

	return actions
}
