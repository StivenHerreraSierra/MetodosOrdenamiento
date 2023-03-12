package metodos

// Algoritmo de ordenamiento Selección en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func Seleccion(array []int) []int {
	n := len(array)
	
	// Recorrer todo el array
	for i := 0; i < n-1; i++ {
		// Encontrar el elemento mínimo en el array sin ordenar
		min := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		// Intercambiar el elemento mínimo con el primer elemento sin ordenar
		temp := array[min]
		array[min] = array[i]
		array[i] = temp
	}

	return array
}
