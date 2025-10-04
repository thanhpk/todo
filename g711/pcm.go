package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/ebitengine/oto/v3"
)

// This code generate sample PCM stream and play it
// Turn on your speaker to hear it
func playpcm() {
	const (
		sampleRate = 8000
		duration   = 2 * time.Second
		freq       = 440.0 // A4 note
	)

	// Generate 16-bit PCM samples
	numSamples := int(float64(sampleRate) * duration.Seconds())
	data := make([]byte, numSamples*2) // 2 bytes per sample (int16)

	for i := 0; i < numSamples; i++ {
		s := int16(3000 * math.Sin(2*math.Pi*freq*float64(i)/float64(sampleRate)))
		binary.LittleEndian.PutUint16(data[i*2:], uint16(s))
	}

	// Initialize Oto audio context
	opts := &oto.NewContextOptions{
		SampleRate:   sampleRate,
		ChannelCount: 1,
		Format:       oto.FormatSignedInt16LE,
	}
	ctx, ready, err := oto.NewContext(opts)
	if err != nil {
		panic(err)
	}
	<-ready // wait until ready

	player := ctx.NewPlayer(bytes.NewReader(data))
	player.Play()

	fmt.Println("Playing... Press Ctrl+C to exit.")
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	player.Close()
}

var musicNoteFreqM = map[string]float64{
	"C0": 16.35, "C#0": 17.32, "D0": 18.35, "D#0": 19.45, "E0": 20.60, "F0": 21.83, "F#0": 23.12, "G0": 24.50, "G#0": 25.96, "A0": 27.50, "A#0": 29.14, "B0": 30.87,
	"C1": 32.70, "C#1": 34.65, "D1": 36.71, "D#1": 38.89, "E1": 41.20, "F1": 43.65, "F#1": 46.25, "G1": 49.00, "G#1": 51.91, "A1": 55.00, "A#1": 58.27, "B1": 61.74,
	"C2": 65.41, "C#2": 69.30, "D2": 73.42, "D#2": 77.78, "E2": 82.41, "F2": 87.31, "F#2": 92.50, "G2": 98.00, "G#2": 103.83, "A2": 110.00, "A#2": 116.54, "B2": 123.47,
	"C3": 130.81, "C#3": 138.59, "D3": 146.83, "D#3": 155.56, "E3": 164.81, "F3": 174.61, "F#3": 185.00, "G3": 196.00, "G#3": 207.65, "A3": 220.00, "A#3": 233.08, "B3": 246.94,
	"C4": 261.63, "C#4": 277.18, "D4": 293.66, "D#4": 311.13, "E4": 329.63, "F4": 349.23, "F#4": 369.99, "G4": 392.00, "G#4": 415.30, "A4": 440.00, "A#4": 466.16, "B4": 493.88,
	"C5": 523.25, "C#5": 554.37, "D5": 587.33, "D#5": 622.25, "E5": 659.26, "F5": 698.46, "F#5": 739.99, "G5": 783.99, "G#5": 830.61, "A5": 880.00, "A#5": 932.33, "B5": 987.77,
	"C6": 1046.50, "C#6": 1108.73, "D6": 1174.66, "D#6": 1244.51, "E6": 1318.51, "F6": 1396.91, "F#6": 1479.98, "G6": 1567.98, "G#6": 1661.22, "A6": 1760.00, "A#6": 1864.66, "B6": 1975.53,
	"C7": 2093.00, "C#7": 2217.46, "D7": 2349.32, "D#7": 2489.02, "E7": 2637.02, "F7": 2793.83, "F#7": 2959.96, "G7": 3135.96, "G#7": 3322.44, "A7": 3520.00, "A#7": 3729.31, "B7": 3951.07,
	"C8": 4186.01,
}

type SoundWave struct {
	startTime    float64
	amplitute    int
	duration     float64
	freq         float64
	attack       float64
	sustainLevel float64
	decay        float64
	release      float64
}

