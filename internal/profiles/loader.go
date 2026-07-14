package profiles

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Load(name string) (*Profile, error) {
	if name == "" { return nil, fmt.Errorf("profile name is required") }
	data, err := os.ReadFile(filepath.Join("configs", "profiles", name+".yaml"))
	if err != nil {
		return nil, err
	}

	var p Profile
	if err := yaml.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	if p.Name == "" { p.Name = name }

	return &p, nil
}
