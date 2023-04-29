package generator

import (
	"fmt"
	html "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/html"
	//model "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/model"
	"strings"
)

func GenerateHTMLFields(args []string) string {
    html.StartHtml()
	var fields string
    for _, arg := range args {
        fields += fmt.Sprintf(`<label for="%s">%s:</label>
            <input type="text" id="%s" name="%s"><br><br>`, arg, arg, arg, arg)
    }
    return fields
}

func ProcessInput(args []string) ([]string, string) {
    var flags []string
    var arguments []string

    for _, arg := range args {
        if strings.HasPrefix(arg, "-") {
            flags = append(flags, arg)
        } else {
            arguments = append(arguments, arg)
        }
    }

    fields := GenerateHTMLFields(arguments)

    return arguments, fields
}



