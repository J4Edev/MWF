package engine

import (
 "context"
 "fmt"
 "MWF/internal/actions"
 "MWF/internal/profiles"
 "MWF/internal/snapshot"
)

// RunPipeline implements Profile -> Planner -> Validator -> Snapshot -> Executor -> Reporter.
func RunPipeline(ctx context.Context, p *profiles.Profile, exec *Executor, snapshots *snapshot.Manager) (string,error) {
 return RunActions(ctx,p.Name,Planner(p),exec,snapshots)
}
func RunActions(ctx context.Context, profile string, planned []actions.Action, exec *Executor, snapshots *snapshot.Manager) (string,error) {
 valid,err:=Validator(planned); if err!=nil{return "",err}
 Reporter{}.Plan(valid,exec.DryRun)
 var id string
 if !exec.DryRun { id,err=snapshots.Create(ctx,profile,valid,snapshot.Systems{Registry:exec.Systems.Registry,Power:exec.Systems.Power,Startup:exec.Systems.Startup,Windows:exec.Systems.Windows}); if err!=nil{return "",fmt.Errorf("create snapshot: %w",err)} }
 if err=exec.Execute(ctx,valid);err!=nil{return id,err}
 Reporter{}.Complete(id,exec.DryRun)
 return id,nil
}
