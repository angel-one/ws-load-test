package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type providers struct {
	providers map[string]*viper.Viper
	mu        sync.Mutex
}

var baseConfigPath string
var p *providers

// Init is used to initialize the configurations
func Init(path string) {
	baseConfigPath = path
	p = &providers{
		providers: make(map[string]*viper.Viper),
	}
}

// Get is used to get the instance to the config provider for the configuration name
func Get(name string) (*viper.Viper, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// see for an existing provider
	if provider, ok := p.providers[name]; ok {
		// provider already exists
		// use it
		return provider, nil
	}

	// try to get the provider
	provider := viper.New()
	provider.SetConfigName(name)
	provider.AddConfigPath(baseConfigPath)
	err := provider.ReadInConfig()
	if err != nil {
		// config not found or some other parsing errors
		return nil, fmt.Errorf("config %s error : %v", name, err.Error())
	}

	// add a watcher for this provider
	provider.WatchConfig()

	// successfully found config, store it for future use
	p.providers[name] = provider

	return provider, nil
}
