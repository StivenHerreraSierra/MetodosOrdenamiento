package metodos

// Algoritmo de ordenamiento Bitonic en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots

// Función que realiza la ordenación bitónica
func BitonicSort(arr []int, asc bool) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	// Se divide el arreglo en dos partes y se ordenan
	m := n / 2
	a1 := BitonicSort(arr[:m], true)
	a2 := BitonicSort(arr[m:], false)
	// Se mezclan ambas partes
	return bitonicMerge(a1, a2, asc)
}

// Función que mezcla dos arreglos bitónicos
func bitonicMerge(a1, a2 []int, asc bool) []int {
	n := len(a1) + len(a2)
	arr := make([]int, n)
	// Se ordenan los arreglos según el ordenamiento bitónico
	for k := 0; k < n; k++ {
		bitonicCompare(a1, a2, k, asc)
		arr[k] = a1[k]
	}
	return arr
}

// Función que compara dos elementos de arreglos bitónicos según el ordenamiento bitónico
func bitonicCompare(a1, a2 []int, k int, asc bool) {
	if asc == (a1[k%len(a1)] > a2[k%len(a2)]) {
		a1[k%len(a1)], a2[k%len(a2)] = a2[k%len(a2)], a1[k%len(a1)]
	}
}
