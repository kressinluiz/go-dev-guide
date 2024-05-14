package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	txtFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, txtFile)
}
