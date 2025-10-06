package main

import (
	"fmt"
	"strconv"
)

var (
	Mem map[int]int
	Op map[int]string
)

func main() {
	var input string
    fmt.Print("Digite um número inteiro positivo: ")
    fmt.Scanln(&input)

    n, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Erro: o valor digitado deve ser um número inteiro válido")
        return
    }

    if n < 0 {
        fmt.Println("Erro: o número deve ser positivo")
        return
    }

	executarDiagnosticos(n)
}

func executarDiagnosticos(n int) {
	fmt.Println("\n 1º resultado - Função Recursiva Simples:")
	total, _ := diagnosticadorDeProblemaRecursivo(n)
	fmt.Printf("Resultado ótimo: %d - %s\n", total, imprimirCaminho(n, Op))

	fmt.Println("\n 2º resultado - Função Recursiva com Memorização:")
	Mem = make(map[int]int)
	Op = make(map[int]string)
	total, _ = diagnosticadorDeProblemaComMemorizacao(n)
	fmt.Printf("Resultado ótimo: %d - %s\n", total, imprimirCaminho(n, Op))
	
	fmt.Println("\n 3º resultado - Função Não Recursiva:")
	total = diagnosticadorDeProblemaNaoRecursivo(n)
	fmt.Printf("Resultado ótimo: %d - %s\n", total, imprimirCaminho(n, Op))
}

//? Operações possíveis: -1, x/2 (exata), x/3 (exata)
//* Foco: número mínimo de operações(das operações possíveis) para chegar a 1 ou em um múltiplo de 7

//# Versão recursiva simples
func diagnosticadorDeProblemaRecursivo(n int) (int, string) {

	if n == 0 || n == 1 || n%7 == 0 {
		return 0, ""
	}

	melhor := 9999999
	opLocal := ""

	op1, _ := diagnosticadorDeProblemaRecursivo(n - 1)
	if op1+1 < melhor {
		melhor = op1 + 1
		opLocal = "-1"
	}

	if n%2 == 0 {
		op2, _ := diagnosticadorDeProblemaRecursivo(n / 2)
		if op2+1 < melhor {
			melhor = op2 + 1
			opLocal = "/2"
		}
	}

	if n%3 == 0 {
		op3, _ := diagnosticadorDeProblemaRecursivo(n / 3)
		if op3+1 < melhor {
			melhor = op3 + 1
			opLocal = "/3"
		}
	}

	if Op == nil {
		Op = make(map[int]string)
	}

	Op[n] = opLocal
	return melhor, opLocal
}

//# Versão recursiva com memorização
func diagnosticadorDeProblemaComMemorizacao(n int) (int, string) {

	if n == 0 || n == 1 || n%7 == 0 {
		return 0, ""
	}

	if Mem == nil {
		Mem = make(map[int]int)
		Op = make(map[int]string)
	}

	if val, ok := Mem[n]; ok {
		return val, Op[n]
	}

	melhor := 9999999
	opLocal := ""

	op1, _ := diagnosticadorDeProblemaComMemorizacao(n - 1)
	if op1+1 < melhor {
		melhor = op1 + 1
		opLocal = "-1"
	}

	if n%2 == 0 {
		op2, _ := diagnosticadorDeProblemaComMemorizacao(n / 2)
		if op2+1 < melhor {
			melhor = op2 + 1
			opLocal = "/2"
		}
	}

	if n%3 == 0 {
		op3, _ := diagnosticadorDeProblemaComMemorizacao(n / 3)
		if op3+1 < melhor {
			melhor = op3 + 1
			opLocal = "/3"
		}
	}

	Mem[n] = melhor
	Op[n] = opLocal
	return melhor, opLocal
}

//# Versão não-recursiva
func diagnosticadorDeProblemaNaoRecursivo(n int) int {

	if n == 0 || n == 1 || n%7 == 0 {
		return 0
	}

	total := make([]int, n+1)
	opLocal := make(map[int]string)

	for i := 2; i <= n; i++ {
		if i%7 == 0 {
			total[i] = 0
			continue
		}
		total[i] = total[i-1] + 1
		opLocal[i] = "-1"

		if i%2 == 0 && 1+total[i/2] < total[i] {
			total[i] = 1 + total[i/2]
			opLocal[i] = "/2"
		}
		if i%3 == 0 && 1+total[i/3] < total[i] {
			total[i] = 1 + total[i/3]
			opLocal[i] = "/3"
		}
	}
	Op = opLocal

	return total[n]
}

func imprimirCaminho(n int, operacoes map[int]string) string {
	var ops []string
	for n != 1 && n%7 != 0 {
		opLocal := operacoes[n]
		ops = append(ops, opLocal)
		switch opLocal {
		case "-1":
			n = n - 1
		case "/2":
			n = n / 2
		case "/3":
			n = n / 3
		default:
			return ""
		}
	}
	fmt.Printf("Parou em %d \n", n)
	return fmt.Sprintf("Operações: %v\n", ops)
}
