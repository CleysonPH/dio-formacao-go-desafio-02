package main

import "fmt"

// Você e seus colegas querem criar em formate de código uma antiga brincadeira:
// Ao falar os números sempre que aparecer um mútiplo de 3, o participante deve
// falar "Pin" e nos múltiplos de 5 o participante deve falar "Pan". Então, seu programa
// deve imprimir números de 1 a 100 e nos casos citados, não devem aparecer os números, mas sim,
// o que o participante deve falar.
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Pin")
		}
		if i%5 == 0 {
			fmt.Print("Pan")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
