package snapshot

import (
	"MWF/internal/actions"
	"MWF/internal/system/power"
	"MWF/internal/system/registry"
	"MWF/internal/system/startup"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Manager struct{ Root string }
type Systems struct {
	Registry registry.Manager
	Power    power.Manager
	Startup  startup.Manager
	Windows  interface {
		Read(context.Context, string) (string, bool, error)
	}
}

func New(root string) *Manager {
	if root == "" {
		root = "snapshots"
	}
	return &Manager{Root: root}
}
func (m *Manager) Create(ctx context.Context, profile string, items []actions.Action, s Systems) (string, error) {
	id, err := newID()
	if err != nil {
		return "", err
	}
	records := make([]Record, 0, len(items))
	for _, a := range items {
		if a.Type == actions.StartupRemove {
			for _, name := range a.Target {
				one := a
				one.Target = []string{name}
				r, e := capture(ctx, one, s)
				if e != nil {
					return "", e
				}
				records = append(records, r)
			}
			continue
		}
		r, e := capture(ctx, a, s)
		if e != nil {
			return "", e
		}
		records = append(records, r)
	}
	dir := filepath.Join(m.Root, id)
	if err = os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	meta := Metadata{ID: id, Timestamp: time.Now().UTC(), Profile: profile, Actions: items}
	if err = write(filepath.Join(dir, "metadata.json"), meta); err != nil {
		return "", err
	}
	if err = write(filepath.Join(dir, "state.json"), State{Records: records}); err != nil {
		return "", err
	}
	return id, nil
}
func capture(ctx context.Context, a actions.Action, s Systems) (Record, error) {
	r := Record{Action: a}
	switch a.Type {
	case actions.RegistrySet, actions.RegistrySetDWORD, actions.RegistryDelete:
		v, e := s.Registry.Read(ctx, a.Key, a.Name)
		if e != nil {
			return r, e
		}
		r.Previous, r.Existed, r.Reversible = v.Data, v.Exists, true
		if v.DWORD {
			r.PreviousType = actions.RegistrySetDWORD
		}
	case actions.PowerPlan:
		v, e := s.Power.CurrentPlan(ctx)
		if e != nil {
			return r, e
		}
		r.Previous, r.Existed, r.Reversible = v, true, true
	case actions.StartupRemove:
		v, e := s.Startup.Lookup(ctx, a.Target[0])
		if e != nil {
			return r, e
		}
		r.Previous, r.Existed, r.Reversible = v.Data, v.Exists, true
	case actions.WindowsSetting:
		v, known, e := s.Windows.Read(ctx, a.Value)
		if e != nil {
			return r, e
		}
		r.Previous, r.Existed, r.Reversible = v, known, known
	}
	return r, nil
}
func (m *Manager) Load(id string) (Metadata, State, error) {
	var meta Metadata
	var state State
	dir := filepath.Join(m.Root, id)
	if err := read(filepath.Join(dir, "metadata.json"), &meta); err != nil {
		return meta, state, err
	}
	if err := read(filepath.Join(dir, "state.json"), &state); err != nil {
		return meta, state, err
	}
	return meta, state, nil
}
func (m *Manager) Reverse(id string) (string, []actions.Action, error) {
	meta, state, err := m.Load(id)
	if err != nil {
		return "", nil, err
	}
	out := make([]actions.Action, 0, len(state.Records))
	for i := len(state.Records) - 1; i >= 0; i-- {
		r := state.Records[i]
		if !r.Reversible {
			continue
		}
		switch r.Action.Type {
		case actions.RegistrySet, actions.RegistrySetDWORD, actions.RegistryDelete:
			a := r.Action
			if r.Existed {
				a.Type = actions.RegistrySet
				if r.PreviousType != "" {
					a.Type = r.PreviousType
				}
				a.Value = r.Previous
			} else {
				a.Type = actions.RegistryDelete
			}
			out = append(out, a)
		case actions.PowerPlan:
			out = append(out, actions.Action{Type: actions.PowerPlan, Value: r.Previous})
		case actions.StartupRemove:
			if r.Existed {
				out = append(out, actions.Action{Type: actions.RegistrySet, Key: startup.RunKey(), Name: r.Action.Target[0], Value: r.Previous})
			}
		}
	}
	return meta.Profile, out, nil
}
func write(path string, v any) error {
	b, e := json.MarshalIndent(v, "", "  ")
	if e != nil {
		return e
	}
	return os.WriteFile(path, b, 0644)
}
func read(path string, v any) error {
	b, e := os.ReadFile(path)
	if e != nil {
		return e
	}
	return json.Unmarshal(b, v)
}
func newID() (string, error) {
	b := make([]byte, 8)
	if _, e := rand.Read(b); e != nil {
		return "", e
	}
	return time.Now().UTC().Format("20060102T150405Z") + "-" + hex.EncodeToString(b), nil
}
