package main

import (
	"os"
	"path/filepath"
	"testing"

	"tj/internal/config"
)

func TestParseFlagsSound(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--sound"})
	if overrides.Sound == nil || !*overrides.Sound {
		t.Error("expected --sound to set Sound=true")
	}

	_, _, _, _, overrides = parseFlags([]string{"--no-sound"})
	if overrides.Sound == nil || *overrides.Sound {
		t.Error("expected --no-sound to set Sound=false")
	}
}

func TestParseFlagsNotify(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--notify"})
	if overrides.Notify == nil || !*overrides.Notify {
		t.Error("expected --notify to set Notify=true")
	}

	_, _, _, _, overrides = parseFlags([]string{"--no-notify"})
	if overrides.Notify == nil || *overrides.Notify {
		t.Error("expected --no-notify to set Notify=false")
	}
}

func TestParseFlagsSortActivity(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--sort-activity"})
	if overrides.SortActivity == nil || !*overrides.SortActivity {
		t.Error("expected --sort-activity to set SortActivity=true")
	}

	_, _, _, _, overrides = parseFlags([]string{"--sort-watchlist"})
	if overrides.SortActivity == nil || *overrides.SortActivity {
		t.Error("expected --sort-watchlist to set SortActivity=false")
	}
}

func TestParseFlagsColumns(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--columns"})
	if overrides.Columns == nil || !*overrides.Columns {
		t.Error("expected --columns to set Columns=true")
	}
}

func TestParseFlagsRecencyColor(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--recency-color"})
	if overrides.RecencyColor == nil || !*overrides.RecencyColor {
		t.Error("expected --recency-color to set RecencyColor=true")
	}

	_, _, _, _, overrides = parseFlags([]string{"--no-recency-color"})
	if overrides.RecencyColor == nil || *overrides.RecencyColor {
		t.Error("expected --no-recency-color to set RecencyColor=false")
	}
}

func TestParseFlagsPicker(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--picker"})
	if overrides.PickerMode == nil || !*overrides.PickerMode {
		t.Error("expected --picker to set PickerMode=true")
	}
}

func TestParseFlagsMultiple(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--columns", "--sort-activity", "--sound", "--picker"})
	if overrides.Columns == nil || !*overrides.Columns {
		t.Error("expected Columns=true")
	}
	if overrides.SortActivity == nil || !*overrides.SortActivity {
		t.Error("expected SortActivity=true")
	}
	if overrides.Sound == nil || !*overrides.Sound {
		t.Error("expected Sound=true")
	}
	if overrides.PickerMode == nil || !*overrides.PickerMode {
		t.Error("expected PickerMode=true")
	}
}

func TestParseFlagsUnspecified(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{})
	if overrides.Sound != nil {
		t.Error("expected Sound to be nil when not specified")
	}
	if overrides.Notify != nil {
		t.Error("expected Notify to be nil when not specified")
	}
	if overrides.Columns != nil {
		t.Error("expected Columns to be nil when not specified")
	}
	if overrides.PickerMode != nil {
		t.Error("expected PickerMode to be nil when not specified")
	}
}

func TestParseFlagsPreservesExisting(t *testing.T) {
	remaining, configPath, watchlistPath, help, _ := parseFlags([]string{"-c", "/tmp/cfg.yaml", "--sound", "-w", "/tmp/wl.json"})
	if configPath != "/tmp/cfg.yaml" {
		t.Errorf("expected config path, got %q", configPath)
	}
	if watchlistPath != "/tmp/wl.json" {
		t.Errorf("expected watchlist path, got %q", watchlistPath)
	}
	if help {
		t.Error("expected help=false")
	}
	if len(remaining) != 0 {
		t.Errorf("expected no remaining args, got %v", remaining)
	}
}

func TestApplyOverridesNil(t *testing.T) {
	cfg := config.Default()
	origSound := cfg.Alerts.SoundOnReady
	origSort := cfg.Display.SortByActivity

	applyOverrides(cfg, CLIOverrides{}) // all nil

	if cfg.Alerts.SoundOnReady != origSound {
		t.Error("expected SoundOnReady unchanged")
	}
	if cfg.Display.SortByActivity != origSort {
		t.Error("expected SortByActivity unchanged")
	}
}

