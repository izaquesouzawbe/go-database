package main_aux

import (
	"fmt"
	"time"
)

var runtimeVar time.Time

func RuntimeStarted() {
	runtimeVar = time.Now()
}

func RuntimeDone() {
	fmt.Printf("Runtime: %s\n", time.Since(runtimeVar))
	fmt.Println("")
}
