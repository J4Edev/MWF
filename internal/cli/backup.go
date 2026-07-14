package cli

import (
 "context"
 "fmt"
 "MWF/internal/engine"
 "MWF/internal/profiles"
 "MWF/internal/snapshot"
 "github.com/spf13/cobra"
)
// backup records the current state for a profile without applying it.
var backupCmd=&cobra.Command{Use:"backup [profile]",Short:"Create a snapshot of MWF-managed profile state",Args:cobra.MaximumNArgs(1),Run:func(_ *cobra.Command,args []string){
 name:="default";if len(args)==1{name=args[0]};p,err:=profiles.Load(name);if err!=nil{fmt.Println("error:",err);return};items,err:=engine.Validator(engine.Planner(p));if err!=nil{fmt.Println("error:",err);return};exec:=engine.NewExecutor(false);id,err:=snapshot.New("").Create(context.Background(),p.Name,items,snapshot.Systems{Registry:exec.Systems.Registry,Power:exec.Systems.Power,Startup:exec.Systems.Startup,Windows:exec.Systems.Windows});if err!=nil{fmt.Println("error:",err);return};fmt.Println("Snapshot created:",id)
}}
