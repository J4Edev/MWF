package engine

import (
 "fmt"
 "MWF/internal/actions"
)

// Validator is intentionally policy-only: it never changes Windows state.
func Validator(actionsToCheck []actions.Action) ([]actions.Action, error) {
 valid:=make([]actions.Action,0,len(actionsToCheck))
 for _, a:=range actionsToCheck {
  switch a.Type {
  case actions.RegistrySet,actions.RegistryDelete:
   if !a.ValidRegistryKey(){return nil,fmt.Errorf("%s requires an HKCU or HKLM registry key",a.Type)}; if a.Name==""{return nil,fmt.Errorf("%s requires a value name",a.Type)}
  case actions.PowerPlan:
   if a.Value==""{return nil,fmt.Errorf("power_plan requires a value")}
  case actions.StartupRemove:
   if len(a.Target)==0{return nil,fmt.Errorf("startup_remove requires at least one target")}
  case actions.WindowsSetting:
   if a.Value==""{return nil,fmt.Errorf("windows_setting requires a value")}
  default: return nil,fmt.Errorf("unsupported action type %q",a.Type)
  }
  valid=append(valid,a)
 }
 return valid,nil
}
