package metodos

// Algoritmo de ordenamiento Merge en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func MergeSort(arr []int) []int {
	// Caso base: Si el arreglo tiene un solo elemento, retorna el arreglo
	if len(arr) <= 1 {
		return arr
	}

	// Divide el arreglo en dos mitades
	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	// Une las dos mitades ordenadas
	return merge(left, right)
}

// Función que une dos arreglos ordenados
func merge(left, right []int) []int {
	// Crea un arreglo auxiliar para almacenar los elementos ordenados
	result := make([]int, 0)

	// Itera mientras ambos arreglos tengan elementos
	for len(left) > 0 && len(right) > 0 {
		// Compara el primer elemento de cada arreglo
		if left[0] <= right[0] {
			// Si el elemento en left es menor o igual, lo agrega a result y lo remueve de left
			result = append(result, left[0])
			left = left[1:]
		} else {
			// Si el elemento en right es menor, lo agrega a result y lo remueve de right
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// Agrega los elementos restantes de left o right a result
	if len(left) > 0 {
		result = append(result, left...)
	} else {
		result = append(result, right...)
	}

	// Retorna el arreglo ordenado
	return result
}