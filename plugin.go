package plugin

import (
	"github.com/aabri-g4m/rr-pubsub/pubsubjobs"
	"github.com/spiral/roadrunner/v2/plugins/config"
	"github.com/spiral/roadrunner/v2/plugins/logger"
)

// PluginName is the name for the plugin as found inside the container
const PluginName = "rr_pubsub"

// Plugin is the structure that will be initiated as per endure.Container interface
type Plugin struct {
	logger   logger.Logger
	cfg      config.Configurer
	localCfg pubsubjobs.Config
}

// Init initiates the plugin with any injected services implementing endure.Container, returning an error if the
// plugin fails to start, if the error is of type errors.Disabled then the plugin will not be active
func (p *Plugin) Init(cfg config.Configurer, logger logger.Logger) error {
	p.logger = logger
	p.cfg = cfg
	p.localCfg.InitDefaults()
	return nil
}

// Name returns endure.Named interface implementation
func (p *Plugin) Name() string {
	return PluginName
}

func (p *Plugin) Available() {}
