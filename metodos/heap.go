package metodos

// Algoritmo de ordenamiento Burbuja doble en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func HeapSort(arr []int) []int {
	// Primero se construye el heap máximo
	buildMaxHeap(arr)

	// Se va extrayendo el valor máximo y se va colocando en su posición correspondiente en el arreglo
	for i := len(arr) - 1; i > 0; i-- {
		// Intercambio el valor máximo del heap con el último valor del heap
		arr[0], arr[i] = arr[i], arr[0]

		// Restauro la propiedad de max-heap del arreglo de 0 a i-1
		maxHeapify(arr[:i], 0)
	}

	// Devuelve el arreglo ya ordenado
	return arr
}

// Función para construir el heap máximo
func buildMaxHeap(arr []int) {
	// El heap máximo comienza desde la mitad del arreglo hacia atrás
	for i := len(arr) / 2; i >= 0; i-- {
		maxHeapify(arr, i)
	}
}

// Función para asegurar la propiedad de max-heap del arreglo
func maxHeapify(arr []int, i int) {
	// Posición del hijo izquierdo en el heap
	l := 2*i + 1
	// Posición del hijo derecho en el heap
	r := 2*i + 2
	// Posición del elemento máximo en el heap
	max := i

	// Si el hijo izquierdo es mayor que el padre, lo marca como máximo
	if l < len(arr) && arr[l] > arr[max] {
		max = l
	}
	// Si el hijo derecho es mayor que el padre, lo marca como máximo
	if r < len(arr) && arr[r] > arr[max] {
		max = r
	}

	// Si el elemento máximo no es el padre, intercambia el padre con el hijo máximo y llama a maxHeapify de nuevo en el hijo máximo
	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		maxHeapify(arr, max)
	}
}