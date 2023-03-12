package metodos

// Algoritmo de ordenamiento Stooge en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func StoogeSort(arr []int) []int {
	// Función recursiva que ordena el arreglo.

	sort(arr, 0, len(arr)-1) // Llamada a la función recursiva.

	return arr // Retorna el arreglo ordenado.
}

func sort(arr []int, l int, r int) {
	if arr[l] > arr[r] {
		arr[l], arr[r] = arr[r], arr[l] // Intercambia los valores.
	}
	if r-l+1 > 2 {
		t := (r - l + 1) / 3 // Define el tamaño del tercio.
		sort(arr, l, r-t)    // Ordena el primer tercio.
		sort(arr, l+t, r)    // Ordena el último tercio.
		sort(arr, l, r-t)    // Ordena nuevamente el primer tercio.
	}
}
