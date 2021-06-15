package quicksort

import (
	"flag"
	"fmt"
	"github.com/psilva261/timsort"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestQuicksort(t *testing.T) {
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
	quickSort(testData1, 0, len(testData1)-1)
	fmt.Println("single goroutine: ", time.Now().Sub(start))
	//fmt.Println(testData1)
	if !sort.IntsAreSorted(testData1) {
		fmt.Println("wrong quick_sort implementation")
	}
	done := make(chan struct{})
	start = time.Now()
	go quickSort_go(testData2, 0, len(testData2)-1, done, 15)
	<-done
	fmt.Println("multiple goroutine: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData2) {
		fmt.Println("wrong quickSort_go implementation")
	}
	start = time.Now()
	sort.Ints(testData3)
	fmt.Println("std lib: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData3) {
		fmt.Println("wrong std lib implementation")
	}
	start = time.Now()
	timsort.Ints(testData4, func(a, b int) bool { return a <= b })
	fmt.Println("timsort: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData4) {
		fmt.Println("wrong timsort implementation")
	}
}
