package metodos

// Algoritmo de ordenamiento Bitonic en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots

// BitonicSort ordena un arreglo utilizando el algoritmo de Bitonic Sort
func BitonicSort(arr []int, asc bool) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	// Crear un arreglo bitónico
	bitonicSort(arr, 0, n, asc)

	return arr
}

// Función auxiliar para crear un arreglo bitónico
func bitonicSort(arr []int, low, cnt int, asc bool) {
	if cnt > 1 {
		mid := cnt / 2
		bitonicSort(arr, low, mid, !asc)
		bitonicSort(arr, low+mid, cnt-mid, asc)
		bitonicMerge(arr, low, cnt, asc)
	}
}

// Función para mezclar dos arreglos bitónicos
func bitonicMerge(arr []int, low, cnt int, asc bool) {
	if cnt > 1 {
		k := cnt / 2
		for i := low; i < low+k; i++ {
			bitonicCompare(arr, i, i+k, asc)
		}
		bitonicMerge(arr, low, k, asc)
		bitonicMerge(arr, low+k, k, asc)
	}
}

// Función para comparar y ordenar dos elementos de un arreglo en orden bitónico
func bitonicCompare(arr []int, i, j int, asc bool) {
	if (arr[i] > arr[j] && asc) || (arr[i] < arr[j] && !asc) {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
