package quicksort

import (
	"errors"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func testQuickSort(recvArr []int) ([]int, error) {
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
		smallGroup, _ = testQuickSort(smallGroup)
	}

	smallGroup = append(smallGroup, recvArr[0])

	var bigGroup []int
	for _, v := range recvArr[1:] {
		if v > recvArr[0] {
			bigGroup = append(bigGroup, v)
		}
	}
	if len(bigGroup) > 0 {
		bigGroup, _ = testQuickSort(bigGroup)
	}
	ansArr := append(smallGroup, bigGroup...)
	return ansArr, nil
}

func makeAns(min, max int) []int {
	ans := make([]int, max-min)
	for i := range ans {
		ans[i] = min + i
	}
	return ans
}

func TestLengthQucikSort(t *testing.T) {
	testArr := rand.Perm(int(math.Pow10(5)))
	quickArr, err := testQuickSort(testArr)
	if err != nil {
		t.Fatalf(`len(QuickSort("testArr")) = %d, %v`, len(quickArr), err)
	}
}

func TestAnsQuickSort(t *testing.T) {
	testArr := rand.Perm(int(math.Pow10(5)))
	quickArr, _ := testQuickSort(testArr)
	ansArr := makeAns(0, int(math.Pow10(5)))
	if reflect.DeepEqual(ansArr, quickArr) != true {
		t.Fatalf("can't sort")
	}
}
