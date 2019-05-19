package utils

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
)

// This function receives a slice of the paths for all audio clips to be combined to form a word.
// It iterates over the list, combining the next clip with the base clip, or "streamer".
// The file is written to the same directory under the name "output.wav"

func combineAudio(paths []string, output string) {

	f, _ := os.Open(paths[0])
	streamerBase, format, _ := wav.Decode(f)
	defer streamerBase.Close()

	streamerFinal := beep.Seq(streamerBase)

	for i := 1; i < len(paths); i++ {

		g, _ := os.Open(paths[i])

		streamerTemp, _, _ := wav.Decode(g)

		streamerFinal = beep.Seq(streamerFinal, streamerTemp)
	}

	writer, _ := os.Create(output) //temporarily changed from output.wav to output variable

	wav.Encode(writer, streamerFinal, format)
}
