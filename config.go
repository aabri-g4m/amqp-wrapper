package plugin

// Config configures the plugin service.
type Config struct {
	Address string `mapstructure:"address"`
}

// InitDefaults for the plugin config
func (cfg *Config) InitDefaults() {
	if cfg.Address == "" {
		cfg.Address = "tcp://localhost:8088"
	}
}
