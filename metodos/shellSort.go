package metodos

// Algoritmo de ordenamiento Shell en Golang
// Código proporcionado por OpenAI, 2023
// Recuperado de https://github.com/openai/chat-bots
func ShellSort(a []int) []int {
    // Calcula el incremento inicial
    for incr := len(a) / 2; incr > 0; incr /= 2 {
        // Realiza la ordenación con el incremento actual
        for i := incr; i < len(a); i++ {
            j := i - incr
            for j >= 0 {
                // Si el elemento actual es mayor que el de la posición actual + incremento
                // intercambia los elementos
                if a[j] > a[j+incr] {
                    a[j], a[j+incr] = a[j+incr], a[j]
                    j -= incr
                } else {
                    break
                }
            }
        }
    }
    // Retorna el slice ordenado
    return a
}
