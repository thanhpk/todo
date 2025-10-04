# How sound is play

`hello world` this is a text stream, made up of ascii characters

`RGBA RGBA RGBA` this is a stream of colored-pixels, made up of 4 integers number corresponding to red, green, blue and alpha value

a stream of audio is made up of sound samples

```
[Microphone]
   ↓ (Analog → ADC)
[Linear PCM (16-bit)]
   ↓ (Compression)
[A-law 8-bit stream]
   ↓ (Transmit over network)
[Receive side]
   ↓ (Decompress)
[Linear PCM (16-bit)]
   ↓ (Play on speaker)
```

## Pulse Code Modulation (PCM)
Pulse Code Modulation (PCM) is the most basic, raw form of digital audio.

I mean: "Representing an analog sound wave as a sequence of numbers."

Every number is a snapshot of the sound's air pressure at a tiny moment in time.


### How sound becomes PCM

Imagine a microphone.

It converts air pressure → voltage (analog wave).

To digitize it:

1.	Sampling – measure the wave’s height at regular time intervals.
2.	Quantization – round each measurement to the nearest integer.
3.	Encoding – store those integers as binary numbers (bytes).


### PCM's key parameters
* Sample Rate: how many samples per second. Example 8000Hz (telephony, 44100 Hz (CD)
* Bit Depth: how many bits per sample. higher means capture more detailed sound, high fidelity, lower means more noise. Eg: 8-bit, 16-bit, 24-bit
* Channels: number of tracks (each channel = one microphone). Eg: Mono (1), Stereo (2)

Example:

PCM @ 8000 Hz, 16-bit mono
→ 8000 samples × 16 bits = 128,000 bits/sec = 128 kbps

That's raw uncompressed audio

### Play with PCM
This code generate sample PCM stream and then play it so you can hear it in the speaker
```
go run . playpcm
```

### Narrowed band
The range of frequencies that can be transmitted is narrow — not the full human hearing range.
So when we say narrowband audio, we’re talking about how much of the sound spectrum is captured and transmitted.

### Human hearing range
Humans can hear roughly 20 Hz – 20 kHz. That's the full spectrum — from deep bass to sharp treble.

But for speech, most of the important information lives between: `300 Hz – 3400 Hz`

### Visualization
|<-- Bass -->|<-- Mids -->|<-- Treble -->|
0Hz        300Hz       3.4kHz        20kHz

## Alaw

PCM has linear amplitude steps. For example, if it’s 16-bit:

```
-32768  ...  0  ...  +32767
```

That's a huge range.

But human ears don't hear loudness linearly — we're more sensitive to small changes at quiet levels than at loud levels.

So, storing all 16 bits with equal precision is wasteful — we don't need that much detail for loud sounds.


### The idea: Companding

A-law (and μ-law) use a process called companding = compressing + expanding.

It means:
- When encoding: Compress the dynamic range — make quiet sounds bigger, loud sounds smaller.
- When decoding: Expand it back to approximate the original.

This saves space while keeping perceptual quality.


## How to play alaw file

```sh
ffplay -f alaw -ar 8000 -ch_layout mono  hello.alaw
```
* `-f alaw`: raw A-law data
* `-ar 8000`: 8 kHz sampling rate (G.711 standard)
* `-ch_layout` mono: sets channel layout to 1-channel
