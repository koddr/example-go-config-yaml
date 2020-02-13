package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

var configPath string

// Config struct
type Config struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Write int `yaml:"write"`
			Read  int `yaml:"read"`
			Idle  int `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
}

// NewConfig returns a new Config
func NewConfig(configPath string) *Config {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		log.Fatal(err)
	}

	return config
}

func init() {
	// Looking for config flag
	flag.StringVar(&configPath, "config", "config.yml", "path to config file")
}

func main() {
	// Parse a command line arguments
	flag.Parse()

	// Create new config instance
	config := NewConfig(configPath)

	// Define server options
	server := &http.Server{
		Addr:         config.Server.Host + ":" + config.Server.Port,
		ReadTimeout:  time.Duration(config.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(config.Server.Timeout.Write) * time.Second,
		IdleTimeout:  time.Duration(config.Server.Timeout.Idle) * time.Second,
	}

	// Start server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
