package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var allowedSymbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var specialSymbols = "%*)?@#$~"

func main() {
	result := []byte{}
	rand.Seed(time.Now().Unix())

	length := flag.Int("len", 8, "password length")
	useSpecial := flag.Bool("spec", false, "use special symbols")
	flag.Parse()

	if *useSpecial {
		allowedSymbols = allowedSymbols + specialSymbols
	}

	for i := 0; i < *length; i++ {
		result = append(result, allowedSymbols[rand.Intn(len(allowedSymbols))])
	}
	fmt.Println(string(result))
}
