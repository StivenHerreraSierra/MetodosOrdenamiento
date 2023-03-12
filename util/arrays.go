package util

import (
	"encoding/json"
	"math/rand"
	"os"
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

func LeerArray(file *os.File) []int {
	decoder := json.NewDecoder(file)

	array := Array{}

	decoder.Decode(&array)

	return array.Datos
}
