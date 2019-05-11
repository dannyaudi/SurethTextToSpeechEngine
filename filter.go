package main

// This function is called at the start of a word processing
// and removes all characters unnecessary to the pronounciation
// of the word. A few important codes include
// 1797 and 776 - these refer to syameh, two dots that indicate plurality
// 1863 - m6alqana, a symbol which denotes a silent letter

func filter(word []rune) []rune {

	for i := 0; i < len(word); i++ {
		a := word[i]
		if a == 1797 || a == 776 {
			word = remove(word, i)
		} else if a == 1863 {
			word = removeTwo(word, i)
		}
	}
	return word
}

func remove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}

// used exclusively for cases of a m6alqana
func removeTwo(slice []rune, s int) []rune {
	return append(slice[:s-1], slice[s+1:]...)
}
