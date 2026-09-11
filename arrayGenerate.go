package arrayGenerate

import (
	"errors"
	"math/rand/v2"
)

func ArrayGenerate(length int) ([]int, error) {

	if length < 1 {
		return nil, errors.New("la taille doit être de 1 ou plus")
	}

	result := make([]int, length)
	for i := range length {
		result[i] = rand.IntN(10000) + 1
	}
	return result, nil
}
