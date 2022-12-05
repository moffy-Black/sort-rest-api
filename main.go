package main

import (
	"fmt"
	"net/http"

	"github.com/moffy-Black/sort-rest-API/pkg/bubblesort"
	"github.com/moffy-Black/sort-rest-API/pkg/heapsort"
	"github.com/moffy-Black/sort-rest-API/pkg/quicksort"

	"github.com/gin-gonic/gin"
)

type arr struct {
	SortType string `json:"sortType"`
	Array    []int  `json:"array"`
}

func responseSortArr(c *gin.Context) {
	var recvArr arr

	if err := c.BindJSON(&recvArr); err != nil {
		return
	}

	switch recvArr.SortType {
	case "bubblesort":
		bubbleArr, err := bubblesort.BubbleSort(recvArr.Array)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusBadRequest, err)
		}
		fmt.Println(`BubbleSort(sampleArr) = `, bubbleArr)
		c.IndentedJSON(http.StatusOK, bubbleArr)
	case "heapsort":
		heapArr, err := heapsort.HeapSort(recvArr.Array)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusBadRequest, err)
		}
		fmt.Println(`HeapSort(sampleArr) = `, heapArr)
		c.IndentedJSON(http.StatusOK, heapArr)
	case "quicksort":
		quickArr, err := quicksort.QuickSort(recvArr.Array)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusBadRequest, err)
		}
		fmt.Println(`QuickSort(sampleArr) = `, quickArr)
		c.IndentedJSON(http.StatusOK, quickArr)
	}
}

func main() {
	router := gin.Default()
	router.POST("/", responseSortArr)
	router.Run("localhost:8080")
}
