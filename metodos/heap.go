package metodos

// Algoritmo de ordenamiento Heap en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func OrdenarHeapSort(arr []int) []int {
	// Se construye el Heap máximo
	ConstruirHeapMaximo(arr)

	// Se va extrayendo el valor máximo y se va colocando en su posición correspondiente en el arreglo
	for i := len(arr) - 1; i > 0; i-- {
		// Se intercambia el valor máximo del heap con el último valor del heap
		arr[0], arr[i] = arr[i], arr[0]

		// Se restaura la propiedad de max-heap del arreglo de 0 a i-1
		RestaurarMaxHeap(arr[:i], 0)
	}

	// Devuelve el arreglo ya ordenado
	return arr
}

// Función para construir el Heap máximo
func ConstruirHeapMaximo(arr []int) {
	// El Heap máximo comienza desde la mitad del arreglo hacia atrás
	for i := len(arr) / 2; i >= 0; i-- {
		RestaurarMaxHeap(arr, i)
	}
}

// Función para asegurar la propiedad de max-heap del arreglo
func RestaurarMaxHeap(arr []int, i int) {
	// Posición del hijo izquierdo en el heap
	izq := 2*i + 1
	// Posición del hijo derecho en el heap
	der := 2*i + 2
	// Posición del elemento máximo en el heap
	max := i

	// Si el hijo izquierdo es mayor que el padre, lo marca como máximo
	if izq < len(arr) && arr[izq] > arr[max] {
		max = izq
	}

	// Si el hijo derecho es mayor que el padre, lo marca como máximo
	if der < len(arr) && arr[der] > arr[max] {
		max = der
	}

	// Si el elemento máximo no es el padre, intercambia el padre con el hijo máximo y llama a RestaurarMaxHeap de nuevo en el hijo máximo
	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		RestaurarMaxHeap(arr, max)
	}
}
