package main

// AppConfig holds all configuration settings for the application.
// We use 'json' struct tags to easily unmarshal settings from a JSON file.
type AppConfig struct {
	Port           int    `json:"port"`
	DatabaseURL    string `json:"database_url"`
	MaxConnections int    `json:"max_connections"`
	APITimeoutMs   int    `json:"api_timeout_ms"`

	// Internak fields (unexported) for application state
	isInitialized bool
}

// NewConfig is a constructor function for creating the configuration.
func NewConfig() *AppConfig {
	// In a real app, this function would load from a file/environment vars.
	return &AppConfig{
		Port:           8080,
		DatabaseURL:    "postgres://user:pass@localhost:5432/myapp",
		MaxConnections: 10,
		APITimeoutMs:   5000,
		isInitialized:  true,
	}
}

// func main() {
// 	config := NewConfig()
// 	fmt.Println(config)
// }
