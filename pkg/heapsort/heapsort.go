package heapsort

import (
	"errors"
	"math"
)

func heapify(recvArr []int, length int, i int) []int {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < length && recvArr[i] < recvArr[l] {
		largest = l
	}

	if r < length && recvArr[largest] < recvArr[r] {
		largest = r
	}

	if largest != i {
		recvArr[i], recvArr[largest] = recvArr[largest], recvArr[i]
		return heapify(recvArr, length, largest)
	}
	return recvArr
}

func HeapSort(recvArr []int) ([]int, error) {
	length := len(recvArr)
	if length > int(math.Pow10(4)) {
		return recvArr, errors.New("配列の要素数が10^4を超えています")
	} else if length == 0 {
		return recvArr, errors.New("配列の要素数は0です")
	}

	for i := length/2 - 1; i > -1; i-- {
		recvArr = heapify(recvArr, length, i)
	}

	for i := length - 1; i > 0; i-- {
		recvArr[i], recvArr[0] = recvArr[0], recvArr[i]
		recvArr = heapify(recvArr, i, 0)
	}

	return recvArr, nil
}
