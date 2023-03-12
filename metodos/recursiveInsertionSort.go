package metodos

// Algoritmo de ordenamiento Recursive Insertion en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func RecursiveInsertionSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	// ordenamos los primeros n-1 elementos
	RecursiveInsertionSort(arr[:n-1])

	// insertamos el último elemento en la posición correcta
	last := arr[n-1]
	j := n - 2
	for j >= 0 && arr[j] > last {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = last

	return arr
}