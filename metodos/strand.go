package metodos

// Algoritmo de ordenamiento Strand en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots

// Función para ordenar un slice de enteros utilizando Strand Sort
func StrandSort(arr []int) []int {
	// Si el slice tiene longitud menor o igual a 1, ya está ordenado
	if len(arr) <= 1 {
		return arr
	}

	// Se inicializa una lista vacía para la sublista ordenada
	sortedList := []int{}
	// Se inicializa una lista con el primer elemento del slice como sublista actual
	sublist := []int{arr[0]}
	// Se itera sobre el slice desde la segunda posición
	for i := 1; i < len(arr); i++ {
		// Se obtiene el elemento actual del slice
		elem := arr[i]
		// Si el elemento actual es mayor o igual al último elemento de la sublista actual, se agrega a la sublista
		if elem >= sublist[len(sublist)-1] {
			sublist = append(sublist, elem)
		} else {
			// Si el elemento actual es menor al último elemento de la sublista actual, se realiza la fusión de sublists
			sortedList = mergeStrand(sortedList, sublist)
			// Se crea una nueva sublista con el elemento actual
			sublist = []int{elem}
		}
	}
	// Se realiza la fusión de la sublista actual y la lista ordenada
	sortedList = mergeStrand(sortedList, sublist)

	// Se devuelve la lista ordenada
	return sortedList
}

// Función para fusionar dos sublists ordenadas
func mergeStrand(list1, list2 []int) []int {
	// Se inicializa una lista vacía para la lista ordenada resultante de la fusión
	mergedList := []int{}
	// Se inicializan dos índices para recorrer las sublists
	i, j := 0, 0
	// Se itera mientras haya elementos en ambas sublists
	for i < len(list1) && j < len(list2) {
		// Se compara el primer elemento de cada sublist y se agrega el menor a la lista ordenada resultante de la fusión
		if list1[i] < list2[j] {
			mergedList = append(mergedList, list1[i])
			i++
		} else {
			mergedList = append(mergedList, list2[j])
			j++
		}
	}
	// Se agregan los elementos restantes de la sublist que aún tiene elementos
	mergedList = append(mergedList, list1[i:]...)
	mergedList = append(mergedList, list2[j:]...)
	// Se devuelve la lista ordenada resultante de la fusión
	return mergedList
}