func TestApplyOverridesSet(t *testing.T) {
	cfg := config.Default()

	applyOverrides(cfg, CLIOverrides{
		Sound:        boolPtr(true),
		Notify:       boolPtr(true),
		SortActivity: boolPtr(true),
		Columns:      boolPtr(true),
		RecencyColor: boolPtr(false),
		PickerMode:   boolPtr(true),
	})

	if !cfg.Alerts.SoundOnReady {
		t.Error("expected SoundOnReady=true")
	}
	if !cfg.Alerts.NotifyOnReady {
		t.Error("expected NotifyOnReady=true")
	}
	if !cfg.Display.SortByActivity {
		t.Error("expected SortByActivity=true")
	}
	if cfg.Display.LayoutMode != "columns" {
		t.Errorf("expected LayoutMode=columns, got %q", cfg.Display.LayoutMode)
	}
	if cfg.Display.RecencyColor {
		t.Error("expected RecencyColor=false")
	}
	if !cfg.Display.PickerMode {
		t.Error("expected PickerMode=true")
	}
}

func TestConfigLayoutModeYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(configDir, 0755)

	configContent := `
display:
  layout_mode: columns
  picker_mode: true
`
	os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte(configContent), 0644)

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg := config.Load()

	if cfg.Display.LayoutMode != "columns" {
		t.Errorf("expected LayoutMode=columns, got %q", cfg.Display.LayoutMode)
	}
	if !cfg.Display.PickerMode {
		t.Error("expected PickerMode=true")
	}
}

func TestParseFlagsPreview(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--preview"})
	if overrides.Preview == nil || !*overrides.Preview {
		t.Error("expected --preview to set Preview=true")
	}

	_, _, _, _, overrides = parseFlags([]string{"--no-preview"})
	if overrides.Preview == nil || *overrides.Preview {
		t.Error("expected --no-preview to set Preview=false")
	}
}

func TestApplyOverridesPreview(t *testing.T) {
	cfg := config.Default()
	if !cfg.Display.ShowPreview {
		t.Fatal("expected default ShowPreview=true")
	}

	applyOverrides(cfg, CLIOverrides{Preview: boolPtr(false)})
	if cfg.Display.ShowPreview {
		t.Error("expected ShowPreview=false after --no-preview override")
	}

	applyOverrides(cfg, CLIOverrides{Preview: boolPtr(true)})
	if !cfg.Display.ShowPreview {
		t.Error("expected ShowPreview=true after --preview override")
	}

	// Nil should not change value
	cfg.Display.ShowPreview = false
	applyOverrides(cfg, CLIOverrides{})
	if cfg.Display.ShowPreview {
		t.Error("expected ShowPreview unchanged when override is nil")
	}
}

func TestConfigShowPreviewDefault(t *testing.T) {
	cfg := config.Default()
	if !cfg.Display.ShowPreview {
		t.Error("expected default ShowPreview=true")
	}
}

func TestConfigShowPreviewYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(configDir, 0755)

	configContent := `
display:
  show_preview: false
`
	os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte(configContent), 0644)

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg := config.Load()

	if cfg.Display.ShowPreview {
		t.Error("expected ShowPreview=false from config")
	}
}

func TestParseFlagsScan(t *testing.T) {
	_, _, _, _, overrides := parseFlags([]string{"--scan"})
	if overrides.Scan == nil || !*overrides.Scan {
		t.Error("expected --scan to set Scan=true")
	}
}

func TestApplyOverridesScan(t *testing.T) {
	cfg := config.Default()
	if cfg.Display.ScanOnStart {
		t.Fatal("expected default ScanOnStart=false")
	}

	applyOverrides(cfg, CLIOverrides{Scan: boolPtr(true)})
	if !cfg.Display.ScanOnStart {
		t.Error("expected ScanOnStart=true after --scan override")
	}

	// Nil should not change value
	cfg.Display.ScanOnStart = true
	applyOverrides(cfg, CLIOverrides{})
	if !cfg.Display.ScanOnStart {
		t.Error("expected ScanOnStart unchanged when override is nil")
	}
}

func TestConfigScanOnStartYAML(t *testing.T) {
	cfg := config.Default()
	if cfg.Display.ScanOnStart {
		t.Error("expected default ScanOnStart=false")
	}

	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(configDir, 0755)

	configContent := `
display:
  scan_on_start: true
`
	os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte(configContent), 0644)

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg = config.Load()
	if !cfg.Display.ScanOnStart {
		t.Error("expected ScanOnStart=true from config")
	}
}
