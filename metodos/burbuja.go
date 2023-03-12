package metodos

// Algoritmo de ordenamiento Burbuja en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func Burbuja(arr []int) []int {
	n := len(arr)
	// Itera a través del arreglo
	for i := 0; i < n-1; i++ {
		// Itera a través de los elementos no ordenados
		for j := 0; j < n-1; j++ {
			// Si el elemento actual es mayor que el siguiente, intercámbialos
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}