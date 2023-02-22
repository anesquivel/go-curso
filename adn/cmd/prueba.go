package main

import "fmt"

func main() {
	/*adn := "ABCEEFG"
	res := searchInBaseADN(adn)
	fmt.Println("res:", res)

	if res {
		fmt.Printf("¡%s es mutante!", adn)
	}*/

	adn := adnString()
	count := 0
	for i, partOfADN := range adn {
		resHorizontal := searchInBaseADN(partOfADN)
		resVertical := searchInBaseADN(sliceADN(adn, true, i))
		if i <= (len(adn) / 2) {
			fmt.Println("entra acá", sliceADN(adn, false, i))
			resRightToLeft := searchInBaseADN(sliceADN(adn, false, i))
			if resRightToLeft {
				count++
			}
		}
		if resHorizontal {
			count++
		}

		if resVertical {
			count++
		}
	}

	printRes(count)
}

func sliceADN(adn []string, vertical bool, pos int) string {
	newADN := ""
	for _, part := range adn {

		for i, char := range part {
			if i == pos {
				newADN += string(char)
			}
		}
		if !vertical {
			pos++
		}
	}

	return newADN
}

func searchInBaseADN(base string) bool {
	count, lenBase := 1, len(base)-1
	noChar := 0
	for i, char := range base {
		if char >= 65 && char <= 71 && i != lenBase {
			noChar = int(char)
			if base[i+1] != 0 && base[i+1] == byte(noChar) {
				count++
				if count == 4 {
					return true // tiene una cadena con 4 chars iguales
				}

			}
		}
	}
	return false
}

func adnString() []string {
	return []string{
		"AGGCGA",
		"CAGTGC",
		"TTAGGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}
}

func printRes(count int) {
	if count > 1 {
		fmt.Printf("¡FELICIDADES! Encontraste a un mutante.\nConteo de cadenas de ADN iguales: %d\n", count)
	} else {
		fmt.Printf("El inidividuo no es un mutante.\nConteo de cadenas de ADN iguales: %d\n", count)
	}
}
