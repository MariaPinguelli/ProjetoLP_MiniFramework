package src

import "fmt"

type Field struct {
    Name string
    Type string
	Required bool
}

func (f *Field) printName() {
    fmt.Println(f.Name)
}