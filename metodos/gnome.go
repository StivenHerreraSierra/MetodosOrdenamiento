package metodos

// Algoritmo de ordenamiento Gnome en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots

// Función de ordenamiento Gnome Sort
func GnomeSort(arr []int) []int {
	// Índice del arreglo
	i := 1

	// Mientras el índice sea menor al tamaño del arreglo
	for i < len(arr) {
		// Si el elemento actual es mayor o igual que el anterior, avanzamos al siguiente elemento
		if arr[i] >= arr[i-1] {
			i++
		} else {
			// Si el elemento actual es menor que el anterior, intercambiamos los elementos y retrocedemos en el índice
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 {
				i--
			}
		}
	}

	// Devuelve el arreglo ya ordenado
	return arr
}
