package html

import (
	"fmt"
	"net/http"
	"strings"

	fieldBuilder "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/field"
	"github.com/PuerkitoBio/goquery"
)

type Html struct {
	htmlString string
	FieldList  []string
	Values     map[string]string
}

func StartHtml() Html {
	template := Html{htmlString: `
	<!DOCTYPE html>
	<html>
		<head>
			<style>
				form {
					margin: 0 auto;
					width: 400px;
				}
				
				div {
					margin: 30px;
				}
				
				div {
					display: flex;
					justify-content: space-between;
				}
			</style>
		</head>
		<body>
			<form method="post"></form>
		</body>
	</html>`}

	template.Values = make(map[string]string)
	return template
}

func (h *Html) AddTextField(label string, name string, isRequired bool) {
	var required string = ""
	campo := fieldBuilder.CreateField(name, "text", isRequired)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(h.htmlString))

	form := doc.Find("form")

	if campo.Required {
		required = "required"
	} else {
		required = ""
	}

	form.AppendHtml(fmt.Sprintf(`
		<div>
			<label for="%s">%s:</label>
			<input type='%s' name='%s' %s>
		</div>
	`, campo.Name, strings.Title(label), campo.Tipo, campo.Name, required))

	htmlString, _ := doc.Html()

	h.htmlString = htmlString
	h.FieldList = append(h.FieldList, campo.Name)
}

func (h *Html) AddTextAreaField(label string, name string, isRequired bool) {
	var required string = ""
	campo := fieldBuilder.CreateField(name, "text", isRequired)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(h.htmlString))

	form := doc.Find("form")

	if campo.Required {
		required = "required"
	} else {
		required = ""
	}

	form.AppendHtml(fmt.Sprintf(`
		<div>
			<label for="%s">%s:</label>
			<textarea name="%s" %s></textarea>
		</div>
	`, campo.Name, strings.Title(label), campo.Name, required))

	htmlString, _ := doc.Html()

	h.htmlString = htmlString
	h.FieldList = append(h.FieldList, campo.Name)
}

func (h *Html) AddSelectField(label string, name string, isRequired bool, items []string) {
	var required string = ""
	var optionList string
	campo := fieldBuilder.CreateField(name, "text", isRequired)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(h.htmlString))

	form := doc.Find("form")

	if campo.Required {
		required = "required"
	} else {
		required = ""
	}

	for i := 0; i < len(items); i++ {
		optionList = fmt.Sprintf(`%s <option value="%s">%s</option>`, optionList, items[i], strings.Title(items[i]))
	}

	form.AppendHtml(fmt.Sprintf(`
		<div>
			<label for="%s">%s:</label>
			<select name="%s" %s>
				<option value="">-- Selecione uma opção --</option>
				%s
  			</select>
		</div>
	`, campo.Name, strings.Title(label), campo.Name, required, optionList))

	htmlString, _ := doc.Html()

	h.htmlString = htmlString
	h.FieldList = append(h.FieldList, campo.Name)
}

func (h *Html) AddSubmitButton() {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(h.htmlString))

	form := doc.Find("form")

	form.AppendHtml(`
		<div>
			<button type="submit" formmethod="post">Enviar formulário!</button>
		</div>
	`)

	htmlString, _ := doc.Html()

	h.htmlString = htmlString
}

func (h *Html) RunHtml() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(resWriter http.ResponseWriter, resHttp *http.Request) {
		if resHttp.Method == "POST" {
			resHttp.ParseForm()
			for i := 0; i < len(h.FieldList); i++ {
				h.Values[h.FieldList[i]] = resHttp.FormValue(h.FieldList[i])
				fmt.Println(h.FieldList[i], ":", h.Values[h.FieldList[i]])
			}
			fmt.Print("\n\n")
			fmt.Fprint(resWriter, "Formulário enviado!")
			server.Close()
			return
		} else {
			html := h.htmlString
			fmt.Fprint(resWriter, html)
		}
	})

	fmt.Println("Servidor iniciado na porta 8080...")
	server.ListenAndServe()
}