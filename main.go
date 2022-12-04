package main

import (
	"fmt"

	"github.com/moffy-Black/sort-server/pkg/bubblesort"
	"github.com/moffy-Black/sort-server/pkg/heapsort"
	"github.com/moffy-Black/sort-server/pkg/quicksort"
)

func main() {
	sampleArr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	bubbleArr, err := bubblesort.BubbleSort(sampleArr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(`BubbleSort(sampleArr) = `, bubbleArr)

	heapArr, err := heapsort.HeapSort(sampleArr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(`HeapSort(sampleArr) = `, heapArr)

	quickArr, err := quicksort.QuickSort(sampleArr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(`QuickSort(sampleArr) = `, quickArr)
}
