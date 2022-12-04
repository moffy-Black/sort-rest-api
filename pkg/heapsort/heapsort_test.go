package heapsort

import (
	"errors"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func testheapify(recvArr []int, length int, i int) []int {
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
		return testheapify(recvArr, length, largest)
	}
	return recvArr
}

func testHeapSort(recvArr []int) ([]int, error) {
	length := len(recvArr)
	if length > int(math.Pow10(4)) {
		return recvArr, errors.New("配列の要素数が10^4を超えています")
	} else if length == 0 {
		return recvArr, errors.New("配列の要素数は0です")
	}

	for i := length/2 - 1; i > -1; i-- {
		recvArr = testheapify(recvArr, length, i)
	}

	for i := length - 1; i > 0; i-- {
		recvArr[i], recvArr[0] = recvArr[0], recvArr[i]
		recvArr = testheapify(recvArr, i, 0)
	}

	ansArr := makeAns(0, length)

	if reflect.DeepEqual(ansArr, recvArr) != true {
		return recvArr, errors.New("ソートが失敗しました")
	}
	return recvArr, nil
}

func makeAns(min, max int) []int {
	ans := make([]int, max-min)
	for i := range ans {
		ans[i] = min + i
	}
	return ans
}

func TestHeapSort(t *testing.T) {
	testArr := rand.Perm(int(math.Pow10(2)))
	heapArr, err := testHeapSort(testArr)
	if err != nil {
		t.Fatalf(`len(HeapSort("testArr")) = %d, %v`, len(heapArr), err)
	}
}
