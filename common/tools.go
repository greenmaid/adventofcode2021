package common

import (
	"log"
	"time"
)

// handy error checker
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
