package quicksort

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQuickSortOptimizer(t *testing.T) {
	c := flag.Int("c", 100000000, "sort src length")
	count := *c
	rand.Seed(time.Now().UnixNano())
	testData1, testData2, testData3, testData4 := make([]int, 0, count), make([]int, 0, count), make([]int, 0, count), make([]int, 0, count)
	times := count
	for i := 0; i < times; i++ {
		val := rand.Intn(2 * count)
		testData1 = append(testData1, val)
		testData2 = append(testData2, val)
		testData3 = append(testData3, val)
		testData4 = append(testData4, val)
	}
	//fmt.Println(testData1)
	start := time.Now()
	done := make(chan struct{})
	go QuicksortOptimize(testData1, 0, len(testData1)-1, 15, done)
	<-done
	fmt.Println("single goroutine: ", time.Now().Sub(start))
	//fmt.Println(testData1)
}
