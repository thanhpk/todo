package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

func linearToALaw(sample int16) byte {
	const (
		A = 87.6
	)
	sign := 0
	if sample < 0 {
		sign = 1
		sample = -sample
	}
	abs := float64(sample)
	var compressed float64
	if abs < 256 {
		compressed = abs / (1 + math.Log(A))
	} else {
		compressed = (1 + math.Log(abs/256)) / (1 + math.Log(A))
	}
	// Map 0..1 → 0..255
	v := byte(math.Min(255, compressed*255))
	if sign == 1 {
		v ^= 0x55 // flip bits for negative
	}
	return v
}

func main2() {
	const sampleRate = 8000
	const freq = 1000.0
	const durationSec = 5.0

	samples := int(sampleRate * durationSec)
	f, _ := os.Create("hello.alaw")
	defer f.Close()

	for n := 0; n < samples; n++ {
		t := float64(n) / sampleRate
		s := int16(3000 * math.Sin(2*math.Pi*freq*t))
		b := linearToALaw(s)
		f.Write([]byte{b})
	}
}

// just print raw stream
func main3() {
	const sampleRate = 8000 // samples per second
	const freq = 1000.0     // frequency (Hz)
	const durationSec = 12  // shorter duration (2 ms) to print fewer samples

	samples := int(sampleRate * durationSec)

	for n := 0; n < samples; n++ {
		t := float64(n) / sampleRate
		s := int16(3000 * math.Sin(2*math.Pi*freq*t))
		fmt.Printf("sample[%02d] = %6d\n", n, s)
	}
}

func generatePCMFile() {
	const sampleRate = 8000
	const freq = 1000.0
	const durationSec = 5.0

	samples := int(sampleRate * durationSec)
	f, _ := os.Create("hello.pcm")
	defer f.Close()

	for n := 0; n < samples; n++ {
		t := float64(n) / sampleRate
		s := int16(3000 * math.Sin(2*math.Pi*freq*t))
		binary.Write(f, binary.LittleEndian, s)
	}
}

func main() {
	// Ensure the user provides at least one command
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [playpcm|subtract] -a <num> -b <num>")
		os.Exit(1)
	}

	// Get the subcommand (first argument)
	cmd := os.Args[1]

	switch cmd {
	case "playpcm":
		playpcm()
	case "playromancedeamour":
		playromancedeamour()
	default:
		fmt.Println("Expected 'pcm' or 'subtract' subcommands")
		os.Exit(1)
	}
}
