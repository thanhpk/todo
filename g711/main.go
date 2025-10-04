package main

import (
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

func main() {
	const sampleRate = 8000
	const freq = 1000.0
	const durationSec = 1.0

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
