package metodos

import (
	"math"
)

// Algoritmo de ordenamiento Radix en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots

// Función para obtener el dígito en la posición i del número n
func obtenerDigito(n int, i int) int {
	return (n / int(math.Pow(10, float64(i)))) % 10
}

// Función para obtener el número de dígitos del número n
func obtenerNumeroDigitos(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(n)))) + 1)
}

// Función para ordenar un arreglo de números utilizando Radix Sort
func RadixSort(arr []int) []int {
	// Se obtiene el número de dígitos del número más grande en el arreglo
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	numDigitos := obtenerNumeroDigitos(max)

	// Se crea una lista de listas para cada dígito (0-9)
	buckets := make([][]int, 10)
	for i := 0; i < 10; i++ {
		buckets[i] = make([]int, 0)
	}

	// Se realiza el sorting por dígitos, comenzando por el dígito menos significativo
	for i := 0; i < numDigitos; i++ {
		// Se distribuyen los números en los buckets correspondientes según su dígito en la posición i
		for j := 0; j < len(arr); j++ {
			digito := obtenerDigito(arr[j], i)
			buckets[digito] = append(buckets[digito], arr[j])
		}

		// Se vacían los buckets en el arreglo original en el orden correcto
		idx := 0
		for j := 0; j < 10; j++ {
			for k := 0; k < len(buckets[j]); k++ {
				arr[idx] = buckets[j][k]
				idx++
			}
			// Se vacía el bucket
			buckets[j] = make([]int, 0)
		}
	}

	// Devuelve el arreglo ya ordenado
	return arr
}
