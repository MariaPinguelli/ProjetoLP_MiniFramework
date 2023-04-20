package html

import (
	fieldBuilder "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/field"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"net/http"
	"strings"
)

type Html struct{
	htmlString string
	fieldList []string
	values map[string]string
}

func StartHtml() (Html){
	template := Html{htmlString: `<html><body><form method="post"></form></body></html>`}

	return template
}

func (h *Html) AddField(){
	var required string = ""
	campo := fieldBuilder.Field{ Name: "Teste", Tipo: "text", Required: true}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(h.htmlString))

    form := doc.Find("form")

	if (campo.Required) {
		required = "required"
	}else{
		required = ""
	}

    form.AppendHtml(fmt.Sprintf(`
		<div>
			<label for="%s">%s:</label>
			<input type='%s' name='%s' %s>
			<input type="submit" value="Submit">
		</div>
	`, campo.Name, campo.Name, campo.Tipo, campo.Name, required))

    htmlString, _ := doc.Html()

	h.htmlString = htmlString
	h.fieldList = append(h.fieldList, campo.Name)
	h.values = make(map[string]string)
}

func (h* Html) RunHtml() {
	http.HandleFunc("/", func(resWriter http.ResponseWriter, resHttp *http.Request) {
        if resHttp.Method == "POST" {
            resHttp.ParseForm()
			for i := 0; i < len(h.fieldList); i++ {
				h.values[h.fieldList[i]] = resHttp.FormValue(h.fieldList[i])
				fmt.Print(h.fieldList[i],":",h.values[h.fieldList[i]])
			}
            fmt.Fprint(resWriter, "Formulário enviado!")
        } else {
            html := h.htmlString
            fmt.Fprint(resWriter, html)
        }
    })

    fmt.Println("Servidor iniciado na porta 8080...")
    http.ListenAndServe(":8080", nil)
}