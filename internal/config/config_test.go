package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	cfg := Default()

	// Check idle timeout default
	if cfg.Detection.IdleTimeout != 2*time.Second {
		t.Errorf("expected idle timeout 2s, got %v", cfg.Detection.IdleTimeout)
	}

	// Check global patterns are empty
	if len(cfg.Detection.PromptEndings) != 0 {
		t.Errorf("expected empty prompt endings, got %v", cfg.Detection.PromptEndings)
	}
	if len(cfg.Detection.WaitingStrings) != 0 {
		t.Errorf("expected empty waiting strings, got %v", cfg.Detection.WaitingStrings)
	}

	// Check claude app defaults
	claudePatterns, ok := cfg.Detection.Apps["claude"]
	if !ok {
		t.Fatal("expected claude app config")
	}
	if len(claudePatterns.WaitingStrings) == 0 {
		t.Error("expected claude waiting strings")
	}
	found := false
	for _, s := range claudePatterns.WaitingStrings {
		if s == "? for shortcuts" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected '? for shortcuts' in claude waiting strings")
	}

	// Check aider app defaults
	aiderPatterns, ok := cfg.Detection.Apps["aider"]
	if !ok {
		t.Fatal("expected aider app config")
	}
	if len(aiderPatterns.WaitingStrings) == 0 {
		t.Error("expected aider waiting strings")
	}
}

func TestLoadMissingFile(t *testing.T) {
	// Temporarily change home to a temp dir without config
	origHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg := Load()

	// Should return defaults
	if cfg.Detection.IdleTimeout != 2*time.Second {
		t.Errorf("expected default idle timeout, got %v", cfg.Detection.IdleTimeout)
	}
}

func TestLoadValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(configDir, 0755)

	configContent := `
detection:
  idle_timeout: 5s
  prompt_endings:
    - "$"
    - ">"
  waiting_strings:
    - "custom prompt"
  apps:
    myapp:
      waiting_strings:
        - "myapp ready"
`
	os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte(configContent), 0644)

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg := Load()

	if cfg.Detection.IdleTimeout != 5*time.Second {
		t.Errorf("expected idle timeout 5s, got %v", cfg.Detection.IdleTimeout)
	}

	if len(cfg.Detection.PromptEndings) != 2 {
		t.Errorf("expected 2 prompt endings, got %d", len(cfg.Detection.PromptEndings))
	}

	if len(cfg.Detection.WaitingStrings) != 1 || cfg.Detection.WaitingStrings[0] != "custom prompt" {
		t.Errorf("expected custom waiting string, got %v", cfg.Detection.WaitingStrings)
	}

	myappPatterns, ok := cfg.Detection.Apps["myapp"]
	if !ok {
		t.Fatal("expected myapp config")
	}
	if len(myappPatterns.WaitingStrings) != 1 || myappPatterns.WaitingStrings[0] != "myapp ready" {
		t.Errorf("expected myapp waiting strings, got %v", myappPatterns.WaitingStrings)
	}

	// Default apps should still be present
	if _, ok := cfg.Detection.Apps["claude"]; !ok {
		t.Error("expected claude defaults to be preserved")
	}
}

func TestLoadMalformedFile(t *testing.T) {
	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(configDir, 0755)

	// Invalid YAML
	configContent := `
detection:
  idle_timeout: [invalid
`
	os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte(configContent), 0644)

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg := Load()

	// Should return defaults on parse error
	if cfg.Detection.IdleTimeout != 2*time.Second {
		t.Errorf("expected default idle timeout on malformed file, got %v", cfg.Detection.IdleTimeout)
	}
}

func TestGetPatternsForApp(t *testing.T) {
	cfg := Default()

	// App with config should return app-specific patterns (replace globals)
	endings, strings := cfg.GetPatternsForApp("claude")
	if len(endings) != 0 {
		t.Errorf("expected no prompt endings for claude, got %v", endings)
	}
	if len(strings) == 0 {
		t.Error("expected waiting strings for claude")
	}

	// App without config should return global patterns
	cfg.Detection.PromptEndings = []string{"$", ">"}
	cfg.Detection.WaitingStrings = []string{"global prompt"}

	endings, strings = cfg.GetPatternsForApp("unknown-app")
	if len(endings) != 2 {
		t.Errorf("expected 2 global prompt endings, got %v", endings)
	}
	if len(strings) != 1 || strings[0] != "global prompt" {
		t.Errorf("expected global waiting string, got %v", strings)
	}
}

func TestDefaultAlerts(t *testing.T) {
	cfg := Default()

	// Check alerts default to false
	if cfg.Alerts.SoundOnReady {
		t.Error("expected default SoundOnReady to be false")
	}
	if cfg.Alerts.NotifyOnReady {
		t.Error("expected default NotifyOnReady to be false")
	}
}

func TestLoadAlertsSection(t *testing.T) {
	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(configDir, 0755)

	configContent := `
alerts:
  sound_on_ready: true
  notify_on_ready: true
`
	os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte(configContent), 0644)

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg := Load()

	if !cfg.Alerts.SoundOnReady {
		t.Error("expected SoundOnReady to be true")
	}
	if !cfg.Alerts.NotifyOnReady {
		t.Error("expected NotifyOnReady to be true")
	}
}
