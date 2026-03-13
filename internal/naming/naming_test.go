package naming

import (
	"testing"

	"tj/internal/tmux"
)

func TestIsGeneric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Generic shells
		{"bash is generic", "bash", true},
		{"zsh is generic", "zsh", true},
		{"fish is generic", "fish", true},
		{"sh is generic", "sh", true},

		// Case insensitive
		{"BASH is generic", "BASH", true},
		{"Zsh is generic", "Zsh", true},

		// AI tools
		{"claude is generic", "claude", true},
		{"opencode is generic", "opencode", true},
		{"aider is generic", "aider", true},

		// Default names and numbers
		{"0 is generic", "0", true},
		{"5 is generic", "5", true},
		{"main is generic", "main", true},
		{"default is generic", "default", true},

		// Empty is generic
		{"empty is generic", "", true},

		// Non-generic names
		{"nvim is not generic", "nvim", false},
		{"vim is not generic", "vim", false},
		{"cargo is not generic", "cargo", false},
		{"npm is not generic", "npm", false},
		{"python is not generic", "python", false},
		{"go is not generic", "go", false},
		{"editor is not generic", "editor", false},
		{"dev is not generic", "dev", false},
		{"api is not generic", "api", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsGeneric(tt.input)
			if got != tt.expected {
				t.Errorf("IsGeneric(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGuessName_Priority(t *testing.T) {
	tests := []struct {
		name         string
		paneInfo     tmux.PaneInfo
		expectedName string
		expectedGen  bool
	}{
		{
			name: "distinctive session is preferred",
			paneInfo: tmux.PaneInfo{
				Command:    "nvim",
				WindowName: "editor",
				Session:    "dev",
			},
			expectedName: "dev",
			expectedGen:  false,
		},
		{
			name: "window name used when session is generic",
			paneInfo: tmux.PaneInfo{
				Command:    "nvim",
				WindowName: "editor",
				Session:    "main",
			},
			expectedName: "editor",
			expectedGen:  false,
		},
		{
			name: "command used when session and window are generic",
			paneInfo: tmux.PaneInfo{
				Command:    "nvim",
				WindowName: "0",
				Session:    "main",
			},
			expectedName: "nvim",
			expectedGen:  false,
		},
		{
			name: "all generic returns session with generic flag",
			paneInfo: tmux.PaneInfo{
				Command:    "bash",
				WindowName: "0",
				Session:    "main",
			},
			expectedName: "main",
			expectedGen:  true,
		},
		{
			name: "empty session falls through to window",
			paneInfo: tmux.PaneInfo{
				Command:    "bash",
				WindowName: "editor",
				Session:    "",
			},
			expectedName: "editor",
			expectedGen:  false,
		},
		{
			name: "all empty returns empty with generic flag",
			paneInfo: tmux.PaneInfo{
				Command:    "",
				WindowName: "",
				Session:    "",
			},
			expectedName: "",
			expectedGen:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, isGeneric := GuessName(tt.paneInfo)
			if name != tt.expectedName {
				t.Errorf("GuessName() name = %q, want %q", name, tt.expectedName)
			}
			if isGeneric != tt.expectedGen {
				t.Errorf("GuessName() isGeneric = %v, want %v", isGeneric, tt.expectedGen)
			}
		})
	}
}
