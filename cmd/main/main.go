package main

import (
	"fmt"

	h "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/html"
)

func main() {
    c := make(chan string) // cria o canal
    html := h.StartHtml()
    
    html.AddField()
    
    go html.RunHtml(c) // chama a função em uma goroutine
    res := <-c // espera pela resposta no canal

    fmt.Println("Olá",res)
    //Receber variáveis

    //Gerar SQL
}






