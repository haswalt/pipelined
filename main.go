package main

import (
	"context"
	"os"
	"time"

	"pipelined.dev/audio"
	"pipelined.dev/audio/wav"
	"pipelined.dev/pipe"
)

var (
	numChannels = 2
	bufferSize  = 512
)

func main() {
	defer PrintElapsedTime(time.Now(), "complete")

	PrintMemoryUsage("Start")

	mixer := audio.NewMixer(numChannels)

	var lines []pipe.Line

	for i := 0; i < 4; i++ {
		track := &audio.Track{
			SampleRate: 22050,
			Channels:   2,
		}

		// This should use a repeat rather than multiple assets
		for j := 0; j < 20; j++ {
			f, _ := os.Open("./sample.wav")

			source := wav.Source{ReadSeeker: f}

			asset := audio.Asset{}

			// This feels like an unnessecary extra step and causes a copy
			l, _ := pipe.Routing{
				Source: source.Source(),
				Sink:   asset.Sink(),
			}.Line(bufferSize)

			pipe.New(context.Background(), pipe.WithLines(l)).Wait()

			f.Close()

			// just use the entire asset for now

			track.AddClip(j*4, asset.Floating)
		}

		line, _ := pipe.Routing{
			Source: track.Source(0, 0),
			Sink:   mixer.Sink(),
		}.Line(bufferSize)

		lines = append(lines, line)
	}

	PrintMemoryUsage("Pre Mix")

	out, _ := os.Create("demo.wav")
	defer out.Close()

	sink := &wav.Sink{
		WriteSeeker: out,
		BitDepth:    16,
	}

	line, _ := pipe.Routing{
		Source: mixer.Source(),
		Sink:   sink.Sink(),
	}.Line(bufferSize)

	lines = append(lines, line)

	pipe.New(context.Background(), pipe.WithLines(lines...)).Wait()

	PrintMemoryUsage("Post mix")
}
