package go4it

type AppConfig struct {
	App_name            string `json:"app_name" toml:"app_name"`
	App_version         string `json:"app_version" toml:"app_version"`
	App_author          string `json:"app_author" toml:"app_author"`
	App_contact         string `json:"app_contact" toml:"app_contact"`
	App_repo            string `json:"app_repo" toml:"app_repo"`
	App_server_protocol string `json:"app_server_protocol" toml:"app_server_protocol"`
	App_server_host     string `json:"app_server_host" toml:"app_server_host"`
	App_server_port     uint16 `json:"app_server_port" toml:"app_server_port"`
	App_url             string `json:"app_url" toml:"app_url"`
	App_setup_enabled   bool   `json:"app_setup_enabled" toml:"app_setup_enabled"`
	App_debug_mode      bool   `json:"app_debug_mode" toml:"app_debug_mode"`
	App_CORS_origins    string `json:"app_CORS_origins" toml:"app_CORS_origins"`
	App_CORS_headers    string `json:"app_CORS_headers" toml:"app_CORS_headers"`
	DB                  map[string]DatabaseConfig
}
