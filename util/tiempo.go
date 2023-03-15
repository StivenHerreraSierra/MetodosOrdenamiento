package util

import (
	"ordenamiento/modelo"
	"time"
)

func MedirTiempo(now time.Time, resultado *modelo.Resultados, prueba *modelo.Prueba) {
	duracion := time.Since(now)
	prueba.Tiempo = duracion
	resultado.Metodo = append(resultado.Metodo, *(prueba))
}
