// Package antiware manages Windows' documented device-metadata policy.
package antiware

import (
	"context"
	"fmt"

	"MWF/internal/actions"
	"MWF/internal/engine"
	"MWF/internal/snapshot"
)

const (
	policyKey  = `HKLM\SOFTWARE\Policies\Microsoft\Windows\Device Metadata`
	policyName = "PreventDeviceMetadataFromNetwork"
)

// Manager orchestrates Antiware changes through the shared engine pipeline.
type Manager struct {
	executor  *engine.Executor
	snapshots *snapshot.Manager
}

// Status describes the current Antiware policy state.
type Status struct {
	Enabled    bool
	Configured bool
}

func New(executor *engine.Executor, snapshots *snapshot.Manager) *Manager {
	return &Manager{executor: executor, snapshots: snapshots}
}

func (m *Manager) Enable(ctx context.Context) (string, error) {
	return m.run(ctx, "antiware-enable", enabledActions())
}

func (m *Manager) Disable(ctx context.Context) (string, error) {
	return m.run(ctx, "antiware-disable", disabledActions())
}

func (m *Manager) Status(ctx context.Context) (Status, error) {
	if m.executor == nil || m.executor.Systems.Registry == nil {
		return Status{}, fmt.Errorf("Antiware registry system is not configured")
	}
	v, err := m.executor.Systems.Registry.Read(ctx, policyKey, policyName)
	if err != nil {
		return Status{}, fmt.Errorf("read Antiware policy: %w", err)
	}
	return Status{Configured: v.Exists, Enabled: v.Exists && v.DWORD && v.Data == "1"}, nil
}

func (m *Manager) run(ctx context.Context, profile string, planned []actions.Action) (string, error) {
	if m.executor == nil || m.snapshots == nil {
		return "", fmt.Errorf("Antiware manager is not configured")
	}
	return engine.RunActions(ctx, profile, planned, m.executor, m.snapshots)
}

func enabledActions() []actions.Action {
	return []actions.Action{{
		Type: actions.RegistrySetDWORD, Risk: actions.RiskLow,
		Key: policyKey, Name: policyName, Value: "1",
	}}
}

func disabledActions() []actions.Action {
	return []actions.Action{{
		Type: actions.RegistryDelete, Risk: actions.RiskLow,
		Key: policyKey, Name: policyName,
	}}
}
