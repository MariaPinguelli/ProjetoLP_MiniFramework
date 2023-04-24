package main

import (
	//"fmt"
	h "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/html"
    sql "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/sql"
)

func main() {
    html := h.StartHtml()
    
    html.AddTextField("Nome Completo", "nome", true)
    // html.AddTextAreaField("Sobre você", "descricao", false)
    // html.AddSelectField("Ano de nascimento", "dataNascimento", false, []string{"1999", "2000", "2001", "2002", "2003", "2004"})
    html.AddSubmitButton()
    
    html.RunHtml() // chama a função em uma goroutine

    //Gerar SQL
    sql.CreateTableAndInsert("teste", "nome", html, `C:\Users\Maria Fernanda\Desktop\`)
}






