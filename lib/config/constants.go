package config

type ServiceConfig struct {
	Datacenter  string `yaml:"datacenter"`
	Node        string `yaml:"node"`
	Address     string `yaml:"address"`
	Port        int    `yaml:"port"`
	Interface   string `yaml"interface"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Instance    string `yaml:"instance"`
}

type HealthConfig struct {
	Command string `yaml:"command"`
	OkRegex string `yaml:"okregex"`
}

type MetricsConfig struct {
	Command string `yaml:"command"`
}

type Config struct {
	Service ServiceConfig `yaml:"service"`
	Health  HealthConfig  `yaml:"health"`
	Metrics MetricsConfig `yaml"metrics"`
}
