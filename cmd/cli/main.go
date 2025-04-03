package main

import (
	"fmt"
	"time"
)

type errorPerson struct {
	nome string
}

func (err *errorPerson) Error() string {
	fmt.Println()
	return fmt.Sprintf("%s já está na fila ⚠️\n ", err.nome)
}
func adicionarPessoa(fila *[]string, nome string) error {
	for _, person := range *fila {
		if person == nome {
			return &errorPerson{nome: nome}
		}
	}
	*fila = append(*fila, nome)
	time.Sleep(1*time.Second)
	fmt.Printf("%v adicionado a fila\n", nome)
	return nil
}

func atenderPessoa(fila *[]string) {
	if len(*fila) == 0 {
		fmt.Println()
		time.Sleep(1*time.Second)
		fmt.Println("❗️ Não há pessoas na fila para atendimento")
		return
	}
	atendido := (*fila)[0]
	*fila = (*fila)[1:]
	fmt.Printf("🟢 Atendendo: %s\n", atendido)
	time.Sleep(3*time.Second)
}

func main() {
	fila := make([]string, 0)
	if err := adicionarPessoa(&fila, "João"); err != nil {
		fmt.Println(err)
	}
	if err := adicionarPessoa(&fila, "Joaquim"); err != nil {
		fmt.Println(err)
	}
	if err := adicionarPessoa(&fila, "Maria"); err != nil {
		fmt.Println(err)
	}
	if err := adicionarPessoa(&fila, "Maria"); err != nil {
		fmt.Println(err)
	}
	atenderPessoa(&fila)
	atenderPessoa(&fila)
	atenderPessoa(&fila)
	atenderPessoa(&fila)
}
