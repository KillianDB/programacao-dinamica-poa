package main

import (
    "fmt"
    "strconv"
)

var Mem = make(map[int]int)

func main() {
    var input string
    fmt.Print("Digite um número inteiro positivo: ")
    fmt.Scanln(&input)

    n, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Erro: o valor digitado deve ser um número inteiro válido")
        return
    }

    if n <= 0 {
        fmt.Println("Erro: o número deve ser positivo")
        return
    }

    fmt.Println("1 resultado para", n, ":", diagnosticadorDeProblemaRecursivo(n))
    fmt.Println("2 resultado para", n, ":", diagnosticadorDeProblemaComMemorizacao(n))
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