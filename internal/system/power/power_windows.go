//go:build windows
package power

import (
 "context"
 "fmt"
 "os/exec"
 "strings"
)
type WindowsManager struct{}
func New() Manager { return WindowsManager{} }
func (WindowsManager) CurrentPlan(ctx context.Context) (string, error) { b,e:=exec.CommandContext(ctx,"powercfg","/getactivescheme").Output(); if e!=nil{return "",e}; s:=strings.ToLower(string(b)); for _, p:=range []string{"balanced","high performance","power saver"}{if strings.Contains(s,p){return p,nil}}; return strings.TrimSpace(string(b)),nil }
func (WindowsManager) SetPlan(ctx context.Context, plan string) error { aliases:=map[string]string{"balanced":"SCHEME_BALANCED","high_performance":"SCHEME_MIN","power_saver":"SCHEME_MAX"}; v,ok:=aliases[strings.ToLower(plan)]; if !ok{return fmt.Errorf("unsupported power plan %q",plan)}; return exec.CommandContext(ctx,"powercfg","/setactive",v).Run() }
