//go:build !windows

package registry

import (
	"context"
	"fmt"
)

type unavailable struct{}

func New() Manager { return unavailable{} }
func (unavailable) Read(context.Context, string, string) (Value, error) {
	return Value{}, fmt.Errorf("Windows registry is only available on Windows")
}
func (unavailable) Values(context.Context, string) (map[string]string, error) {
	return nil, fmt.Errorf("Windows registry is only available on Windows")
}
func (unavailable) Set(context.Context, string, string, string) error {
	return fmt.Errorf("Windows registry is only available on Windows")
}
func (unavailable) SetDWORD(context.Context, string, string, string) error {
	return fmt.Errorf("Windows registry is only available on Windows")
}
func (unavailable) Delete(context.Context, string, string) error {
	return fmt.Errorf("Windows registry is only available on Windows")
}
