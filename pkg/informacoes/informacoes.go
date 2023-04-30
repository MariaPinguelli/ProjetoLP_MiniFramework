package informacoes

import (
	model "github.com/MariaPinguelli/ProjetoLP_MiniFramework/pkg/model")

type Informacao struct {
	Nome              model.Frase
	Biografia         model.Texto
	UltimaAtualizacao model.Data
}

func NovaInformacao(nome *model.Frase, bio *model.Texto, ultima_atualizacao *model.Data) *Informacao {
	return &Informacao{Nome: *nome, Biografia: *bio, UltimaAtualizacao: *ultima_atualizacao}
}
