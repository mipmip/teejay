//go:build ignore

// This file generates WAV sound files for notification alerts.
// Run with: go run gen_sounds.go
package main

import (
	"encoding/binary"
	"math"
	"os"
)

const sampleRate = 44100

func main() {
	// Generate 5 distinct notification sounds
	generateWAV("chime.wav", generateChime())
	generateWAV("bell.wav", generateBell())
	generateWAV("ping.wav", generatePing())
	generateWAV("pop.wav", generatePop())
	generateWAV("ding.wav", generateDing())
}

// generateChime creates a pleasant two-tone chime
func generateChime() []float64 {
	duration := 0.3
	samples := int(float64(sampleRate) * duration)
	data := make([]float64, samples)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		env := math.Exp(-t * 8) // decay envelope

		// Two harmonious frequencies (C5 and E5)
		freq1 := 523.25 // C5
		freq2 := 659.25 // E5

		// Second tone starts slightly later
		tone1 := math.Sin(2 * math.Pi * freq1 * t)
		tone2 := 0.0
		if t > 0.05 {
			tone2 = math.Sin(2 * math.Pi * freq2 * (t - 0.05))
		}

		data[i] = env * (tone1*0.6 + tone2*0.4) * 0.7
	}
	return data
}

// generateBell creates a bell-like sound with harmonics
func generateBell() []float64 {
	duration := 0.4
	samples := int(float64(sampleRate) * duration)
	data := make([]float64, samples)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		env := math.Exp(-t * 5)

		// Bell harmonics
		freq := 880.0 // A5
		tone := math.Sin(2*math.Pi*freq*t) * 0.5
		tone += math.Sin(2*math.Pi*freq*2.0*t) * 0.25
		tone += math.Sin(2*math.Pi*freq*3.0*t) * 0.125

		data[i] = env * tone * 0.6
	}
	return data
}

// generatePing creates a short, high ping sound
func generatePing() []float64 {
	duration := 0.15
	samples := int(float64(sampleRate) * duration)
	data := make([]float64, samples)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		env := math.Exp(-t * 20)

		freq := 1200.0
		tone := math.Sin(2 * math.Pi * freq * t)

		data[i] = env * tone * 0.7
	}
	return data
}

// generatePop creates a short pop/bubble sound
func generatePop() []float64 {
	duration := 0.1
	samples := int(float64(sampleRate) * duration)
	data := make([]float64, samples)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		env := math.Exp(-t * 30)

		// Frequency drops quickly (pop effect)
		freq := 600.0*math.Exp(-t*10) + 200.0
		tone := math.Sin(2 * math.Pi * freq * t)

		data[i] = env * tone * 0.8
	}
	return data
}

// generateDing creates a simple single-tone ding
func generateDing() []float64 {
	duration := 0.25
	samples := int(float64(sampleRate) * duration)
	data := make([]float64, samples)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		env := math.Exp(-t * 10)

		freq := 1046.50 // C6
		tone := math.Sin(2 * math.Pi * freq * t)

		data[i] = env * tone * 0.7
	}
	return data
}

func generateWAV(filename string, samples []float64) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numSamples := len(samples)
	dataSize := numSamples * 2 // 16-bit samples

	// WAV header
	f.Write([]byte("RIFF"))
	binary.Write(f, binary.LittleEndian, uint32(36+dataSize))
	f.Write([]byte("WAVE"))

	// fmt chunk
	f.Write([]byte("fmt "))
	binary.Write(f, binary.LittleEndian, uint32(16))         // chunk size
	binary.Write(f, binary.LittleEndian, uint16(1))          // PCM format
	binary.Write(f, binary.LittleEndian, uint16(1))          // mono
	binary.Write(f, binary.LittleEndian, uint32(sampleRate)) // sample rate
	binary.Write(f, binary.LittleEndian, uint32(sampleRate*2)) // byte rate
	binary.Write(f, binary.LittleEndian, uint16(2))          // block align
	binary.Write(f, binary.LittleEndian, uint16(16))         // bits per sample

	// data chunk
	f.Write([]byte("data"))
	binary.Write(f, binary.LittleEndian, uint32(dataSize))

	// Write samples
	for _, s := range samples {
		// Clamp and convert to 16-bit
		if s > 1.0 {
			s = 1.0
		}
		if s < -1.0 {
			s = -1.0
		}
		sample := int16(s * 32767)
		binary.Write(f, binary.LittleEndian, sample)
	}
}
