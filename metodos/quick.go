package metodos

// Algoritmo de ordenamiento Quick en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func QuickSort(arr []int) []int {
    // Si el arreglo tiene menos de 2 elementos, entonces ya está ordenado
    if len(arr) < 2 {
        return arr
    }

    // Seleccionamos el pivote como el último elemento del arreglo
    pivot := arr[len(arr)-1]

    // Inicializamos los arreglos que van a contener los elementos menores y mayores al pivote
    left := []int{}
    right := []int{}

    // Recorremos el arreglo y vamos colocando los elementos menores al pivote en el arreglo "left"
    // y los elementos mayores en el arreglo "right"
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] <= pivot {
            left = append(left, arr[i])
        } else {
            right = append(right, arr[i])
        }
    }

    // Ordenamos recursivamente los arreglos "left" y "right"
    left = QuickSort(left)
    right = QuickSort(right)

    // Concatenamos los arreglos "left", el pivote y "right" en un solo arreglo
    return append(append(left, pivot), right...)
}
