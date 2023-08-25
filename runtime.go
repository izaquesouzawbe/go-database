package main

import (
	"fmt"
	"time"
)

var runtimeVar time.Time

func runtimeStarted() {
	runtimeVar = time.Now()
}

func runtimeDone() {
	fmt.Printf("Runtime: %s\n", time.Since(runtimeVar))
	fmt.Println("")
}