func (s *SoundWave) envelop(sec float64) float64 {
	switch {
	case sec < 0:
		return 0
	case sec < s.attack:
		// Attack: ramp up 0 → 1
		return sec / s.attack
	case sec < s.attack+s.decay:
		// Decay: 1 → sustain level
		return 1 - (1-s.sustainLevel)*(sec-s.attack)/s.decay
	case sec < s.duration-s.release:
		// Sustain
		return s.sustainLevel
	case sec < s.duration:
		// Release: fade out
		return s.sustainLevel * (1 - (sec-(s.duration-s.release))/s.release)
	default:
		return 0
	}
}

func NewGuitarNoteWave(note string, startTime float64, amplitute int) *SoundWave {
	freq := musicNoteFreqM[note]
	return &SoundWave{
		freq:         freq,
		startTime:    startTime,
		amplitute:    5000,
		duration:     0.8,
		attack:       0.01,
		decay:        0.15,
		sustainLevel: 0.6,
		release:      0.5,
	}
}

// t: sec
func (s *SoundWave) Sample(t float64) float64 {
	if t < s.startTime {
		return 0
	}

	env := s.envelop(t - s.startTime)
	return float64(s.amplitute) * env * math.Sin(2*math.Pi*s.freq*(t-s.startTime))
}

