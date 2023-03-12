package metodos

// Algoritmo de ordenamiento Binary Insertion en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots

// Función que realiza el ordenamiento Binary Insertion Sort
func BinaryInsertionSort(arr []int) {
	// Recorre el arreglo desde la segunda posición
	for i := 1; i < len(arr); i++ {
		// Guarda temporalmente el valor del elemento en la posición i
		temp := arr[i]

		// Inicializa los índices para la búsqueda binaria
		left := 0
		right := i - 1

		// Realiza la búsqueda binaria para encontrar la posición adecuada donde insertar el elemento
		for left <= right {
			mid := (left + right) / 2

			if temp < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		// Desplaza los elementos hacia la derecha para abrir espacio para el nuevo elemento
		for j := i - 1; j >= left; j-- {
			arr[j+1] = arr[j]
		}

		// Inserta el elemento en la posición correcta
		arr[left] = temp
	}
}