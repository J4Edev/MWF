package profiles

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(name string) (*Profile, error) {
	if name != "default" {
		return nil, fmt.Errorf("only default profile exists in v0.2")
	}

	data, err := os.ReadFile("configs/profiles/default.yaml")
	if err != nil {
		return nil, err
	}

	var p Profile
	if err := yaml.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return &p, nil
}
