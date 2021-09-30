package plugin

import "log"

// GlobalCfg configures the plugin service.
type GlobalCfg struct {
	Addr string `mapstructure:"addr"`
}

// InitDefaults for the plugin config
func (cfg *GlobalCfg) InitDefaults() {
	if cfg.Addr == "" {
		cfg.Addr = "tcp://localhost:8088"
	}
	log.Println(cfg.Addr)
}
