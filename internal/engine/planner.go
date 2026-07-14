package engine

import (
 "MWF/internal/actions"
 "MWF/internal/profiles"
)

func Planner(p *profiles.Profile) []actions.Action { return append([]actions.Action(nil), p.Actions...) }
