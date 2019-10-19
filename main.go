package main

import (
    "fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	opts := ParseOptions()
	if err := Generate(opts); err != nil {
        fmt.Println("Error: " + err.Error())
    }
}
