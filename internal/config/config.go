package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// AppPatterns defines detection patterns for a specific application.
type AppPatterns struct {
	PromptEndings  []string `yaml:"prompt_endings,omitempty"`
	WaitingStrings []string `yaml:"waiting_strings,omitempty"`
	BusyStrings    []string `yaml:"busy_strings,omitempty"`
}

// Detection holds activity detection configuration.
type Detection struct {
	IdleTimeout    time.Duration          `yaml:"idle_timeout"`
	PromptEndings  []string               `yaml:"prompt_endings"`
	WaitingStrings []string               `yaml:"waiting_strings"`
	BusyStrings    []string               `yaml:"busy_strings"`
	Apps           map[string]AppPatterns `yaml:"apps,omitempty"`
}

// Alerts holds alert notification configuration.
type Alerts struct {
	SoundOnReady  bool   `yaml:"sound_on_ready"`
	NotifyOnReady bool   `yaml:"notify_on_ready"`
	SoundType     string `yaml:"sound_type"`
}

// Config holds all application configuration.
type Config struct {
	Detection Detection `yaml:"detection"`
	Alerts    Alerts    `yaml:"alerts"`
}

// configFile is the YAML structure for parsing (with string duration).
type configFile struct {
	Detection struct {
		IdleTimeout    string                 `yaml:"idle_timeout"`
		PromptEndings  []string               `yaml:"prompt_endings"`
		WaitingStrings []string               `yaml:"waiting_strings"`
		BusyStrings    []string               `yaml:"busy_strings"`
		Apps           map[string]AppPatterns `yaml:"apps"`
	} `yaml:"detection"`
	Alerts Alerts `yaml:"alerts"`
}

// Default returns the default configuration with sensible defaults.
func Default() *Config {
	return &Config{
		Detection: Detection{
			IdleTimeout:    2 * time.Second,
			PromptEndings:  []string{},
			WaitingStrings: []string{},
			BusyStrings:    []string{},
			Apps: map[string]AppPatterns{
				"claude": {
					WaitingStrings: []string{"? for shortcuts"},
					BusyStrings:    []string{"Thinking", "Reasoning", "Undulating"},
				},
				"aider": {
					WaitingStrings: []string{"(Y)es/(N)o"},
				},
				"codex": {
					WaitingStrings: []string{"[Y/n]"},
				},
				"opencode": {
					WaitingStrings: []string{"Continue?"},
					BusyStrings:    []string{"Thinking", "Working"},
				},
			},
		},
		Alerts: Alerts{
			SoundOnReady:  false,
			NotifyOnReady: false,
			SoundType:     "chime",
		},
	}
}

// ConfigPath returns the path to the configuration file.
func ConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "teejay", "config.yaml"), nil
}

// Load reads the configuration from the specified path, or ~/.config/teejay/config.yaml if not provided.
// If the file doesn't exist or is malformed, returns defaults.
func Load(customPath ...string) *Config {
	var path string
	var err error

	if len(customPath) > 0 && customPath[0] != "" {
		path = customPath[0]
	} else {
		path, err = ConfigPath()
		if err != nil {
			return Default()
		}
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return Default()
	}
	if err != nil {
		log.Printf("Warning: failed to read config file: %v, using defaults", err)
		return Default()
	}

	var cf configFile
	if err := yaml.Unmarshal(data, &cf); err != nil {
		log.Printf("Warning: failed to parse config file: %v, using defaults", err)
		return Default()
	}

	cfg := Default()

	// Parse idle timeout if specified
	if cf.Detection.IdleTimeout != "" {
		if d, err := time.ParseDuration(cf.Detection.IdleTimeout); err == nil {
			cfg.Detection.IdleTimeout = d
		} else {
			log.Printf("Warning: invalid idle_timeout '%s': %v, using default", cf.Detection.IdleTimeout, err)
		}
	}

	// Override globals if specified in file
	if cf.Detection.PromptEndings != nil {
		cfg.Detection.PromptEndings = cf.Detection.PromptEndings
	}
	if cf.Detection.WaitingStrings != nil {
		cfg.Detection.WaitingStrings = cf.Detection.WaitingStrings
	}
	if cf.Detection.BusyStrings != nil {
		cfg.Detection.BusyStrings = cf.Detection.BusyStrings
	}

	// Merge app patterns: file overrides defaults for specified apps
	if cf.Detection.Apps != nil {
		for app, patterns := range cf.Detection.Apps {
			cfg.Detection.Apps[app] = patterns
		}
	}

	// Copy alerts settings from file, preserving defaults for unset values
	cfg.Alerts.SoundOnReady = cf.Alerts.SoundOnReady
	cfg.Alerts.NotifyOnReady = cf.Alerts.NotifyOnReady
	if cf.Alerts.SoundType != "" {
		cfg.Alerts.SoundType = cf.Alerts.SoundType
	}

	return cfg
}

// GetPatternsForApp returns the patterns to use for a given application.
// If the app has specific config, returns those patterns (replacing globals).
// Otherwise returns global patterns.
func (c *Config) GetPatternsForApp(appName string) (promptEndings, waitingStrings, busyStrings []string) {
	if appPatterns, ok := c.Detection.Apps[appName]; ok {
		return appPatterns.PromptEndings, appPatterns.WaitingStrings, appPatterns.BusyStrings
	}
	return c.Detection.PromptEndings, c.Detection.WaitingStrings, c.Detection.BusyStrings
}
