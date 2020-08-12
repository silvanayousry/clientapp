package configs

import (
	"time"

	"github.com/spf13/viper"
)

// ApplicationConfig contains the application configuration
type ApplicationConfig struct {
	Port        int
	ReadTimeout time.Duration
}

// appconfig is the default application configuration
var appconfig ApplicationConfig

// App returns the default application configuration
func App() ApplicationConfig {
	return appconfig
}

// LoadApp loads application configuration
func LoadApplication() {
	appconfig = ApplicationConfig{
		Port:        viper.GetInt("app.port"),
		ReadTimeout: viper.GetDuration("app.read_timeout") * time.Second,
	}
}
