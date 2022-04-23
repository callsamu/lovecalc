package main

import (
	"flag"
	"fmt"
	"hash/fnv"

	"github.com/callsamu/lovecalc/pkg/core"
)

func main() {
	fmt.Printf("\n=========[ CALCULADORA DO AMOR <3 ]========\n\n")
	first := flag.String("PrimeiroNome", "", "Primeiro nome a ser inserido na calculadora")
	second := flag.String("SegundoNome", "", "Segundo nome a ser inserido na calculadora")
	flag.Parse()

	if *first == "" || *second == "" {
		fmt.Println("Erro: Ambos PrimeiroNome e SegundoNome precisam ser declarados")
		fmt.Println("=============================================")
		return
	}

	calculator := core.HashCalculator{Hash: fnv.New64()}
	result := 100 * calculator.Compute(*first, *second)

	fmt.Printf(">>> Compatibilidade entre %s e %s: %4f%%\n", *first, *second, result)
	fmt.Println("\n=============================================")
}
