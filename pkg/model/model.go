package model

// import (
// 	"time"
// )

type Frase struct {
	Tamanho int
	Verboso string
	Required    bool
}

type Texto struct {
	Verboso string
	Required    bool
}

type Data struct {
	Verboso string
	Required    bool
}

func NovaFrase(tamanho int, verboso string, required bool) *Frase {
	return &Frase{Tamanho: tamanho, Verboso: verboso, Required: required}
}

func NovoTexto(verboso string, required bool) *Texto {
	return &Texto{Verboso: verboso, Required: required}
}

func NovaData(required bool, verboso string) *Data {
	return &Data{Verboso: verboso, Required: required}
}
