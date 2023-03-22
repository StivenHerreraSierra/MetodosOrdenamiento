package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"ordenamiento/modelo"
	"os"
	"sort"
	"strconv"
	"time"
)

type Array struct {
	Datos []int
}

func GenerarArray(tamano int) []int {
	rand.Seed(time.Now().UnixNano())

	datos := make([]int, tamano)

	for i := 0; i < tamano; i++ {
		datos[i] = rand.Intn(999999999)
	}

	array := Array{Datos: datos}

	return array.Datos
}

func EscribirArray(file *os.File, elementos []int) {
	encoder := json.NewEncoder(file)

	array := Array{Datos: elementos}

	encoder.Encode(array)
}

func EscribirResultado(resultados modelo.Resultados) {
	file, _ := os.Create("resultados/prueba-" + strconv.Itoa(resultados.Tamano) + ".json")

	encoder := json.NewEncoder(file)

	encoder.Encode(resultados)
}

func LeerResultado(nombre string) modelo.Resultados {
	file, _ := os.Open("resultados/" + nombre)

	decoder := json.NewDecoder(file)

	resultados := modelo.Resultados{}

	decoder.Decode(&resultados)

	return resultados
}

func LeerArray(file *os.File) []int {
	decoder := json.NewDecoder(file)

	array := Array{}

	decoder.Decode(&array)

	return array.Datos
}

func OrdenarResultados(resultados []modelo.Resultados) []modelo.Resultados {
	ordenado := make([]modelo.Resultados, len(resultados))
	copy(ordenado, resultados)

	for res := 0; res < len(resultados); res++ {
		pruebas := resultados[res].Metodo
		sort.Slice(pruebas, func(i, j int) bool {
			return pruebas[i].Tiempo > pruebas[j].Tiempo
		})

		ordenado[res].Metodo = pruebas
	}

	return ordenado
}

func GenerarTabla(resultados []modelo.Resultados) {
	file, _ := os.Create("resultados/ordenados.txt")
	defer file.Close()

	for _, resultado := range resultados {
		file.WriteString(fmt.Sprintf("Prueba con arreglo de %d elementos\n\n", resultado.Tamano))
		for pos, prueba := range resultado.Metodo {
			file.WriteString(fmt.Sprintf("%d. %s: %s\n", (pos + 1), prueba.Metodo, prueba.Tiempo))
		}
		file.WriteString("======================\n\n")
	}
}
