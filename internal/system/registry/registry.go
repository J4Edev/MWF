// Package registry is the sole owner of Windows registry access.
package registry

import (
	"context"
	"fmt"
)

type Value struct {
	Exists bool   `json:"exists"`
	Data   string `json:"data,omitempty"`
	DWORD  bool   `json:"dword,omitempty"`
}

type Manager interface {
	Read(context.Context, string, string) (Value, error)
	Values(context.Context, string) (map[string]string, error)
	Set(context.Context, string, string, string) error
	SetDWORD(context.Context, string, string, string) error
	Delete(context.Context, string, string) error
}

func ValidateKey(key string) error {
	if key == "" {
		return fmt.Errorf("registry key is required")
	}
	if len(key) < 5 {
		return fmt.Errorf("registry key %q must include a hive", key)
	}
	return nil
}
