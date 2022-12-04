package bubblesort

import (
	"errors"
	"math"
)

func BubbleSort(recvArr []int) ([]int, error) {
	length := len(recvArr)

	if length > int(math.Pow10(4)) {
		return recvArr, errors.New("配列の要素数が10^4を超えています")
	} else if length == 0 {
		return recvArr, errors.New("配列の要素数は0です")
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length-(i+1); j++ {
			if recvArr[j] > recvArr[j+1] {
				recvArr[j], recvArr[j+1] = recvArr[j+1], recvArr[j]
			}
		}
	}

	return recvArr, nil
}