func playromancedeamour() {
	const (
		sampleRate = 8000
	)

	// String 1: High E (e)

	// Transcription of the provided 2-measure tab snippet.

	// String 1: E (Melody)
	string1 := []string{
		// Measure 1
		"  ", "  ", "B4", "  ", "  ", "B4", "  ", "  ", "B4",
		// Measure 2
		"  ", "  ", "B4", "  ", "  ", "A4", "  ", "  ", "G4",
		// Measure 3
		"  ", "  ", "G4", "  ", "  ", "F#4", "  ", "  ", "E4",
		// Measure 4
		"  ", "  ", "E4", "  ", "  ", "G4", "  ", "  ", "B4",
		// Measure 5
		"  ", "  ", "E5", "  ", "  ", "E5", "  ", "  ", "E5",
		// Measure 6
		"  ", "  ", "E5", "  ", "  ", "D5", "  ", "  ", "C5",
		// Measure 7
		"  ", "  ", "C5", "  ", "  ", "B4", "  ", "  ", "A4",
		// Measure 8
		"  ", "  ", "A4", "  ", "  ", "B4", "  ", "  ", "C5",
		// Measure 9
		"  ", "  ", "B4", "  ", "  ", "C5", "  ", "  ", "B4",
		// Measure 10
		"  ", "  ", "D#5", "  ", "  ", "C5", "  ", "  ", "B4",
		// Measure 11
		"  ", "  ", "B4", "  ", "  ", "A4", "  ", "  ", "G4",
		// Measure 12
		"  ", "  ", "G4", "  ", "  ", "F#4", "  ", "  ", "E4",
		// Measure 13
		"F#4", "  ", "  ", "F#4", "  ", "  ", "F#4", "  ", "  ",
		// Measure 14
		"F#4", "  ", "  ", "G4", "  ", "  ", "F#4", "  ", "  ",
		// Measure 15
		"E4", "  ", "  ", "E4", "  ", "  ", "E4", "  ", "  ",
		// Measure 16 (Sextuplet, then Triplet)
		"  ", "  ", "  ", "  ", "  ", "E4", "E4", "  ", "  ",
		// Measure 17
		"G#4", "  ", "  ", "G#4", "  ", "  ", "G#4", "  ", "  ",
		// Measure 18
		"G#4", "  ", "  ", "F#4", "  ", "  ", "E4", "  ", "  ",
		// Measure 19
		"A4", "  ", "  ", "G#4", "  ", "  ", "G#4", "  ", "  ",
		// Measure 20
		"G#4", "  ", "  ", "G4", "  ", "  ", "G#4", "  ", "  ",
		// Measure 21
		"C#5", "  ", "  ", "C#5", "  ", "  ", "C#5", "  ", "  ",
		// Measure 22
		"C#5", "  ", "  ", "D#5", "  ", "  ", "C#5", "  ", "  ",
		// Measure 23
		"C#5", "  ", "  ", "B4", "  ", "  ", "B4", "  ", "  ",
		// Measure 24
		"B4", "  ", "  ", "C#5", "  ", "  ", "D#5", "  ", "  ",
	}

	// String 2: B (Harmony)
	string2 := []string{
		// Measure 1
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 2
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 3
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 4
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 5
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 6
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 7
		"  ", "E4", "  ", "  ", "E4", "  ", "  ", "E4", "  ",
		// Measure 8
		"  ", "E4", "  ", "  ", "E4", "  ", "  ", "E4", "  ",
		// Measure 9
		"  ", "F#4", "  ", "  ", "F#4", "  ", "  ", "F#4", "  ",
		// Measure 10
		"  ", "F#4", "  ", "  ", "F#4", "  ", "  ", "F#4", "  ",
		// Measure 11
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 12
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 13
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 14
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 15
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 16
		"  ", "  ", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 17
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 18
		"  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3", "  ",
		// Measure 19
		"  ", "C#4", "  ", "  ", "C#4", "  ", "  ", "C#4", "  ",
		// Measure 20
		"  ", "C#4", "  ", "  ", "C#4", "  ", "  ", "C#4", "  ",
		// Measure 21
		"  ", "B4", "  ", "  ", "B4", "  ", "  ", "B4", "  ",
		// Measure 22
		"  ", "B4", "  ", "  ", "B4", "  ", "  ", "B4", "  ",
		// Measure 23
		"  ", "D#5", "  ", "  ", "D#5", "  ", "  ", "D#5", "  ",
		// Measure 24
		"  ", "D#5", "  ", "  ", "D#5", "  ", "  ", "D#5", "  ",
	}

	// String 3: G (Harmony)
	string3 := []string{
		// Measure 1
		"  ", "G3", "  ", "  ", "G3", "  ", "  ", "G3", "  ",
		// Measure 2
		"  ", "G3", "  ", "  ", "G3", "  ", "  ", "G3", "  ",
		// Measure 3
		"  ", "G3", "  ", "  ", "G3", "  ", "  ", "G3", "  ",
		// Measure 4
		"  ", "G3", "  ", "  ", "G3", "  ", "  ", "G3", "  ",
		// Measure 5
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 6
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 7
		"  ", "C4", "  ", "  ", "C4", "  ", "  ", "C4", "  ",
		// Measure 8
		"  ", "C4", "  ", "  ", "C4", "  ", "  ", "C4", "  ",
		// Measure 9
		"  ", "D#4", "  ", "  ", "D#4", "  ", "  ", "D#4", "  ",
		// Measure 10
		"  ", "D#4", "  ", "  ", "D#4", "  ", "  ", "D#4", "  ",
		// Measure 11
		"  ", "G3", "  ", "  ", "G3", "  ", "  ", "G3", "  ",
		// Measure 12
		"  ", "G3", "  ", "  ", "G3", "  ", "  ", "G3", "  ",
		// Measure 13
		"  ", "  ", "A3", "  ", "  ", "A3", "  ", "  ", "A3",
		// Measure 14
		"  ", "  ", "A3", "  ", "  ", "A3", "  ", "  ", "A3",
		// Measure 15
		"  ", "  ", "G3", "  ", "  ", "G3", "  ", "  ", "G3",
		// Measure 16
		"  ", "  ", "  ", "G3", "  ", "  ", "  ", "  ", "G3",
		// Measure 17
		"  ", "  ", "G#3", "  ", "  ", "G#3", "  ", "  ", "G#3",
		// Measure 18
		"  ", "  ", "G#3", "  ", "  ", "G#3", "  ", "  ", "G#3",
		// Measure 19
		"  ", "  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3",
		// Measure 20
		"  ", "  ", "B3", "  ", "  ", "B3", "  ", "  ", "B3",
		// Measure 21
		"  ", "  ", "E4", "  ", "  ", "E4", "  ", "  ", "E4",
		// Measure 22
		"  ", "  ", "E4", "  ", "  ", "E4", "  ", "  ", "E4",
		// Measure 23
		"  ", "  ", "F#4", "  ", "  ", "F#4", "  ", "  ", "F#4",
		// Measure 24
		"  ", "  ", "F#4", "  ", "  ", "F#4", "  ", "  ", "F#4",
	}

	// String 4: D
	string4 := []string{
		// All measures are empty for this interpretation
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 13 & 14 are empty
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 15
		"  ", "  ", "  ", "  ", "  ", "  ", "F3", "  ", "  ",
		// Measure 16
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 17 & 18 are empty
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 19 & 20 are empty
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 21 & 22 are empty
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 23 & 24 are empty
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
	}

	// String 5: A (Bass)
	string5 := []string{
		// Measure 1
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 2
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 3
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 4
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 5
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 6
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 7
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 8
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 9
		"E3", "  ", "  ", "E3", "  ", "  ", "E3", "  ", "  ",
		// Measure 10
		"E3", "  ", "  ", "E3", "  ", "  ", "E3", "  ", "  ",
		// Measure 11
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 12
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 13
		"B2", "  ", "  ", "B2", "  ", "  ", "B2", "  ", "  ",
		// Measure 14
		"B2", "  ", "  ", "B2", "  ", "  ", "B2", "  ", "  ",
		// Measure 15
		"B2", "  ", "  ", "B2", "  ", "  ", "  ", "  ", "  ",
		// Measure 16
		"  ", "B2", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 17
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 18
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 19
		"B2", "  ", "  ", "B2", "  ", "  ", "B2", "  ", "  ",
		// Measure 20
		"B2", "  ", "  ", "B2", "  ", "  ", "B2", "  ", "  ",
		// Measure 21
		"E3", "  ", "  ", "E3", "  ", "  ", "E3", "  ", "  ",
		// Measure 22
		"E3", "  ", "  ", "E3", "  ", "  ", "E3", "  ", "  ",
		// Measure 23
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 24
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
	}

	// String 6: Low E (Bass)
	string6 := []string{
		// Measure 1
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 2
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 3
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 4
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 5
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 6
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 7
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 8
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 9
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 10
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 11
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 12
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 13
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 14
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 15
		"  ", "  ", "  ", "  ", "  ", "  ", "G2", "  ", "  ",
		// Measure 16
		"E2", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 17
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 18
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 19
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 20
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 21
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 22
		"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ",
		// Measure 23
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
		// Measure 24
		"E2", "  ", "  ", "E2", "  ", "  ", "E2", "  ", "  ",
	}
	waves := []*SoundWave{}
	for _, tabs := range [][]string{string1, string2, string3, string4, string5, string6} {
		for i, note := range tabs {
			wave := NewGuitarNoteWave(note, float64(i)*0.20, 3000)
			waves = append(waves, wave)
		}
	}

	durSec := 50
	pcmdata := []byte{}
	for sec := range durSec {
		for sample := 0; sample < sampleRate; sample++ {
			t := float64(sec) + float64(sample)/float64(sampleRate) // ms
			var s float64
			for _, wave := range waves {
				s += wave.Sample(t)
			}

			data := []byte{0x00, 0x00}
			binary.LittleEndian.PutUint16(data[:], uint16(s))
			pcmdata = append(pcmdata, data...)
		}
	}

	// Initialize Oto audio context
	opts := &oto.NewContextOptions{
		SampleRate:   sampleRate,
		ChannelCount: 1,
		Format:       oto.FormatSignedInt16LE,
	}
	ctx, ready, err := oto.NewContext(opts)
	if err != nil {
		panic(err)
	}
	<-ready // wait until ready

	player := ctx.NewPlayer(bytes.NewReader(pcmdata))
	player.Play()

	fmt.Println("Playing... Press Ctrl+C to exit.")
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	player.Close()
}
