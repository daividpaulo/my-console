package main

import (
	"fmt"
	"log"
	"my-console/app"
	"os"
)

func main() {
	fmt.Println("")
	fmt.Println("Executando")
	fmt.Println("")

	aplicacao := app.NewClient()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("")
	fmt.Println("Operação concluida!")
	fmt.Println("")
}
