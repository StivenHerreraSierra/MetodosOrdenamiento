package metodos

// Algoritmo de ordenamiento Bucket en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func BucketSortInt(a []int) []int {
	// Obtener el valor máximo del arreglo
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	// Inicializar los buckets
	n := len(a)
	buckets := make([][]int, n)
	for i := range buckets {
		buckets[i] = []int{}
	}

	// Distribuir los elementos del arreglo en los buckets
	for _, v := range a {
		idx := int(v / (max + 1) * (n - 1))
		buckets[idx] = append(buckets[idx], v)
	}

	// Ordenar cada bucket y unirlos en un solo arreglo
	result := []int{}
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			// Ordenar el bucket utilizando Insertion Sort
			InsercionInt(bucket)

			// Agregar los elementos ordenados al resultado final
			result = append(result, bucket...)
		}
	}

	return result
}

// Algoritmo de ordenamiento Bucket en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func BucketSortFloat(a []float64) []float64 {
	// Obtener el valor máximo del arreglo
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	// Inicializar los buckets
	n := len(a)
	buckets := make([][]float64, n)
	for i := range buckets {
		buckets[i] = []float64{}
	}

	// Distribuir los elementos del arreglo en los buckets
	for _, v := range a {
		idx := int(v / (max + 1) * float64(n-1))
		buckets[idx] = append(buckets[idx], v)
	}

	// Ordenar cada bucket y unirlos en un solo arreglo
	result := []float64{}
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			// Ordenar el bucket utilizando Insertion Sort
			InsercionFloat(bucket)

			// Agregar los elementos ordenados al resultado final
			result = append(result, bucket...)
		}
	}

	return result
}
