package utils

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
)

func overlap(paths []string, name string) {

	// Opening the first clip in the path
	i, _ := os.Open(paths[0])
	streamer1, format1, _ := wav.Decode(i)
	defer streamer1.Close()

	// Opening the second clip in the path
	j, _ := os.Open(paths[1])
	streamer2, _, _ := wav.Decode(j)
	defer streamer2.Close()

	// Mixing the first two clips, with a 0.075 second (3300 sample) overlap between them
	// This is done by adding silence for the length of streamer1 minus 0.075 seconds to the beginning of streamer 2
	// And then mixing these two clips over each other
	mixed := beep.Mix(
		streamer1,
		beep.Seq(beep.Silence(streamer1.Len()-3300), streamer2),
	)

	// This loop iterates over the list of paths, mixing all clips with a 0.075 second delay
	for x := 2; x < len(paths); x++ {

		// Reading the next clip in the paths, creating a temporary streamer
		streamerTemp, _ := os.Open(paths[x])
		streamerT, _, _ := wav.Decode(streamerTemp)
		defer streamerTemp.Close()

		// Creating a buffer to hold the previously mixed clips
		// The buffer was implemented in order to make use of the .Len() function in mixing during this loop
		mixedBuffer := beep.NewBuffer(format1)
		mixedBuffer.Append(mixed)

		// Converting the buffer to a streamer
		mixedStreamer := mixedBuffer.Streamer(0, mixedBuffer.Len())

		// Mixing the previous mix, and the next clip which is streamerT, with the 0.075 second (3300 sample) delay
		mixed = beep.Mix(
			mixedStreamer,
			beep.Seq(beep.Silence(mixedStreamer.Len()-3300), streamerT),
		)
	}

	// Writing the new file
	writer, _ := os.Create(name)
	wav.Encode(writer, mixed, format1)

}
