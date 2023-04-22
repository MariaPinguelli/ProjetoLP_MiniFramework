package field

type Field struct {
    Name string
    Tipo string
	Required bool
    SelectedValue string
}

func CreateField(newName string, newTipo string, isRequired bool) (Field){
    newField := Field{Name: newName, Tipo: newTipo, Required: isRequired}
    return newField
}