package plugin

// GlobalCfg configures the plugin service.
type GlobalCfg struct {
	Addr string `mapstructure:"addr"`
}

// Config is used to parse pipeline configuration
type Config struct{}

func (c *Config) InitDefaults() {
	// all default options

}

// InitDefaults for the plugin config
func (cfg *GlobalCfg) InitDefaults() {
	if cfg.Addr == "" {
		cfg.Addr = "tcp://localhost:8088"
	}
}
