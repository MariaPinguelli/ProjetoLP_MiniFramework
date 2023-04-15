package main

import (
	"fmt"
	field "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/field"
)

func main() {
	//Um pequeno campo de texto simples
	f := field.Field{ Name: "Testerson", Type: "Teste", Required: false}
	f2 := field.Field{ Name: "Testerson o Segundo", Type: "Teste Dois", Required: false}
	fmt.Println("Hello: ", f.Name)
	fmt.Println("Hello: ", f2.Name)
}
