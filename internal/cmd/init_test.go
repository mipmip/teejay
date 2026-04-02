package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestInitConfig_Defaults(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Simulate pressing Enter for all prompts (accept defaults)
	Stdin = strings.NewReader("\n\n\n\n")
	defer func() { Stdin = os.Stdin }()

	err := InitConfig()
	if err != nil {
		t.Fatalf("InitConfig() error = %v", err)
	}

	cfgPath := filepath.Join(tmpDir, ".config", "teejay", "config.yaml")
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		t.Fatalf("failed to read generated config: %v", err)
	}

	// Verify it's valid YAML
	var parsed map[string]interface{}
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("generated config is not valid YAML: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "sound_on_ready: false") {
		t.Error("expected sound_on_ready: false (default)")
	}
	if !strings.Contains(content, "notify_on_ready: false") {
		t.Error("expected notify_on_ready: false (default)")
	}
	if !strings.Contains(content, `layout_mode: "default"`) {
		t.Error("expected layout_mode: default")
	}
	if !strings.Contains(content, "sort_by_activity: false") {
		t.Error("expected sort_by_activity: false (default)")
	}
}

func TestInitConfig_CustomChoices(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// y for sound, y for notify, 2 for columns, 2 for activity sort
	Stdin = strings.NewReader("y\ny\n2\n2\n")
	defer func() { Stdin = os.Stdin }()

	err := InitConfig()
	if err != nil {
		t.Fatalf("InitConfig() error = %v", err)
	}

	cfgPath := filepath.Join(tmpDir, ".config", "teejay", "config.yaml")
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		t.Fatalf("failed to read generated config: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "sound_on_ready: true") {
		t.Error("expected sound_on_ready: true")
	}
	if !strings.Contains(content, "notify_on_ready: true") {
		t.Error("expected notify_on_ready: true")
	}
	if !strings.Contains(content, `layout_mode: "columns"`) {
		t.Error("expected layout_mode: columns")
	}
	if !strings.Contains(content, "sort_by_activity: true") {
		t.Error("expected sort_by_activity: true")
	}
}

func TestInitConfig_ExistingFileCancel(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Create existing config
	cfgDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(cfgDir, 0o755)
	cfgPath := filepath.Join(cfgDir, "config.yaml")
	os.WriteFile(cfgPath, []byte("original content"), 0o644)

	// Choose cancel (option 3)
	Stdin = strings.NewReader("3\n")
	defer func() { Stdin = os.Stdin }()

	err := InitConfig()
	if err != nil {
		t.Fatalf("InitConfig() error = %v", err)
	}

	// Original file should be unchanged
	data, _ := os.ReadFile(cfgPath)
	if string(data) != "original content" {
		t.Error("config file was modified despite choosing cancel")
	}
}

func TestInitConfig_ExistingFileBackup(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Create existing config
	cfgDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(cfgDir, 0o755)
	cfgPath := filepath.Join(cfgDir, "config.yaml")
	os.WriteFile(cfgPath, []byte("original content"), 0o644)

	// Choose backup (option 2), then accept all defaults
	Stdin = strings.NewReader("2\n\n\n\n\n")
	defer func() { Stdin = os.Stdin }()

	err := InitConfig()
	if err != nil {
		t.Fatalf("InitConfig() error = %v", err)
	}

	// Backup should exist
	bakPath := cfgPath + ".bak"
	bakData, err := os.ReadFile(bakPath)
	if err != nil {
		t.Fatalf("backup file not found: %v", err)
	}
	if string(bakData) != "original content" {
		t.Error("backup file does not contain original content")
	}

	// New config should be valid YAML
	data, _ := os.ReadFile(cfgPath)
	var parsed map[string]interface{}
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("new config is not valid YAML: %v", err)
	}
}

func TestInitConfig_ExistingFileOverwrite(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Create existing config
	cfgDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(cfgDir, 0o755)
	cfgPath := filepath.Join(cfgDir, "config.yaml")
	os.WriteFile(cfgPath, []byte("original content"), 0o644)

	// Choose overwrite (option 1), then accept all defaults
	Stdin = strings.NewReader("1\n\n\n\n\n")
	defer func() { Stdin = os.Stdin }()

	err := InitConfig()
	if err != nil {
		t.Fatalf("InitConfig() error = %v", err)
	}

	// File should be new content, not original
	data, _ := os.ReadFile(cfgPath)
	if string(data) == "original content" {
		t.Error("config was not overwritten")
	}

	// Should be valid YAML
	var parsed map[string]interface{}
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("overwritten config is not valid YAML: %v", err)
	}
}
