package profiles

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

func LoadProfile(name string) (*Profile, error) {
	if name != "default" {
		return nil, fmt.Errorf("profile not found: %s", name)
	}

	data, err := os.ReadFile("internal/profiles/default.yaml")
	if err != nil {
		return nil, err
	}

	var p Profile
	err = yaml.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
