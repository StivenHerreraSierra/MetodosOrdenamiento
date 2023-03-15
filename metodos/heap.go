package metodos

func HeapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		temp := arr[0]
		arr[0] = arr[i]
		arr[i] = temp

		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		swap := arr[i]
		arr[i] = arr[largest]
		arr[largest] = swap

		heapify(arr, n, largest)
	}
}
