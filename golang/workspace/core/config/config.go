package config

// Config is struct which includes neccessary configs
type Config struct {
	HTTPServiceConfig
}

// HTTPServiceConfig is struct which includes http service config
type HTTPServiceConfig struct {
	HTTPServiceConfigBindIP  string `json:"http_service_bind_ip" default:"0.0.0.0" envconfig:"HTTP_SERVICE_CONFIG_BIND_IP"`
	HTTPServiceConfigPort    string `json:"http_service_port" default:"8000" envconfig:"HTTP_SERVICE_CONFIG_PORT"`
	HTTPServiceConfigURLBase string `json:"http_service_url_base" default:"/" envconfig:"HTTP_SERVICE_CONFIG_URL_BASE"`
}
