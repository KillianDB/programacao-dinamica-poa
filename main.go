package main

import "fmt"

var Mem = make(map[int]int)

func main() {
	fmt.Println("1 resultado para 1777: ", diagnosticadorDeProblemaRecursivo(1777))
    fmt.Println("2 resultado para 1777: ", diagnosticadorDeProblemaComMemorizacao(1777))
}

//? Operações possíveis: -1, x/2 (exata), x/3 (exata)
//* Foco: número mínimo de operações para chegar a 1 ou em um múltiplo de 7

//# Versão recursiva simples
func diagnosticadorDeProblemaRecursivo(n int) int {
	fmt.Printf("Analisando %d\n", n)

	if n == 0 {
		return 0
	}

    if n == 1 || n%7 == 0 {
        fmt.Printf("Encontrado %d, retornando 1\n", n)
        return 1
    }

    return diagnosticadorDeProblemaRecursivo(n-1) + 1
}

//# Versão recursiva com memorização
func diagnosticadorDeProblemaComMemorizacao(n int) int {
	fmt.Printf("Analisando %d\n", n)

	if n == 0 {
		return 0
	}

	if n == 1 || n%7 == 0 {
		fmt.Printf("Encontrado %d, retornando 1\n", n)
		return 1
	}

	if val, ok := Mem[n]; ok {
		return val
	}

	Mem[n] = diagnosticadorDeProblemaComMemorizacao(n-1) + 1

	return Mem[n]
}

// //# Versão não-recursiva
// func diagnosticadorDeProblemaNaoRecursivo(n int) int {
// }