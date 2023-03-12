package main

import (
	"ordenamiento/metodos"
	"ordenamiento/util"
	"os"
	"strconv"
	"time"
)

func main() {
	//testMetodos(100)
	//testMetodos(1000)
	//testMetodos(10000)
	//testMetodos(100000)
	//testMetodos(1000000)
	testMetodos(10000000)
}

func testMetodos(tamanoArray int) {
	file2, _ := os.Open("arreglos/Array_" + strconv.Itoa(tamanoArray) + ".json")
	array := util.LeerArray(file2)
	defer file2.Close()

	//testBurbuja(array)
	//testBurbujaDoble(array)
	testSeleccion(array)
	testInsercion(array)
	testRecursiveInsertion(array)
	testShell(array)
	testBucket(array)
	testMerge(array)
	testQuick(array)
	//testStooge(array)
	//testHeap(array)
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
