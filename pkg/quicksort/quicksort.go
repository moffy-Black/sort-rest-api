package quicksort

import (
	"errors"
	"math"
)

func QuickSort(recvArr []int) ([]int, error) {
	length := len(recvArr)

	if length > int(math.Pow10(5)) {
		return recvArr, errors.New("配列の要素数が10^5を超えています")
	} else if length == 0 {
		return recvArr, errors.New("配列の要素数は0です")
	}

	var smallGroup []int
	for _, v := range recvArr[1:] {
		if v <= recvArr[0] {
			smallGroup = append(smallGroup, v)
		}
	}
	if len(smallGroup) > 0 {
		smallGroup, _ = QuickSort(smallGroup)
	}

	smallGroup = append(smallGroup, recvArr[0])

	var bigGroup []int
	for _, v := range recvArr[1:] {
		if v > recvArr[0] {
			bigGroup = append(bigGroup, v)
		}
	}
	if len(bigGroup) > 0 {
		bigGroup, _ = QuickSort(bigGroup)
	}
	ansArr := append(smallGroup, bigGroup...)
	return ansArr, nil
}
