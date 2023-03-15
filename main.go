package main

import (
	"log"
	"ordenamiento/metodos"
	"ordenamiento/modelo"
	"ordenamiento/util"
	"os"
	"strconv"
	"time"
)

func main() {
	iniciarPrueba()

	resultados := obtenerResultados()

	resultados = util.OrdenarResultados(resultados)

	util.GenerarTabla(resultados)
}

func iniciarPrueba() {
	var resultados []modelo.Resultados

	for i, tamano := range []int{100, 1000, 10000, 50000, 100000, 500000} {
		resultados = append(resultados, modelo.Resultados{Tamano: tamano})

		file2, _ := os.Open("arreglos/Array_" + strconv.Itoa(tamano) + ".json")
		array := util.LeerArray(file2)

		log.Printf("\nPrueba método Burbuja con %d elementos", tamano)
		testBurbuja(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Burbuja Doble con %d elementos", tamano)
		testBurbujaDoble(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Seleccion con %d elementos", tamano)
		testSeleccion(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Insertion con %d elementos", tamano)
		testInsercion(array, tamano, &resultados[i])

		log.Printf("\nPrueba método RecursiveInsertion con %d elementos", tamano)
		testRecursiveInsertion(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Shell con %d elementos", tamano)
		testShell(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Bucket con %d elementos", tamano)
		testBucket(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Merge con %d elementos", tamano)
		testMerge(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Quick con %d elementos", tamano)
		testQuick(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Bitonic con %d elementos", tamano)
		testBitonic(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Gnome con %d elementos", tamano)
		testGnome(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Binary Insertion con %d elementos", tamano)
		testBinaryInsertion(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Strand con %d elementos", tamano)
		testStrand(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Radix con %d elementos", tamano)
		testRadix(array, tamano, &resultados[i])

		log.Printf("\nPrueba método Heap con %d elementos", tamano)
		testHeap(array, tamano, &resultados[i])

		if tamano <= 10000 {
			log.Printf("\nPrueba método Stooge con %d elementos", tamano)
			testStooge(array, tamano, &resultados[i])
		}

		file2.Close()

		util.EscribirResultado(resultados[i])
	}
}

func testBurbuja(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Burbuja"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.Burbuja(arrayCopy)
}

func testBurbujaDoble(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Burbuja doble"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.BurbujaDoble(arrayCopy)
}

func testSeleccion(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Seleccion"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.Seleccion(arrayCopy)
}

func testInsercion(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Insercion"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.InsercionInt(arrayCopy)
}

func testRecursiveInsertion(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Recursive insercion"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.RecursiveInsertionSort(arrayCopy)
}

func testShell(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Shell"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.ShellSort(arrayCopy)
}

func testBucket(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Bucket"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.BucketSortInt(arrayCopy)
}

func testMerge(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Merge"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.MergeSort(arrayCopy)
}

func testQuick(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Quick"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.QuickSort(arrayCopy)
}

func testStooge(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Stooge"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.StoogeSort(arrayCopy)
}

func testHeap(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Heap"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.HeapSort(arrayCopy)
}

func testBitonic(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Bitonic"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.BitonicSort(arrayCopy, true)
}

func testGnome(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Gnome"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.GnomeSort(arrayCopy)
}

func testBinaryInsertion(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Insertion"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.BinaryInsertionSort(arrayCopy)
}

func testStrand(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Strand"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.StrandSort(arrayCopy)
}

func testRadix(array []int, tamano int, resultado *modelo.Resultados) {
	arrayCopy := make([]int, tamano)
	copy(arrayCopy, array)

	prueba := &modelo.Prueba{Metodo: "Radix"}

	now := time.Now()
	defer util.MedirTiempo(now, resultado, prueba)

	metodos.RadixSort(arrayCopy)
}

func obtenerResultados() []modelo.Resultados {
	var resultados []modelo.Resultados
	for _, tamano := range []int{100, 1000, 10000, 50000, 100000, 500000} {
		resultados = append(resultados, util.LeerResultado("prueba-"+strconv.Itoa(tamano)+".json"))
	}

	return resultados
}
