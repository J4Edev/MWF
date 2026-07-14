package profiles

import (
 "os"
 "path/filepath"
 "testing"
 "MWF/internal/actions"
)
func TestLoadParsesTypedActions(t *testing.T){dir:=t.TempDir();old,err:=os.Getwd();if err!=nil{t.Fatal(err)};defer os.Chdir(old);if err=os.MkdirAll(filepath.Join(dir,"configs","profiles"),0755);err!=nil{t.Fatal(err)};data:=[]byte("name: test\nactions:\n  - type: registry_set\n    key: HKCU\\Software\\MWF\n    name: Enabled\n    value: '1'\n    risk: low\n");if err=os.WriteFile(filepath.Join(dir,"configs","profiles","test.yaml"),data,0644);err!=nil{t.Fatal(err)};if err=os.Chdir(dir);err!=nil{t.Fatal(err)};p,err:=Load("test");if err!=nil{t.Fatal(err)};if len(p.Actions)!=1||p.Actions[0].Type!=actions.RegistrySet||p.Actions[0].Key!="HKCU\\Software\\MWF"{t.Fatalf("unexpected actions: %#v",p.Actions)}}
