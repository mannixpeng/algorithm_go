package quicksort

func Partition(a []int, l, h int) int {
	i := l - 1
	pivot := a[h]
	for j := l; j < h; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[h] = a[h], a[i+1]
	return i + 1
}

func QuicksortOptimize(a []int, l, h, depth int, done chan struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	if l >= h{
		return
	}
	depth--
	p := Partition(a, l, h)

	if depth > 0 {
		ch := make(chan struct{}, 2)
		go QuicksortOptimize(a, l, p-1, depth, ch)
		go QuicksortOptimize(a, p+1, h, depth, ch)
		<-ch
		<-ch
	} else {
		quickSort(a, l, p-1)
		quickSort(a, p+1, h)
	}
}
