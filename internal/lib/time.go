package lib

import (
	"log"
	"time"
)

func LogDuration(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("OK (%vms)\n", elapsed.Milliseconds())
}
