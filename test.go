package main

import (
	"fmt"
	"github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/model"
	"github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/generator"
)

var Data []interface{}

func main(){
	Data = append(Data, *model.NovaFrase(255, "FraseTeste", true))
	Data = append(Data, *model.NovoTexto("TextoTeste", false))
	Data = append(Data, *model.NovaData(true, "DataTeste"))

	generator.GenerateTableAndFields("dados", Data)
	fmt.Println(Data)
}