package profiles

type Profile struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Actions     []Action `yaml:"actions"`
}

type Action struct {
	Type   string   `yaml:"type"`
	Target []string `yaml:"target,omitempty"`
	Value  string   `yaml:"value,omitempty"`
}
