package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Host 			string 		  `envconfig:"HOST" required:"true"`
	Port 			int			  `envconfig:"PORT" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

type DatabaseConfig struct {
	Host     string        `envconfig:"POSTGRES_HOST" required:"true"`
	Port     string        `envconfig:"POSTGRES_PORT" required:"true"`
	User     string        `envconfig:"POSTGRES_USER" required:"true"`
	Password string        `envconfig:"POSTGRES_PASSWORD" required:"true"`
	DB       string        `envconfig:"POSTGRES_DB" required:"true"`
	DBUrl    string        `envconfig:"POSTGRES_URL" required:"true"`
	Timeout   time.Duration `envconfig:"POSTGRES_TIMEOUT" required:"true"`
}

type LoggerConfig struct {
	Level  string `envconfig:"LEVEL" required:"true"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

type EnvironmentConfig struct {
	Debug bool `envconfig:"DEBUG" required:"false"`
}

type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Logger      LoggerConfig
	Environment EnvironmentConfig
}

func LoadConfig() (*Config, error) {
	var config Config

	// Try to load env file
	err := tryLoadEnvFile()
	if err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Load environment variables
	processors := map[string]any{
		"SERVER":      &config.Server,
		"DATABASE":    &config.Database,
		"LOGGER":      &config.Logger,
		"ENVIRONMENT": &config.Environment,
	}

	for prefix, cfg := range processors {
		if err := envconfig.Process(prefix, cfg); err != nil {
			return nil, fmt.Errorf("unable to get %s config: %w", prefix, err)
		}
	}
	
	return &config, nil
}


func tryLoadEnvFile() error {
	paths := []string{
		".env",                    
		"../../.env",             
		"../.env",               
	}
	
	if root := os.Getenv("PROJECT_ROOT"); root != "" {
		paths = append([]string{filepath.Join(root, ".env")}, paths...)
	}
	
	for _, path := range paths {
		if err := godotenv.Load(path); err == nil {
			log.Printf("Loaded .env from: %s", path)
			return nil
		}	
	}
	
	return fmt.Errorf("no .env file found in paths: %v", paths)
}