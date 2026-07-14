package engine
import("testing";"MWF/internal/actions")
func TestValidatorAcceptsRegistrySet(t *testing.T){got,err:=Validator([]actions.Action{{Type:actions.RegistrySet,Key:`HKCU\Software\MWF`,Name:"Enabled",Value:"1"}});if err!=nil||len(got)!=1{t.Fatalf("got %v, %v",got,err)}}
func TestValidatorRejectsUnknownAction(t *testing.T){if _,err:=Validator([]actions.Action{{Type:"dangerous"}});err==nil{t.Fatal("expected error")}}
