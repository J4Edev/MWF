// Package actions contains the engine-facing representation of profile changes.
// It deliberately contains no Windows implementation details.
package actions

import (
	"fmt"
	"strings"
)

type Type string

const (
	RegistrySet      Type = "registry_set"
	RegistrySetDWORD Type = "registry_set_dword"
	RegistryDelete   Type = "registry_delete"
	PowerPlan        Type = "power_plan"
	StartupRemove    Type = "startup_remove"
	WindowsSetting   Type = "windows_setting"
)

type Risk string

const (
	RiskLow    Risk = "low"
	RiskMedium Risk = "medium"
	RiskHigh   Risk = "high"
)

// Action is the normalized profile action passed through the engine.
// Key and Name are used by registry actions. Target remains for existing
// profile actions such as startup_remove.
type Action struct {
	Type   Type     `yaml:"type" json:"type"`
	Risk   Risk     `yaml:"risk,omitempty" json:"risk,omitempty"`
	Target []string `yaml:"target,omitempty" json:"target,omitempty"`
	Key    string   `yaml:"key,omitempty" json:"key,omitempty"`
	Name   string   `yaml:"name,omitempty" json:"name,omitempty"`
	Value  string   `yaml:"value,omitempty" json:"value,omitempty"`
}

func (a Action) String() string {
	if a.Key != "" {
		return fmt.Sprintf("%s %s/%s", a.Type, a.Key, a.Name)
	}
	return fmt.Sprintf("%s %s", a.Type, a.Value)
}

func (a Action) EffectiveRisk() Risk {
	if a.Risk != "" {
		return a.Risk
	}
	switch a.Type {
	case RegistryDelete, StartupRemove:
		return RiskMedium
	default:
		return RiskLow
	}
}

func (a Action) ValidRegistryKey() bool {
	return strings.Contains(a.Key, `\`) && (strings.HasPrefix(strings.ToUpper(a.Key), `HKCU\`) || strings.HasPrefix(strings.ToUpper(a.Key), `HKLM\`) || strings.HasPrefix(strings.ToUpper(a.Key), `HKEY_CURRENT_USER\`) || strings.HasPrefix(strings.ToUpper(a.Key), `HKEY_LOCAL_MACHINE\`))
}
