package metodos

// Algoritmo de ordenamiento Burbuja doble en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func BurbujaDoble(arr []int) []int {
	primero, ultimo := 1, len(arr)-1 // índices del inicio y fin del subarreglo a ordenar
	dir := 0                         // índice del último intercambio realizado

	// el ciclo externo se ejecuta hasta que el subarreglo esté completamente ordenado
	for ultimo >= primero {
		// en el primer ciclo, se recorre el subarreglo de derecha a izquierda
		for i := ultimo; i >= primero; i-- {
			// si el elemento anterior es mayor que el actual, se intercambian
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1] // intercambio
				dir = i                             // se actualiza el índice del último intercambio
			}
		}
		// se actualiza el inicio del subarreglo a ordenar
		primero = dir + 1

		// en el segundo ciclo, se recorre el subarreglo de izquierda a derecha
		for i := primero; i <= ultimo; i++ {
			// si el elemento anterior es mayor que el actual, se intercambian
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1] // intercambio
				dir = i                             // se actualiza el índice del último intercambio
			}
		}
		// se actualiza el fin del subarreglo a ordenar
		ultimo = dir - 1
	}
	return arr // se devuelve el arreglo ordenado
}
