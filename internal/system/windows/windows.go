// Package windows provides a narrow extension point for supported Windows
// settings. Keeping it separate prevents profiles from becoming registry hacks.
package windows

import (
 "context"
 "fmt"
)
type Manager interface { Read(context.Context,string)(string,bool,error); Apply(context.Context,string,string) error }
type ManagerStub struct{}
func New() Manager{return ManagerStub{}}
func(ManagerStub)Read(context.Context,string)(string,bool,error){return "",false,nil}
func(ManagerStub)Apply(_ context.Context, name, _ string)error{return fmt.Errorf("Windows setting %q is not implemented",name)}
