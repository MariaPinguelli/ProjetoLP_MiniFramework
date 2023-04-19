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
}

func (h Html) RunHtml(c chan string) {
	var res string

	http.HandleFunc("/", func(resWriter http.ResponseWriter, resHttp *http.Request) {
        if resHttp.Method == "POST" {
            resHttp.ParseForm()
			res = resHttp.FormValue("Teste")
            fmt.Fprintf(resWriter, "Você digitou: %s", res)
			c <- res
			return
        } else {
            html := h.htmlString
            fmt.Fprint(resWriter, html)
        }
    })

    fmt.Println("Servidor iniciado na porta 8080...")
    http.ListenAndServe(":8080", nil)
}