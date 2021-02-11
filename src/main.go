package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x = somaRaizQuadrada(x)
	}
	fmt.Fprintf(w, "Code.education Rocks!!")
	fmt.Print("Code.education Rocks!")
}

func somaRaizQuadrada(x float64) float64 {
	x += math.Sqrt(x)
	return x
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
