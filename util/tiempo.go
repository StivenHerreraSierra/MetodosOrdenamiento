package util

import (
	"log"
	"time"
)

func MedirTiempo(now time.Time, texto string) {
	duracion := time.Since(now)

	log.Printf(texto, duracion)
}