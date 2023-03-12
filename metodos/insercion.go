package metodos

// Algoritmo de ordenamiento Inserción en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func InsercionInt(arreglo []int) []int {
    n := len(arreglo)
    for i := 1; i < n; i++ {
        // Selecciona el elemento en la posición i
        valorActual := arreglo[i]
        j := i - 1

        // Mueve los elementos mayores que el valor actual hacia la derecha
        for j >= 0 && arreglo[j] > valorActual {
            arreglo[j+1] = arreglo[j]
            j = j - 1
        }

        // Inserta el valor actual en su posición correcta
        arreglo[j+1] = valorActual
    }

    return arreglo
}

// Algoritmo de ordenamiento Inserción en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func InsercionFloat(arreglo []float64) []float64 {
    n := len(arreglo)
    for i := 1; i < n; i++ {
        // Selecciona el elemento en la posición i
        valorActual := arreglo[i]
        j := i - 1

        // Mueve los elementos mayores que el valor actual hacia la derecha
        for j >= 0 && arreglo[j] > valorActual {
            arreglo[j+1] = arreglo[j]
            j = j - 1
        }

        // Inserta el valor actual en su posición correcta
        arreglo[j+1] = valorActual
    }

    return arreglo
}