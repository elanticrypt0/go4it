package go4it

type AppConfig struct {
	App_name            string `json:"app_App_name"`
	App_version         string `json:"app_App_version"`
	App_author          string `json:"app_App_author"`
	App_contact         string `json:"app_App_contact"`
	App_repo            string `json:"app_App_repo"`
	App_server_protocol string `json:"app_App_server_protocol"`
	App_server_host     string `json:"app_App_server_host"`
	App_server_port     uint16 `json:"app_App_server_port"`
	App_url             string `json:"app_App_url"`
	App_setup_enabled   bool   `json:"app_setup_enabled"`
	App_debug_mode      bool   `json:"app_debug_mode"`
	App_CORS_origins    string `json:"app_App_CORS_origins"`
	App_CORS_headers    string `json:"app_App_CORS_headers"`
	DB                  map[string]DatabaseConfig
}
