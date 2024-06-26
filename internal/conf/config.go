// conf/config.go
package conf

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var configFiles embed.FS

type Settings struct {
	Debug bool // true to enable debug mode

	Main struct {
		Name      string // name of go-birdnet node, can be used to identify source of notes
		TimeAs24h bool   // true 24-hour time format, false 12-hour time format
		Log       LogConfig
	}

	BirdNET struct {
		Sensitivity             float64 // birdnet analysis sigmoid sensitivity
		Threshold               float64 // threshold for prediction confidence to report
		Overlap                 float64 // birdnet analysis overlap between chunks
		Longitude               float64 // longitude of recording location for prediction filtering
		Latitude                float64 // latitude of recording location for prediction filtering
		Threads                 int     // number of CPU threads to use for analysis
		Locale                  string  // language to use for labels
		LocationFilterThreshold float32 // threshold for prediction confidence to report
	}

	Input struct {
		Path      string // path to input file or directory
		Recursive bool   // true for recursive directory analysis
	}

	Realtime struct {
		Interval       int  // minimum interval between log messages in seconds
		ProcessingTime bool // true to report processing time for each prediction

		AudioExport struct {
			Enabled bool   // export audio clips containing indentified bird calls
			Path    string // path to audio clip export directory
			Type    string // audio file type, wav, mp3 or flac
		}

		Log struct {
			Enabled bool   // true to enable OBS chat log
			Path    string // path to OBS chat log
		}

		Birdweather struct {
			Enabled          bool    // true to enable birdweather uploads
			Debug            bool    // true to enable debug mode
			ID               string  // birdweather ID
			Threshold        float64 // threshold for prediction confidence for uploads
			LocationAccuracy float64 // accuracy of location in meters
		}

		PrivacyFilter struct {
			Enabled bool // true to enable privacy filter
		}

		DogBarkFilter struct {
			Enabled bool // true to enable dog bark filter
		}

		Retention struct {
			MinEvictionHours   int // minimum number of hours to keep audio clips
			MinClipsPerSpecies int // minimum number of clips per species to keep
		}

		RTSP struct {
			Url       string // RTSP stream URL
			Transport string // RTSP Transport Protocol
		}

		MQTT struct {
			Enabled  bool   // true to enable MQTT
			Broker   string // MQTT (tcp://host:port)
			Topic    string // MQTT topic
			Username string // MQTT username
			Password string // MQTT password
		}

		Telemetry struct {
			Enabled bool   // true to enable Prometheus compatible telemetry endpoint
			Listen  string // IP address and port to listen on
		}
	}

	WebServer struct {
		Enabled bool   // true to enable web server
		Port    string // port for web server
		AutoTLS bool   // true to enable auto TLS
		Log     LogConfig
	}

	Output struct {
		File struct {
			Enabled bool   // true to enable file output
			Path    string // directory to output results
			Type    string // table, csv
		}

		SQLite struct {
			Enabled bool   // true to enable sqlite output
			Path    string // path to sqlite database
		}

		MySQL struct {
			Enabled  bool   // true to enable mysql output
			Username string // username for mysql database
			Password string // password for mysql database
			Database string // database name for mysql database
			Host     string // host for mysql database
			Port     string // port for mysql database
		}
	}
}

// LogConfig defines the configuration for a log file
type LogConfig struct {
	Enabled     bool         // true to enable this log
	Path        string       // Path to the log file
	Rotation    RotationType // Type of log rotation
	MaxSize     int64        // Max size in bytes for RotationSize
	RotationDay time.Weekday // Day of the week for RotationWeekly
}

// RotationType defines different types of log rotations.
type RotationType string

const (
	RotationDaily  RotationType = "daily"
	RotationWeekly RotationType = "weekly"
	RotationSize   RotationType = "size"
)

// buildTime is the time when the binary was built.
var buildDate string

// Load reads the configuration file and environment variables into GlobalConfig.
func Load() (*Settings, error) {
	// Create a new settings struct
	settings := &Settings{}

	// Initialize viper and read config
	if err := initViper(); err != nil {
		return nil, fmt.Errorf("error initializing viper: %w", err)
	}

	// Unmarshal config into struct
	if err := viper.Unmarshal(settings); err != nil {
		return nil, fmt.Errorf("error unmarshaling config into struct: %w", err)
	}

	// Validate MQTT settings
	if settings.Realtime.MQTT.Enabled {
		if settings.Realtime.MQTT.Broker == "" {
			return nil, errors.New("MQTT broker URL is required when MQTT is enabled")
		}
	}
	return settings, nil
}

// initViper initializes viper with default values and reads the configuration file.
func initViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Get OS specific config paths
	configPaths, err := GetDefaultConfigPaths()
	if err != nil {
		return fmt.Errorf("error getting default config paths: %w", err)
	}

	// Assign config paths to Viper
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}

	// Set default values for each configuration parameter
	// function defined in defaults.go
	setDefaultConfig()

	// Read configuration file
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, create config with defaults
			return createDefaultConfig()
		}
		return fmt.Errorf("fatal error reading config file: %w", err)
	}

	// Print build date and config file used
	fmt.Printf("BirdNET-Go build date: %s, using config file: %s\n", buildDate, viper.ConfigFileUsed())

	return nil
}

// createDefaultConfig creates a default config file and writes it to the default config path
func createDefaultConfig() error {
	configPaths, err := GetDefaultConfigPaths() // Again, adjusted for error handling
	if err != nil {
		return fmt.Errorf("error getting default config paths: %w", err)
	}
	configPath := filepath.Join(configPaths[0], "config.yaml")
	defaultConfig := getDefaultConfig()

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("error creating directories for config file: %w", err)
	}

	if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}

	fmt.Println("Created default config file at:", configPath)
	return viper.ReadInConfig()
}

// getDefaultConfig reads the default configuration from the embedded config.yaml file.
func getDefaultConfig() string {
	data, err := fs.ReadFile(configFiles, "config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	return string(data)
}
