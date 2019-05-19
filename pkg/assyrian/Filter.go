package assyrian

// This function is called at the start of a word processing
// and removes all characters unnecessary to the pronounciation
// of the word. A few important codes include
// 1797 and 776 - these refer to syameh, two dots that indicate plurality
// 1863 - m6alqana, a symbol which denotes a silent letter
//
// Additionally, it simplifies the word by replacing characters that can
// be expressed in a simpler form, such as the replaceMorkikha function.
// Letters which become silent due to grammatical rules are also removed.

// IF DALAT SHLEETA BEFORE TAW, CHECK TAW BEFORE DALAT
// TWO TAW SHILYEH BAR OODALEH, AND REMOVE TWO DOTS ABOVE
// ZAYN BECOMES SIMKATH

func Filter(word []rune) []rune {

	for i := 0; i < len(word); i++ {
		a := word[i]
		if a == 1797 || a == 776 { // syameh
			word = remove(word, i)
		} else if a == 1863 { // m6alqana
			word = removeTwo(word, i)
		} else if a == 1858 { // rukkakha
			word = replaceMorkikha(word, i, word[i-1])
		} else if a == 1836 { // taw, for dalat before rule
			if word[i-1] == 1813 {
				if IsShleeta(word, i-1) {
					remove(word, i-1)
				}
			}
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

// used to replace bet and cap with waw and khet respectively when morkikheh
func replaceMorkikha(slice []rune, s int, a rune) []rune {
	if a == 1810 {
		slice[s] = 1816
		return append(slice[:s-1], slice[s:]...)
	} else if a == 1823 {
		slice[s] = 1818
		return append(slice[:s-1], slice[s:]...)
	} else { // letter must be gamal, dalat, or taw, do not replace
		return slice
	}
}
