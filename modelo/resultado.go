package modelo

import "time"

type Resultados struct {
	Tamano int
	Metodo []Prueba
}

type Prueba struct {
	Metodo string
	Tiempo time.Duration
}
