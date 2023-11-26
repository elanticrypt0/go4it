package go4it

type AppConfig struct {
	App_name            string
	App_version         string
	App_author          string
	App_contact         string
	App_repo            string
	App_server_protocol string
	App_server_host     string
	App_server_port     uint16
	App_url             string
	App_setup_enabled   bool
	App_debug_mode      bool
	App_CORS_origins    string
	App_CORS_headers    string
	DB                  map[string]DatabaseConfig
}
