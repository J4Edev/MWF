//go:build !windows
package power
import ("context"; "fmt")
type unavailable struct{}
func New() Manager{return unavailable{}}
func(unavailable)CurrentPlan(context.Context)(string,error){return "",fmt.Errorf("power management is only available on Windows")}
func(unavailable)SetPlan(context.Context,string)error{return fmt.Errorf("power management is only available on Windows")}
