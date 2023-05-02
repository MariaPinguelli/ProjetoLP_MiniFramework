package generator

import (
	"fmt"
	"strings"

	"github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/html"
	model "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/model"
	"github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/sql"
)

func ProcessInput(args []string) []string {
	var flags []string
	var arguments []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		} else {
			arguments = append(arguments, arg)
		}
	}

	return arguments
}

func GenerateForm(models []interface{}) {
	htmlForm := html.StartHtml()
	for _, m := range models {
		// Checar o tipo do modelo e gerar um campo HTML apropriado
		switch m.(type) {
		case model.Frase:
			frase := m.(model.Frase)
			htmlForm.AddTextField("Frase", strings.ToLower(fmt.Sprintf("%T", frase)), true)
		case model.Texto:
			texto := m.(model.Texto)
			htmlForm.AddTextAreaField("Texto", strings.ToLower(fmt.Sprintf("%T", texto)), true)
		case model.Data:
			data := m.(model.Data)
			htmlForm.AddTextField("Data", strings.ToLower(fmt.Sprintf("%T", data)), true)
		default:
			fmt.Print("Model desconhecida!")
		}
	}
    htmlForm.AddSubmitButton()
    htmlForm.RunHtml()
	sql.CreateTableAndInsert(models, htmlForm.Values)
}
