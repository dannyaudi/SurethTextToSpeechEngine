package main

import "testing"

func TestCombineAudio(t *testing.T) {

	paths := []string{"audio/1811/1811.wav", "audio/1824/1852/1808.wav", "audio/1836/1845/1845.wav"} //path for gleeta, "ܓܠܝܼܬܵܐ"

	combineAudio(paths, "output1.wav")

	paths = []string{"audio/1829/1845/1816.wav", "audio/1824/1845/1808.wav"} // path for eawla, "ܥܵܘܠܵܐ"

	combineAudio(paths, "output2.wav")

	paths = []string{"audio/1834/1848/1818.wav", "audio/1819/1845/1808.wav"} // path for rikhta, "ܪܸܚܛܵܐ"

	combineAudio(paths, "output3.wav")

	paths = []string{"audio/1836/1842/1824.wav", "audio/1824/1852/1808.wav", "audio/1824/1845/1808.wav"} // path for taleela, "ܬܲܠܝܼܠܵܐ"

	combineAudio(paths, "output4.wav")

	paths = []string{"audio/1824/1845/1808.wav", "audio/1819/1845/1808.wav"} // path for lata, "ܠܵܛܵܐ"

	combineAudio(paths, "output5.wav")
}
