package cli

import (
 "context"
 "fmt"
 "MWF/internal/engine"
 "MWF/internal/snapshot"
 "github.com/spf13/cobra"
)
var restoreCmd=&cobra.Command{Use:"restore <snapshot-id>",Short:"Restore MWF-managed changes from a snapshot",Args:cobra.ExactArgs(1),Run:func(_ *cobra.Command,args []string){
 manager:=snapshot.New("");profile,items,err:=manager.Reverse(args[0]);if err!=nil{fmt.Println("error:",err);return};if len(items)==0{fmt.Println("Snapshot contains no reversible changes.");return}
 if _,err=engine.RunActions(context.Background(),"restore-"+profile,items,engine.NewExecutor(false),manager);err!=nil{fmt.Println("error:",err)}
}}
