// Very rough draft, extremely subject to change

package main

import (
	"SurethTextToSpeechEngine/pkg/assyrian"
	"fmt"
	"strconv"
)

func main() {

	input := "ܛܘܼܒ݂ܵܐ"

	word := []rune(input)
	fmt.Println(word)

	word = assyrian.Filter(word)
	fmt.Println(word)

	var paths []string

	for i := 0; i < len(word); i++ {

		var clipPath string

		currentLetter := word[i]

		// checks if we are on the last letter of the word
		if nextLetter(word, i) == -1 {
			break
		}

		// first letter is shleeta
		if assyrian.IsShleeta(word, i) { // was previously if nextCharacter == word[indexOfSecondLetter]
			clipPath = "/" + strconv.Itoa(int(currentLetter)) + "/" + strconv.Itoa(int(currentLetter)) + ".wav"
			fmt.Println(clipPath)
			paths = append(paths, clipPath)
			continue
		}
		fmt.Println("hi")
		// first letter has zoweh (including khwasa rwasa rwakha)

	}
}

// this function finds the next letter in the word, used in many situations
// returns the index of the next letter, skipping over other symbols.
// if the index received is the last letter in the word, -1 is returned
func nextLetter(w []rune, a int) int {

	if a == (len(w) - 1) {
		return -1
	}
	for i := a + 1; i < len(w); i++ {
		if w[i] >= 1808 && w[i] <= 1836 {
			return i
		}
	}
	return -1
}
