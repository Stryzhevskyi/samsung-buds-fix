package main

import (
	"fmt"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

func main() {
	// Initialize the speaker
	err := speaker.Init(beep.SampleRate(1000), beep.SampleRate(1000).N(time.Second))
	if err != nil {
		fmt.Println("Failed to initialize speaker:", err)
		return
	}

	// Generate silence
	silence := beep.Silence(-1)

	// Play silence
	done := make(chan bool)
	speaker.Play(beep.Seq(silence, beep.Callback(func() {
		close(done)
	})))

	// Wait for the silence to finish playing
	<-done

	// Close the speaker
	speaker.Close()
}
