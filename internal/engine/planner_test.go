package engine

import (
 "testing"
 "MWF/internal/actions"
 "MWF/internal/profiles"
)
func TestPlannerPreservesTypedAction(t *testing.T){in:=actions.Action{Type:actions.PowerPlan,Value:"balanced"};got:=Planner(&profiles.Profile{Actions:[]actions.Action{in}});if len(got)!=1||got[0]!=in{t.Fatalf("got %#v",got)}}
