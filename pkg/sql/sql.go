package sql

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/html"
)

func CreateTableAndInsert(tableName string, idField string, h h.Html, destPath string) {
    sqlQuery := "CREATE TABLE "+tableName
    sqlQueryFields := "INSERT INTO "+tableName+" ("
    sqlQueryValues := "VALUES ("

    _, ok := h.Values[idField]

    if ok {
        sqlQuery = fmt.Sprintf("%s %s %s PRIMARY KEY,", sqlQuery, idField, getType(idField))
    }else{
        sqlQuery = sqlQuery + " id INT PRIMARY KEY,"
    }

    for i := 0; i < len(h.FieldList); i++ {
        if idField != h.FieldList[i]{
            sqlQuery = fmt.Sprintf("%s %s %s", sqlQuery, h.Values[h.FieldList[i]], getType(h.Values[h.FieldList[i]]))
            sqlQueryFields = fmt.Sprintf("%s%s",sqlQueryFields,h.FieldList[i])
            if strings.Contains(getType(h.Values[h.FieldList[i]]), "VARCHAR"){
                sqlQueryValues = fmt.Sprintf("%s'%s'", sqlQueryValues, h.Values[h.FieldList[i]])
            }else{
                sqlQueryValues = fmt.Sprintf("%s%s", sqlQueryValues, h.Values[h.FieldList[i]])
            }
        }

        if i == len(h.FieldList)-1{
            sqlQuery = sqlQuery+");\n"
            sqlQueryFields = sqlQueryFields+") "
            sqlQueryValues = sqlQueryValues+");\n"
        }else{
            sqlQuery = sqlQuery+", "
            sqlQueryFields = sqlQueryFields+", "
            sqlQueryValues = sqlQueryValues+", "
        }
    }

    sqlQueryInsert := sqlQueryFields + sqlQueryValues

    file, err := os.Create(destPath+tableName+".sql")
    
	if err != nil {
        fmt.Println("Erro ao criar o arquivo SQL:", err)
        return
    }

    _, err = file.WriteString(sqlQuery)
    if err != nil {
        fmt.Println("Erro ao inserir querys no arquivo: ", err)
        return
    }

    _, err = file.WriteString(sqlQueryInsert)
    if err != nil {
        fmt.Println("Erro ao inserir querys no arquivo: ", err)
        return
    }

	defer file.Close()

    fmt.Println("Querys SQL escrita com sucesso <3")
}

func getType(value interface{}) string {
    tipo := reflect.TypeOf(value)
    switch tipo.Kind() {
    case reflect.Bool:
        return "BOOLEAN"
    case reflect.String:
        return "VARCHAR(" + strconv.Itoa(int(math.Ceil(float64(len(value.(string))/10) * 10))) + ")"
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return "INTEGER"
    case reflect.Float32, reflect.Float64:
        return "FLOAT"
    }
    return "UNKNOWN"
}