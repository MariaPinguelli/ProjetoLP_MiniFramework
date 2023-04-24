# ProjetoLP_MiniFramework

<h2> Links importantes </h2>
Link no moodle: https://moodle.utfpr.edu.br/mod/assign/view.php?id=1295950

Link de instalação do Go: https://go.dev/doc/install

Essa estrutura de arquitetura de projeto é diferente, porque o Go tem sua própria convenção para modularidade de projetos, e não é possível rodar o projeto, sem usar essa convenção :)
Link sobre project layout: https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md

<h2> Documentação da linguagem </h2>
<h3> Sobre Go </h3>
Go é uma linguagem de programação imperativa (ou seja, os programas são escritos como uma sequência lógica de instruções que são executadas na ordem) que possui algumas características de linguagens funcionais, como a passagem de funções como argumentos e closures ("captura" uma variável definida em um escopo, mesmo que este não esteja mais disponível). Além disso, é uma linguagem com tipagem estática, isto é, os tipos de variáveis são definidos em tempo de execução e não podem ser alterados durante a execução do programa. 
<h3> TAD e OO em Go </h3>
Go não tem classes. Em vez disso, usa a composição de structs e interfaces para a implementação de polimorfismo e encapsulamento. AS structs são usadas para definir tipos de dados com campos que podem ser acessados diretamente; também podem incorporar outros tipos de dados (composição de tipos em vez de herança de classes). Uso de pacotes para a modularidade.<br>
O encapsulamento em Go é dado pela delaração de visibilidade de tipos e variáveis. Variáveis e tipos declarados com letras maiúsculas são públicos e podem ser acessados de qualquer pacote (ou módulo); já em minúsculos são privados e, portanto, só podem ser acessados dentro do pacote no qual foram definidos. <br>
<p> Exemplo de código em Go: </p>

```
package main

import "fmt"

type Pessoa struct { // Struct
    Nome  string // Atributo público
    idade int // Atributo privado
}

func main() {
    p := Pessoa{"João", 30}
    fmt.Println(p.Nome) // Imprime "João"
    fmt.Println(p.idade) // Erro de compilação: `idade` não é pública
}
```
O polimorfismo em Go é alcançado por meio de interfaces. Observação: em Go, um tipo é considerado compatível com uma interface se implementar todos os métodos definidos na interface. Exemplo:
```
type Forma interface {
    area() float64
    perimetro() float64
}

func (r Retangulo) area() float64 {
    return r.largura * r.altura
}

func (r Retangulo) perimetro() float64 {
    return 2*r.largura + 2*r.altura
}
```
Um tipo é considerado compatível com uma interface se implementar todos os métodos definidos na interface. Temos uma interface Forma e Retangulo implementa todos os métodos de Forma. Logo, é compatível e pode-se atribuir uma variável do tipo 'Retangulo' a uma do tipo 'Formula'.

<h2> Documentação do projeto </h2
Para desenvolver um mini-framework web, nos inspiramos no pacote "html/template", que fornece um mecanismo de templates para gerar código HTML dinamicamente, e no pacote "database/sql", que fornece uma interface comum para acessar diferentes bancos de dados, sendo ambos da biblioteca padrão.
