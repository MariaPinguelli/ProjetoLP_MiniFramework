package main

import (
	//"fmt"
	informacoes "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/informacoes"
	model "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/model"
	generator "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/generator"

	"os"
	//"os/exec"
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

	// Processa os argumentos
	input := generator.ProcessInput(args)

	// Cria as informações a partir dos argumentos
	nome := model.NovaFrase(50, input[0])
	bio := model.NovoTexto(input[1])
	ultimaAtualizacao := model.NovaData()

	informacao := informacoes.NovaInformacao(nome, bio, ultimaAtualizacao)

	// Gera o HTML com as informações
	generator.GenerateHTMLFields([]interface{}{informacao.Nome, informacao.Biografia, informacao.UltimaAtualizacao})

}
