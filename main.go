package main

import (
	"log"
	"ordenamiento/metodos"
	"ordenamiento/util"
	"os"
	"strconv"
	"time"
)

func main() {
	iniciarPruebas()
}

func iniciarPruebas() {
	log.Println("\nInicia la prueba con 100 elementos")
	testMetodos(100)

	log.Println("\nInicia la prueba con 1.000 elementos")
	testMetodos(1000)

	log.Println("\nInicia la prueba con 10.000 elementos")
	testMetodos(10000)

	log.Println("\nInicia la prueba con 50.000 elementos")
	testMetodos(50000)

	log.Println("\nInicia la prueba con 100.000 elementos")
	testMetodos(100000)

	log.Println("\nInicia la prueba con 500.000 elementos")
	testMetodos(500000)
}

func testMetodos(tamanoArray int) {
	file2, _ := os.Open("arreglos/Array_" + strconv.Itoa(tamanoArray) + ".json")
	array := util.LeerArray(file2)
	defer file2.Close()

	testBurbuja(array)
	testBurbujaDoble(array)
	testSeleccion(array)
	testInsercion(array)
	testRecursiveInsertion(array)
	testShell(array)
	testBucket(array)
	testMerge(array)
	testQuick(array)
	testStooge(array)
	testHeap(array)
	testBitonic(array)
	testGnome(array)
	testBinaryInsertion(array)
	testStrand(array)
	testRadix(array)
}

func testBurbuja(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Burbuja con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.Burbuja(arrayCopy)
}

func testBurbujaDoble(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Burbuja Doble con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.BurbujaDoble(arrayCopy)
}

func testSeleccion(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Seleccion con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.Seleccion(arrayCopy)
}

func testInsercion(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Insercion con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.InsercionInt(arrayCopy)
}

func testRecursiveInsertion(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Recursive Insertion con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.RecursiveInsertionSort(arrayCopy)
}

func testShell(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Shell con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.ShellSort(arrayCopy)
}

func testBucket(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Bucket con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.BucketSortInt(arrayCopy)
}

func testMerge(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Merge con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.MergeSort(arrayCopy)
}

func testQuick(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Quick con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.QuickSort(arrayCopy)
}

func testStooge(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Stooge con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.StoogeSort(arrayCopy)
}

func testHeap(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Heap con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.StoogeSort(arrayCopy)
}

func testBitonic(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Bitonic con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.BitonicSort(arrayCopy, true)
}

func testGnome(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Gnome con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.GnomeSort(arrayCopy)
}

func testBinaryInsertion(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Binary Insertion con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.BinaryInsertionSort(arrayCopy)
}

func testStrand(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Strand con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.StrandSort(arrayCopy)
}

func testRadix(array []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	now := time.Now()
	defer util.MedirTiempo(now, "Tiempo empleado método Radix con arreglo de "+strconv.Itoa(len(array))+" elementos: %s\n")

	metodos.RadixSort(arrayCopy)
}
