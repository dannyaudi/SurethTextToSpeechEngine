package main

// This file contains various small rules pertaining to rukkakha

// Checks if a character sent has no zoweh
func isShleeta(w []rune, i int) bool {
	if w[i+1] == 1845 { // zqapa
		return false
	} else if w[i+1] == 1842 { // ptakha
		return false
	} else if w[i+1] == 1848 { // zlameh keryeh
		return false
	} else if w[i+1] == 1849 { // zlameh yareekheh
		return false
	} else if w[i+1] == 1816 { // waw, checking for rwakha/rwasa
		if w[i+2] == 1852 || w[i+2] == 1855 {
			return false
		}
	} else if w[i+1] == 1821 { // yut, checking for khwasa
		if w[i+2] == 1852 {
			return false
		}
	}
	return true
}
