package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	allowedSymbols := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := []byte{}
	rand.Seed(time.Now().Unix())

	length := flag.Int("len", 8, "password length")
	flag.Parse()

	for i := 0; i < *length; i++ {
		result = append(result, allowedSymbols[rand.Intn(len(allowedSymbols))])
	}
	fmt.Println(string(result))
}
