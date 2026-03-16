package sounds

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/wav"
)

// Sound type constants
const (
	SoundChime = "chime"
	SoundBell  = "bell"
	SoundPing  = "ping"
	SoundPop   = "pop"
	SoundDing  = "ding"
)

// DefaultSound is the default sound type
const DefaultSound = SoundChime

// ValidSounds returns all valid sound type names
func ValidSounds() []string {
	return []string{SoundChime, SoundBell, SoundPing, SoundPop, SoundDing}
}

// IsValidSound returns true if the sound type is valid
func IsValidSound(soundType string) bool {
	for _, s := range ValidSounds() {
		if s == soundType {
			return true
		}
	}
	return false
}

// NextSound returns the next sound in the cycle
func NextSound(current string) string {
	sounds := ValidSounds()
	for i, s := range sounds {
		if s == current {
			return sounds[(i+1)%len(sounds)]
		}
	}
	return DefaultSound
}

var (
	speakerInitOnce sync.Once
	speakerInitErr  error
	speakerReady    bool
)

// initSpeaker initializes the audio speaker (lazy, called on first play)
func initSpeaker() error {
	speakerInitOnce.Do(func() {
		// Load a sound to get the sample rate
		f, err := SoundFiles.Open("chime.wav")
		if err != nil {
			speakerInitErr = fmt.Errorf("failed to open sound file: %w", err)
			return
		}
		defer f.Close()

		// Read into memory for seeking
		data, err := io.ReadAll(f)
		if err != nil {
			speakerInitErr = fmt.Errorf("failed to read sound file: %w", err)
			return
		}

		streamer, format, err := wav.Decode(io.NopCloser(bytes.NewReader(data)))
		if err != nil {
			speakerInitErr = fmt.Errorf("failed to decode WAV: %w", err)
			return
		}
		streamer.Close()

		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			speakerInitErr = fmt.Errorf("failed to init speaker: %w", err)
			return
		}
		speakerReady = true
	})
	return speakerInitErr
}

// GetSound returns a beep.StreamSeekCloser for the given sound type
func GetSound(soundType string) (beep.StreamSeekCloser, beep.Format, error) {
	if !IsValidSound(soundType) {
		soundType = DefaultSound
	}

	filename := soundType + ".wav"
	f, err := SoundFiles.Open(filename)
	if err != nil {
		return nil, beep.Format{}, fmt.Errorf("failed to open %s: %w", filename, err)
	}

	// Read into memory for seeking support
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		return nil, beep.Format{}, fmt.Errorf("failed to read %s: %w", filename, err)
	}

	streamer, format, err := wav.Decode(io.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return nil, beep.Format{}, fmt.Errorf("failed to decode %s: %w", filename, err)
	}

	return streamer, format, nil
}

// PlaySound plays the specified sound type. Falls back to terminal bell on error.
func PlaySound(soundType string) {
	if err := initSpeaker(); err != nil {
		log.Printf("Audio init failed, using terminal bell: %v", err)
		playTerminalBell()
		return
	}

	streamer, _, err := GetSound(soundType)
	if err != nil {
		log.Printf("Failed to get sound %s, using terminal bell: %v", soundType, err)
		playTerminalBell()
		return
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Wait for sound to finish (non-blocking in background)
	go func() {
		<-done
		streamer.Close()
	}()
}

// playTerminalBell outputs the terminal bell character
func playTerminalBell() {
	fmt.Print("\a")
}
