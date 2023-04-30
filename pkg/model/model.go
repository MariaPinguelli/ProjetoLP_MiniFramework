package model

import (
	"time"
)

type Frase struct {
	Tamanho int
	Verboso string
	Nulo    bool
}

type Texto struct {
	Verboso string
	Nulo    bool
}

type Data struct {
	Verboso string
	Nulo    bool
}

func NovaFrase(tamanho int, verboso string) *Frase {
	return &Frase{Tamanho: tamanho, Verboso: verboso}
}

func NovoTexto(verboso string) *Texto {
	return &Texto{Verboso: verboso}
}

func NovaData() *Data {
	return &Data{Verboso: time.Now().Format("2006-01-02")}
}
