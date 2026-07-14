package startup

import (
	"context"
	"MWF/internal/system/registry"
)

const runKey = `HKCU\Software\Microsoft\Windows\CurrentVersion\Run`

type Application struct { Name string; Command string }
type Manager interface { List(context.Context) ([]Application, error); Lookup(context.Context, string) (registry.Value, error); Remove(context.Context, string) error }

// RegistryManager implements the user Run-key portion of startup management.
// More startup locations can be added without leaking registry details to engine.
type RegistryManager struct { Registry registry.Manager }
func New(r registry.Manager) Manager { return RegistryManager{Registry:r} }
func (m RegistryManager) Lookup(ctx context.Context, name string) (registry.Value,error) { return m.Registry.Read(ctx,runKey,name) }
func (m RegistryManager) Remove(ctx context.Context, name string) error { return m.Registry.Delete(ctx,runKey,name) }
func (m RegistryManager) List(ctx context.Context) ([]Application,error) { values,err:=m.Registry.Values(ctx,runKey);if err!=nil{return nil,err};apps:=make([]Application,0,len(values));for name,command:=range values{apps=append(apps,Application{Name:name,Command:command})};return apps,nil }
func RunKey() string { return runKey }
