//go:build windows

package registry

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/sys/windows/registry"
)

type WindowsManager struct{}

func New() Manager { return WindowsManager{} }

func split(key string) (registry.Key, string, error) {
	parts := strings.SplitN(strings.ReplaceAll(key, "/", `\`), `\`, 2)
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("registry key %q must be HIVE\\path", key)
	}
	var hive registry.Key
	switch strings.ToUpper(parts[0]) {
	case "HKCU", "HKEY_CURRENT_USER":
		hive = registry.CURRENT_USER
	case "HKLM", "HKEY_LOCAL_MACHINE":
		hive = registry.LOCAL_MACHINE
	default:
		return 0, "", fmt.Errorf("unsupported registry hive %q", parts[0])
	}
	return hive, parts[1], nil
}

func (WindowsManager) Read(_ context.Context, key, name string) (Value, error) {
	hive, path, err := split(key)
	if err != nil {
		return Value{}, err
	}
	k, err := registry.OpenKey(hive, path, registry.QUERY_VALUE)
	if err != nil {
		return Value{}, err
	}
	defer k.Close()
	v, _, err := k.GetStringValue(name)
	if err == nil {
		return Value{Exists: true, Data: v}, nil
	}
	if err == registry.ErrNotExist {
		return Value{}, nil
	}
	dword, _, dwordErr := k.GetIntegerValue(name)
	if dwordErr == nil {
		return Value{Exists: true, Data: fmt.Sprintf("%d", dword), DWORD: true}, nil
	}
	return Value{}, err
}
func (WindowsManager) Values(_ context.Context, key string) (map[string]string, error) {
	hive, path, err := split(key)
	if err != nil {
		return nil, err
	}
	k, err := registry.OpenKey(hive, path, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer k.Close()
	names, err := k.ReadValueNames(-1)
	if err != nil {
		return nil, err
	}
	values := make(map[string]string, len(names))
	for _, name := range names {
		value, _, err := k.GetStringValue(name)
		if err == nil {
			values[name] = value
		}
	}
	return values, nil
}
func (WindowsManager) Set(_ context.Context, key, name, value string) error {
	hive, path, err := split(key)
	if err != nil {
		return err
	}
	k, _, err := registry.CreateKey(hive, path, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()
	return k.SetStringValue(name, value)
}
func (WindowsManager) SetDWORD(_ context.Context, key, name, value string) error {
	hive, path, err := split(key)
	if err != nil {
		return err
	}
	number, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid DWORD value %q: %w", value, err)
	}
	k, _, err := registry.CreateKey(hive, path, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()
	return k.SetDWordValue(name, uint32(number))
}
func (WindowsManager) Delete(_ context.Context, key, name string) error {
	hive, path, err := split(key)
	if err != nil {
		return err
	}
	k, err := registry.OpenKey(hive, path, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()
	err = k.DeleteValue(name)
	if err == registry.ErrNotExist {
		return nil
	}
	return err
}
