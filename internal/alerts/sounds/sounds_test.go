package sounds

import (
	"testing"
)

func TestValidSounds(t *testing.T) {
	sounds := ValidSounds()
	expected := []string{SoundChime, SoundBell, SoundPing, SoundPop, SoundDing}

	if len(sounds) != len(expected) {
		t.Errorf("ValidSounds() returned %d sounds, want %d", len(sounds), len(expected))
	}

	for i, s := range expected {
		if sounds[i] != s {
			t.Errorf("ValidSounds()[%d] = %q, want %q", i, sounds[i], s)
		}
	}
}

func TestIsValidSound(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{SoundChime, true},
		{SoundBell, true},
		{SoundPing, true},
		{SoundPop, true},
		{SoundDing, true},
		{"invalid", false},
		{"", false},
		{"CHIME", false}, // case sensitive
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsValidSound(tt.input)
			if got != tt.want {
				t.Errorf("IsValidSound(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNextSound(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{SoundChime, SoundBell},
		{SoundBell, SoundPing},
		{SoundPing, SoundPop},
		{SoundPop, SoundDing},
		{SoundDing, SoundChime}, // wraps around
		{"invalid", DefaultSound},
		{"", DefaultSound},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := NextSound(tt.input)
			if got != tt.want {
				t.Errorf("NextSound(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDefaultSound(t *testing.T) {
	if DefaultSound != SoundChime {
		t.Errorf("DefaultSound = %q, want %q", DefaultSound, SoundChime)
	}
}

func TestSoundConstants(t *testing.T) {
	// Ensure constants match expected string values
	tests := []struct {
		constant string
		value    string
	}{
		{SoundChime, "chime"},
		{SoundBell, "bell"},
		{SoundPing, "ping"},
		{SoundPop, "pop"},
		{SoundDing, "ding"},
	}

	for _, tt := range tests {
		if tt.constant != tt.value {
			t.Errorf("Sound constant %q != %q", tt.constant, tt.value)
		}
	}
}

func TestGetSoundInvalidFallback(t *testing.T) {
	// GetSound should fall back to default sound for invalid input
	// This tests that an invalid sound type doesn't cause an error
	// (it uses the embedded files, so we can test this)
	_, _, err := GetSound("invalid")
	if err != nil {
		t.Errorf("GetSound(\"invalid\") returned error: %v", err)
	}
}

func TestGetSoundAllValid(t *testing.T) {
	// Test that all valid sounds can be loaded from embedded files
	for _, soundType := range ValidSounds() {
		t.Run(soundType, func(t *testing.T) {
			streamer, format, err := GetSound(soundType)
			if err != nil {
				t.Errorf("GetSound(%q) returned error: %v", soundType, err)
				return
			}
			if streamer == nil {
				t.Errorf("GetSound(%q) returned nil streamer", soundType)
				return
			}
			if format.SampleRate == 0 {
				t.Errorf("GetSound(%q) returned zero sample rate", soundType)
			}
			streamer.Close()
		})
	}
}
