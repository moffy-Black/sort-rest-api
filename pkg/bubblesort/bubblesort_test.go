package bubblesort

import (
	"errors"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func testBubbleSort(recvArr []int) ([]int, error) {
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

func TestBubbleSort(t *testing.T) {
	testArr := rand.Perm(int(math.Pow10(4)))
	bubbleArr, err := testBubbleSort(testArr)
	if err != nil {
		t.Fatalf(`len(BubbleSort("testArr")) = %d, %v`, len(bubbleArr), err)
	}
}
