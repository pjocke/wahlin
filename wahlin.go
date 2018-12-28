package wahlin

import (
	"strings"
)

func consonant(letter rune) bool {
	var consonants = [20]rune{'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Z'}

	var consonant rune
	for _, consonant = range consonants {
		if consonant == letter {
			return true
		}
	}
	return false
}

// Encode a string to its Wåhlin encoded equivalent
func Encode(decoded string) string {
	var word = []rune(strings.ToUpper(decoded))
	var encoded []rune

	var i int
	var letter rune
	for i = 0; i < len(word); i++ {
		letter = word[i]

		if letter == 'A' {
			// Rule set A
			if i+1 < len(word) && word[i+1] == 'E' {
				continue
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'C' && i+1 < len(word) && word[i+1] != 'H' {
			// Rule set C
			if word[i+1] == 'A' || word[i+1] == 'O' || word[i+1] == 'U' || word[i+1] == 'L' || word[i+1] == 'S' || word[i+1] == 'R' {
				encoded = append(encoded, 'K')
			} else if word[i+1] == 'E' || word[i+1] == 'I' || word[i+1] == 'Y' {
				encoded = append(encoded, 'S')
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'C' && i+1 < len(word) && word[i+1] == 'H' {
			// Rule set CH
			if i+2 < len(word) && (word[i+2] == 'R' || word[i+2] == 'S' || word[i+2] == 'T') {
				encoded = append(encoded, 'K')
			}
		} else if letter == 'D' {
			// Rule set D
			if i+1 < len(word) && (word[i+1] == 'J' || word[i+1] == 'T') {
				continue
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'F' {
			// Rule set F
			if (i > 0 && (word[i-1] == 'A' || word[i-1] == 'E' || word[i-1] == 'I' || word[i-1] == 'O' || word[i-1] == 'U' || word[i-1] == 'Y' || word[i-1] == 'Å' || word[i-1] == 'F' || word[i-1] == 'L' || word[i-1] == 'R')) ||
				(i+1 < len(word) && (word[i+1] == 'B' || word[i+1] == 'C' || word[i+1] == 'D' || word[i+1] == 'G' || word[i+1] == 'H' || word[i+1] == 'K' || word[i+1] == 'M' || word[i+1] == 'N' || word[i+1] == 'P' || word[i+1] == 'S' || word[i+1] == 'V')) {
				encoded = append(encoded, 'V')
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'G' {
			// Rule set G
			if (i-1 > 0 && (word[i-1] == 'R')) ||
				(i == 0 && i+1 < len(word) && word[i+1] == 'E') ||
				(i+1 < len(word) && (word[i+1] == 'I' || word[i+1] == 'Y' || word[i+1] == 'Ö')) {
				encoded = append(encoded, 'J')
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'H' {
			// Rule set H
			if i-1 > 0 && word[i-1] == 'T' {
				continue
			} else if i+1 < len(word) && consonant(word[i+1]) {
				continue
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'I' {
			// Rule set I
			if (i > 0 && (word[i-1] == 'H')) &&
				(i+1 < len(word) && (word[i+1] == 'A' || word[i+1] == 'E' || word[i+1] == 'O' || word[i+1] == 'U' || word[i+1] == 'Y' || word[i+1] == 'Å' || word[i+1] == 'Ö')) {
				encoded = append(encoded, 'J')
			} else if (i > 0) && (word[i-1] == 'A' || word[i-1] == 'E' || word[i-1] == 'O' || word[i-1] == 'Ö') {
				encoded = append(encoded, 'J')
			} else if i+1 < len(word) && (word[i+1] == 'E' || word[i+1] == 'O' || word[i+1] == 'Y' || word[i+1] == 'Ö') {
				encoded = append(encoded, 'J')
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'K' {
			// Rule set K, KJ
			if i+1 < len(word) && word[i+1] == 'J' {
				encoded = append(encoded, '+')
				i++
			} else if i == 0 && i+1 < len(word) && (word[i+1] == 'E' || word[i+1] == 'I' || word[i+1] == 'Y' || word[i+1] == 'Ö') {
				encoded = append(encoded, '+')
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'L' {
			if i+1 < len(word) && word[i+1] == 'J' {
				continue
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'L' {
			// Rule set L
			if i+1 < len(word) && word[i+1] == 'J' {
				continue
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'P' {
			// Rule set P, PH
			if (i+1 < len(word) && word[i+1] == 'H') && (i == 0 || (i > 0 && word[i-1] != 'P')) {
				encoded = append(encoded, 'F')
				i++
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'Q' {
			// Rule set Q, QU
			if i+1 < len(word) && word[i+1] == 'U' {
				encoded = append(encoded, 'K')
				encoded = append(encoded, 'V')
				i++
			} else {
				encoded = append(encoded, 'K')
			}
		} else if letter == 'S' {
			// Rule set SJ, SCH, STJ, SKJ, SK
			if i+1 < len(word) && word[i+1] == 'J' {
				encoded = append(encoded, '*')
				i++
			} else if i+2 < len(word) && word[i+1] == 'C' && word[i+2] == 'H' {
				encoded = append(encoded, '*')
				i = i + 2
			} else if i+2 < len(word) && word[i+1] == 'T' && word[i+2] == 'J' {
				encoded = append(encoded, '*')
				i = i + 2
			} else if i+2 < len(word) && word[i+1] == 'K' && word[i+2] == 'J' {
				encoded = append(encoded, '*')
				i = i + 2
			} else if i+2 < len(word) && word[i+1] == 'K' && word[i+2] == 'Ö' {
				encoded = append(encoded, '*')
				i++
			} else if i+2 < len(word) && word[i+1] == 'K' && word[i+2] == 'E' {
				encoded = append(encoded, '*')
				i++
			} else if i+2 < len(word) && word[i+1] == 'K' && word[i+2] == 'I' {
				encoded = append(encoded, '*')
				i++
			} else if i+2 < len(word) && word[i+1] == 'K' && word[i+2] == 'Y' {
				encoded = append(encoded, '*')
				i++
			} else {
				encoded = append(encoded, letter)
			}
		} else if letter == 'T' && i+1 < len(word) && word[i+1] == 'J' && i == 0 {
			// Rule set TJ
			encoded = append(encoded, '+')
			i++
		} else if letter == 'U' && i > 0 && word[i-1] == 'O' {
			// Rule set OU
			continue
		} else if letter == 'W' {
			// Rule set W
			encoded = append(encoded, 'V')
		} else if letter == 'X' {
			// Rule set X
			encoded = append(encoded, 'K')
			encoded = append(encoded, 'S')
		} else if letter == 'Y' && (i > 0 && (word[i-1] == 'A' || word[i-1] == 'E' || word[i-1] == 'O' || word[i-1] == 'U')) {
			// Rule set Y
			encoded = append(encoded, 'J')
		} else if letter == 'Z' {
			// Rule set Z
			encoded = append(encoded, 'S')
		} else if letter == 'Å' {
			// Rule set Å
			if (i+1 < len(word) && (word[i+1] == 'V' || word[i+1] == 'N')) || (i+2 < len(word) && word[i+2] == 'G') {
				encoded = append(encoded, 'O')
			}
		} else if letter == 'Ä' {
			// Rule set Ä
			encoded = append(encoded, 'E')
		} else {
			encoded = append(encoded, letter)
		}
	}

	return string(encoded)
}
