package sql

import (
	"fmt"
	// "math"
	// "os"
	// "reflect"
	// "strconv"
	// "strings"
	// h "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/html"
	model "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/model"
)

func CreateTableAndInsert(tableName string, data []interface{}, values map[string]string) {
    retorno := "CREATE TABLE ("+tableName+" id INTEGER PRIMARY KEY,"
    camposInsert := "INSERT INTO "+tableName+" ("
    valuesInsert := "VALUES ("
    insert := ""
    
    resTable := ""
    resInsert := ""
    resValues := ""
    
    for _, c := range data {
		// Checar o tipo do modelo e gerar um campo HTML apropriado
		switch c.(type) {
		case model.Frase:
			frase := c.(model.Frase)
            resTable, resInsert, resValues = AddFrase(frase, values[frase.Verboso])
		case model.Texto:
			texto := c.(model.Texto)
			resTable, resInsert, resValues = AddTexto(texto, values[texto.Verboso])
        case model.Data:
			data := c.(model.Data)
            resTable, resInsert, resValues = AddData(data, values[data.Verboso])
        default:
			fmt.Print("Model desconhecida!")
		}
        retorno = retorno+resTable
        camposInsert = camposInsert+resInsert
        valuesInsert = valuesInsert+resValues
	}

    insert = camposInsert+valuesInsert
    fmt.Println(retorno)
    fmt.Println(insert)
}

func AddFrase(frase model.Frase, value string) (string, string, string){
    return fmt.Sprintf(`%s VARCHAR (%d),`, frase.Verboso, frase.Tamanho), 
    fmt.Sprintf(`%s,`, frase.Verboso),
    fmt.Sprintf(`"%s",`, value)
}

func AddTexto(texto model.Texto, value string) (string, string, string){
    return fmt.Sprintf(`%s TEXT,`, texto.Verboso), 
    fmt.Sprintf(`%s,`, texto.Verboso),
    fmt.Sprintf(`"%s",`, value)
}

func AddData(data model.Data, value string) (string, string, string) {
    return fmt.Sprintf(`%s DATE);`, data.Verboso), 
    fmt.Sprintf(`%s)`, data.Verboso),
    fmt.Sprintf(`"%s");`, value)
}