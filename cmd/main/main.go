package main

import (
	//"fmt"

	h "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/html"
)

func main() {
    html := h.StartHtml()
    
    html.AddField()
    html.AddField()
    html.AddField()
    
    html.RunHtml() // chama a função em uma goroutine

    //Receber variáveis

    //Gerar SQL
}






