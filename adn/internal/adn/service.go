package adn

import (
	"errors"

	"github.com/anesquivel/go-curso/adn/internal/domain"
)

type service struct {
	adn   domain.ADN
	count int
}

type Service interface {
	IsMutant() (bool, error)
	PlusCount()
}

func NewService(adn domain.ADN) Service {
	return &service{
		adn:   adn,
		count: 0,
	}
}

func sliceADN(adn []string, vertical bool, pos int) (string, error) {
	if len(adn) == 0 {
		return "", errors.New("No hay secuencias en el ADN.")
	}
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

	return newADN, nil
}

func searchInSecuenceADN(base string) (bool, error) {
	count, lenBase := 1, len(base)-1
	noChar := 0
	for i, char := range base {
		if char >= 65 && char <= 71 && i != lenBase {
			noChar = int(char)
			if base[i+1] != 0 && base[i+1] == byte(noChar) {
				count++
				if count == 4 {
					return true, nil
				}

			}
		} else {
			return false, errors.New("Caracter inválido en secuencia")
		}
	}
	return false, nil
}

func horizontalSearching(secuence string) (bool, error) {
	return searchInSecuenceADN(secuence)
}

func verticalSearching(adn domain.ADN, i int) (bool, error) {
	secuence, err := sliceADN(adn.Secuence, true, i)

	if err != nil {

	}
	return searchInSecuenceADN(secuence)
}
func verticalSearchinv(secuence string) (bool, error)
func (s *service) IsMutant() (bool, error) {
	if len(s.adn.Secuence) == 0 {
		return false, errors.New("No hay secuencias registradas")
	}
	for i, partOfADN := range s.adn.Secuence {
		resHorizontal, err := searchInSecuenceADN(partOfADN)

		if err != nil {
			//mal horizontal
			return false, err
		}
		resVertical, err := verticalSearching(s.adn, i)

		if err != nil {
			//mal vertical
			return false, err
		}
		if i <= (len(s.adn.Secuence) / 2) {
			resultADN, err := sliceADN(s.adn.Secuence, false, i)
			resRightToLeft, err := searchInSecuenceADN(resultADN)

			if err != nil {
				//mal vertical
				return false, err
			}

			if resRightToLeft { //si la secuencia en diagonal retornó true
				s.PlusCount()
			}
		}
		if resHorizontal { //si la secuencia en horizontal retornó true
			s.PlusCount()
		}

		if resVertical { // si la secuencia en vertical retornó true
			s.PlusCount()
		}
	}

	return s.count > 1, nil
}

func (s *service) PlusCount() {
	s.count++
}
