package config

import "os"

// Config is the configuration of this app
type Config struct {
	HerokuPort string
	HerokuURL  string
	TgToken    string
	MongoDBURI string
}

// Get the configuration from environment variable
func Get() *Config {
	return &Config{
		HerokuPort: os.Getenv("PORT"),
		HerokuURL:  os.Getenv("HEROKU_URL"),
		TgToken:    os.Getenv("TG_TOKEN"),
		MongoDBURI: os.Getenv("MONGODB_URI"),
	}
}
