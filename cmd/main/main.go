package main

import (
	"fmt"
	"os"
	//"os/exec"
	generator "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/generator"
)

func main() {
	/*
		if len(os.Args) < 2 {
			fmt.Println("Por favor declare no comando o arquivo a ser executado!")
			return
		}

		programName := "C:/Users/Maria Fernanda/Desktop/ProjetoLP_MiniFramework/"+os.Args[1]

		cmd := exec.Command("go", "run", programName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Erro ao executar o programa %s: %v\n", programName, err)
	*/

	args := os.Args[1:]
	fmt.Println(args)
	arguments, flag := generator.ProcessInput(args)

	fmt.Println("Argumentos separados:")
	fmt.Println(arguments)

	fmt.Println("Campos HTML gerados:")
	fmt.Println(flag)

}
