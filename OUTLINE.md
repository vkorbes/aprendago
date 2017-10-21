# Outline

## 01 – Visão Geral do Curso

### Por que Go?

### Sucesso

### Recursos do curso

- gobyexample.com
- A Linguagem de Programação Go (Alan Donovan, Brian Kernighan): https://www.amazon.com.br/dp/8575225464/
- Go em Ação (William Kennedy, Brian Ketelsen, Erik St. Martin): https://www.amazon.com.br/dp/8575225065/
- Introdução à Linguagem Go (Caleb Doxsey): https://www.amazon.com.br/dp/8575224891/

### Documentação

### Português vs. inglês

### Por trás do panos



## 02 – Variáveis, Valores & Tipos

### Go Playground

- É online, funciona sem instalar nem configurar nada.
- Assim você pode começar a programar o mais rápido possível.
- Mais pra frente no curso vou explicar direitinho como configurar tudo no seu computador.
- [Go Playground](https://play.golang.org/).
    - Função share. Use para compartilhar código, por exemplo pra fazer uma pergunta em um fórum.
    - Função imports.
    - Função format. 
        - Maneira idiomática: a gente fala da mesma maneira que a comunidade onde estamos.
    - Função run.

### Hello world!

- Estrutura básica: 
    - package main.
    - func main: é aqui que tudo começa, é aqui que tudo acaba.
    - import.
- Packages:
    - Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
    - Notação: pacote.Identificador. Exemplo: fmt.Println()
    - Documentação: fmt.Println.
- Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
- Variáveis não utilizadas? Não pode: _ nelas.
- ...funções variádicas.
- Lição principal: package main, func main, pacote.Identificador.

### Operador curto de declaração

- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática
    - Só pode repetir se houverem variáveis novas 
    - != do assignment operator (operador de atribuição)
    - Só funciona dentro de codeblocks
- Terminologia:
    - keywords (palavras-chave) são termos reservados
    - operadores, operandos
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões 
    - expressão -> qualquer coisa que "produz um resultado"
    - scope (abrangência)
        - package-level scope
- Lição principal:
    - := utilizado pra criar novas variáveis, dentro de code blocks
    - = para atribuir valores a variáveis já existentes

### A palavra-chave var

- Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos `var`
- Funciona em qualquer lugar
- Prestar atenção: chaves, colchetes, parênteses

### Explorando tipos

- Tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
- Tipos em Go são estáticos.
- Ao declarar uma variável para conter valores de um certo tipo, essa variável só poderá conter valores desse tipo.
- O tipo pode ser deduzido pelo compilador:
    - x := 10
    - var y = "a tia do batima"
- Ou pode ser declarado especificamente:
    - var w string = "isso é uma string"
    - var z int = 15
    - na declaração var z int com package scope, atribuição z = 15 no codeblock (somente)
- Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção
    - int, string, bool
- Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário
    - slice, array, struct, map
- O ato de definir, criar, estruturar tipos compostos chama-se composição. Veremos muito disso futuramente.

### Valor zero

- Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
- O que é valor zero?
- Os zeros:
    - ints: 0
    - floats: 0.0
    - booleans: false
    - strings: ""
    - pointers, functions, interfaces, slices, channels, maps: nil
- Use := sempre que possível.
- Use var para package-level scope.

### O pacote fmt

- Setup: strings, ints, bools.
- Strings: interpreted string literals vs. raw string literals.
    - Rune literals.
    - Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte. 
- Format printing: documentação.
    - Grupo #1: Print -> standard out
        - func Print(a ...interface{}) (n int, err error)
        - func Println(a ...interface{}) (n int, err error)
        - func Printf(format string, a ...interface{}) (n int, err error)
            - Format verbs. (%v %T)
    - Grupo #2: Print -> string, pode ser usado como variável
        - func Sprint(a ...interface{}) string
        - func Sprintf(format string, a ...interface{}) string
        - func Sprintln(a ...interface{}) string
    - Grupo #3: Print -> file, writer interface, e.g. arquivo ou resposta de servidor
        - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
        - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
        - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

### Criando seu próprio tipo

- Revisando: tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
- Tem uma história que Bill Kennedy dizia que se um dia fizesse uma tattoo, ela diria "type is life."
- Grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.
- Como fundação para estas ferramentas, vamos aprender a declarar nossos próprios tipos.
- Revisando: tipos são fixos. Uma vez declarada uma variável como de um certo tipo, isso é imutável.
- type hotdog int → var b hotdog (main hotdog)
- Uma variável de tipo hotdog não pode ser atribuida com o valor de uma variável tipo int, mesmo que este seja o tipo subjacente de hotdog. 

### Conversão, não coerção

- Conversão de tipos é o que soa.
- Em Go não se diz casting, se diz conversion.
- a = int(b)
- ref/spec#Conversions
- Fim da sessão. Parabéns! Dicas, motivação e exercícios.

## 03 – Exercícios: Ninja Nível 1

### Contribua seu código

- Nesse curso a gente vai fazer um monte de exercícios.
    - Talvez você queira contribuir suas próprias soluções.
    - Talvez você tenha exemplos melhores que os que estamos mostrando aqui.
- Para compartilhar me manda o link no twitter do seu código no Go Playground, twitter.com/ellenkorbes!
- "@ellenkorbes Olha essa solução pro exercício Ninja nível 5, exercício 2: <link> O que vc acha?"

### Na prática: exercício #1

- Esses são seus primeiros exercícios, e seus primeiros passos.
- Completando os exercícios dessa seção, você será um ninja nível 1.
- É o seu primeiro passo pra se tornar um developer ninja.
- Esses exercícios servem pra reforçar seu aprendizado. Só se aprende a programar programando. Ninguem aprende a andar de bicicleta assistindo vídeos de pessoas andando de bicicleta. É necessário botar a mão na massa.
- Eu vou começar explicando qual é o exercício. Aí vou pedir pra você dar pausa. Esse é o momento de por os miolos pra trabalhar, encontrar sua solução, tec-tec-tec, e rodar pra ver se funciona. Depois é só dar play novamente, ver a minha abordagem para a mesma questão, e comparar nossas soluções.
- Vamos lá:

- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print
- Solução: https://play.golang.org/p/yYXnWXIQNa

### Na prática: exercício #2

- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
    1. Identificador "x" deverá ter tipo int
    2. Identificador "y" deverá ter tipo string
    3. Identificador "z" deverá ter tipo bool
- Na função main:
    1. Demonstre os valores de cada identificador
    2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
- Solução: https://play.golang.org/p/pAFd-F7uGZ

### Na prática: exercício #3

- Utilizando a solução do exercício anterior:
    1. Em package-level scope, atribua os seguintes valores às variáveis:
        1. para "x" atribua 42
        2. para "y" atribua "James Bond"
        3. para "z" atribua true
    2. Na função main:
        1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
        2. Demonstre a variável "s".
- Solução: https://play.golang.org/p/QFctSQB_h3

### Na prática: exercício #4

- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
- Para os aventureiros: https://golang.org/ref/spec#Types
- Agora já somos quase ninjas nível 1!
- Solução: https://play.golang.org/p/snm4WuuYmG

### Na prática: exercício #5

- Utilizando a solução do exercício anterior:
    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    2. Na função main:
        1. Isto já deve estar feito:
            1. Demonstre o valor da variável "x"
            2. Demonstre o tipo da variável "x"
            3. Atribua 42 à variável "x" utilizando o operador "="
            4. Demonstre o valor da variável "x"
        2. Agora faça tambem:
            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
            2. Demonstre o valor de "y"
            3. Demonstre o tipo de "y"
- Solução: https://play.golang.org/p/uq81T_fasP

### Na prática: exercício #6

- Prova!
- Link: https://goo.gl/forms/s9y91iVSPvA4iahj1
- Se você deu pausa e fez todos os exercícios anteriores você mesmo, e só viu a resposta depois... e se você der pausa agora e fizer a prova inteira por conta própria, e só assistir as respostas depois... sabe o que isso quer dizer? Que você é ninja. Ninja nível 1. Tá no caminho certo pra ser um developer ninja mestre.

## 04 – Fundamentos da Programação

### Tipo booleano

- Agora vamos explorar os tipos de maneira mais detalhada. golang.org/ref/spec. A começar pelo bool.
- O tipo bool é um tipo binário, que só pode conter um dos dois valores: true e false. (Verdadeiro ou falso, sim ou não, zero ou um, etc.)
- Sempre que você ver operadores relacionais (==, <=, >=, !=, <, >), o resultado da expressão será um valor booleano.
- Booleans são fundamentais nas tomadas de decisões em lógica condicional, declarações switch, declarações if, fluxo de controle, etc.
- Na prática:
    - Zero value
    - Atribuindo um valor
    - Bool como resultado de operadores relacionais
- Go Playground: https://play.golang.org/p/7joj615nZw

### Como os computadores funcionam

- Isso é importante pois daqui pra frente vamos falar de ints, bytes, e etc.
- Não é necessário um conhecimento a fundo mas é importante ter uma idéia de como as coisas funcionam por trás dos panos.
- https://docs.google.com/presentation/d/1aVytiGOBVDMISFW-ZARJ5iFY1osU2XJIw0hQpNICXm8/
- ASCII: https://en.wikipedia.org/wiki/ASCII
- Filme: Alan Turing, The Immitation Game.

### Tipos numéricos

- int vs. float: Números inteiros vs. números com frações.
- golang.org/ref/spec → numeric types
- Integers:
    - Números inteiros
    - int & uint → “implementation-specific sizes”
    - Todos os tipos numéricos são distintos, exceto:
        - byte = uint8
        - rune = int32 (UTF8)
        (O código fonte da linguagem Go é sempre em UTF-8).
    - Tipos são únicos
        - Go é uma linguagem estática
        - int e int32 não são a mesma coisa
        - Para "misturá-los" é necessário conversão
    - Regra geral: use somente int
- Floating point:
    - Números racionais ou reais
    - Regra geral: use somente float64
- Na prática:
    - Defaults com :=
    - Tipagem com var
    - Dá pra colocar número com vírgula em tipo int?
    - Overflow
    - Go Playground: https://play.golang.org/p/dt2x1ies5b
- “implementation-specific sizes”? Runtime package. Word.
    - GOOS
    - GORUNTIME
    - https://play.golang.org/p/1vp5DImIMM


### Tipo string (cadeias de caracteres)

- Strings são sequencias de bytes.
- Imutáveis.
- Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
- Na prática:
    - %v %T
    - Raw string literals
    - Conversão para slice of bytes: []byte(x)
    - %#U, %#x
    - Go Playground: https://play.golang.org/p/dt2x1ies5b
- https://blog.golang.org/strings

### Sistemas numéricos

- Base-10: decimal, 0–9
- Base-2: binário, 0–1
- Base-16: hexadecimal, 0–f
- https://docs.google.com/document/d/1GqXpubhMMIr4Sy5xwgiPIDh5PGVmVpF2u0c9vDrvykE/
- Demonstração em Go.

### Constantes

- São valores imutáveis.
- Podem ser tipadas ou não:
    - const oi = "Bom dia"
    - const oi string = "Bom dia"
- As não tipadas só terão um tipo atribuido a elas quando forem usadas.
    - Ex. qual o tipo de 42? int? uint? float64?
    - Ou seja, é uma flexibilidade conveniente.
- Na prática: int, float, string.
    - const x = y
    - const ( x = y )

### Iota

- golang.org/ref/spec
- Numa declaração de constantes, o identificador iota representa números sequenciais.
- Na prática.
    - iota, iota + 1, a = iota b c, reinicia em cada const, _
- Go Playground: https://play.golang.org/p/eSrwoQjuYR

### Deslocamento de bits

- Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
- Na prática:
    - %d %b
    - x << y
    - iota * 10 << 10 = kb, mb, gb
    
- https://play.golang.org/p/7MOnbhx4R4
- https://splice.com/blog/iota-elegant-constants-golang/
- https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

- Fim da sessão. Massa!

## 05 – Exercícios: Ninja Nível 2

### Na prática: exercício #1

- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
- Solução: https://play.golang.org/p/X7qm3aWSa6

### Na prática: exercício #2

- Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
    - ==
    - !=
    - <=
    - <
    - >=
    - >
- Demonstre estes valores.
- Solução: https://play.golang.org/p/BMYEch6_s8

### Na prática: exercício #3

- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.
- Solução: https://play.golang.org/p/eWnKI59ual

### Na prática: exercício #4

- Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal
- Solução: https://play.golang.org/p/IiwgT0v3Mp

### Na prática: exercício #5

- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
- Solução: https://play.golang.org/p/RkpqPpRWuo

### Na prática: exercício #6

- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
- Solução: https://play.golang.org/p/zRBEooRvo4

### Na prática: exercício #7

- Prova!
- Link: https://goo.gl/forms/fnPXMmxvKAEUD8xP2
- Motivação. Ninja nível 2!

## 06 – Fluxo de Controle

### Entendendo fluxo de controle

- Computadores lêem programas de uma certa maneira, do mesmo jeito que nós lemos livros, por exemplo, de uma certa maneira.
- Quando nós ocidentais lemos livros, lemos da frente pra trás, da esquerda pra direito, de cima pra baixo.
- Computadores lêem de cima pra baixo.
- Ou seja, sua leitura é sequencial. Isso chama-se fluxo de controle sequencial.
- Alem do fluxo de controle sequencial, há duas declarações que podem afetar como o computador lê o código:
    - Uma delas é o fluxo de controle de repetição (loop). Nesse caso, o computador vai repetir a leitura de um mesmo código de uma maneira específica. O fluxo de controle de repetição tambem é conhecido como fluxo de controle iterativo.
    - E o outro é o fluxo de controle condicional, ou fluxo de controle de seleção. Nesse caso o computador encontra uma condição e, através de uma declaração if ou switch, toma um curso ou outro dependendo dessa condição.
- Ou seja, há três tipos de fluxo de controle: sequencial, de repetição e condicional.

- Nesse capítulo:
    - Sequencial
    - Iterativo (loop)
        - for: inicialização, condição, pós
        - for: hierarquicamente
        - for: condição ("while")
        - for: ...ever?
        - for: break
        - for: continue
    - Condicional
        - declarações switch/case/default
            - não há fall-through por padrão
            - criando fall-through
            - default
            - múltiplos casos
            - casos podem ser expressões
                - se resultarem em true, rodam
            - tipo
        - if
            - bool
            - o operador "!"
            - declaração de inicialização
            - if, else
            - if, else if, else
            - if, else if, else if, ..., else

### Loops: inicialização, condição, pós

- For
    - Inicialização, condição, pós
    - Ponto e vírgula?
    - gobyexample.com
    - Não existe while!

### Loops: nested loop (repetição hierárquica)

- For
    - Repetição hierárquica
    - Exemplos: relógio, calendário

### Loops: a declaração for

- For: inicialização, condição, pós
- For: condição ("while")
- For: ...ever? (http servers)
- For: break
- golang.org/ref/spec#For_statements, Effective Go
- (Range vem mais pra frente.)

### Loops: break & continue

- Operação módulo: %
- For: break
- For: continue
- Go Playground: https://play.golang.org/p/gpKMP1wAEM

### Loops: utilizando ascii

- Desafio surpresa!
- Format printing:
    - Decimal       %d
    - Hexadecimal   %#x
    - Unicode       %#U
    - Tab           \t
    - Linha nova    \n
- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
- Solução: https://play.golang.org/p/REm2WHyzzz

### Condicionais: a declaração if

- If: bool
- If: o operador não → "!"
- If: declaração de inicialização
- Go Playground: https://play.golang.org/p/6nq2Tjb07i

### Condicionais: if, else if, else

- If, else.
- If, else if, else.
- If, else if, else if, ..., else.
- Go Playground: https://play.golang.org/p/18VrRX2pec

### Loops, condicionais, e operação módulo

- Revisão:
- For + condicionais + módulo
- Números múltiplos de x

### Condicionais: a declaração switch

- Switch:
    - pode avaliar uma expressão 
        - switch statement == case (value)
        - default switch statement == true (bool)
    - não há fall-through por padrão
    - criando fall-through
    - default
    - cases compostos

### Condicionais: a declaração switch pt. 2 & documentação

- Pode avaliar tipos
- Pode haver uma expressão de inicialização

### Operadores lógicos condicionais

- &&
- ||
- !
- Go Playground: https://play.golang.org/p/MFwrt93xlc
- Qual o resultado de fmt.Println...
    - true && true
    - true && false
    - true || true
    - true || false
    - !true

## 07 – Exercícios: Ninja Nível 3

### Na prática: exercício #1

- Põe na tela: todos os números de 1 a 10000.
- Solução: https://play.golang.org/p/MkdZiDW8SQ

### Na prática: exercício #2

- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
    65
        U+0041 'A'
        U+0041 'A'
        U+0041 'A'
    66
        U+0042 'B'
        U+0042 'B'
        U+0042 'B' 
    ...e por aí vai.
- Solução: https://play.golang.org/p/QlP6nVchq-

### Na prática: exercício #3

- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/PzUqGLODmW

### Na prática: exercício #4

- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/zCU9QWkzJ9

### Na prática: exercício #5

- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
- Solução: https://play.golang.org/p/zcEsXqnBr8

### Na prática: exercício #6

- Crie um programa que demonstre o funcionamento da declaração if.
- Solução: https://play.golang.org/p/OGPgTJZbiy

### Na prática: exercício #7

- Utilizando a solução anterior, adicione as opções else if e else.
- Solução: https://play.golang.org/p/_cO7E-yL0o

### Na prática: exercício #8

- Crie um programa que utilize a declaração switch, sem switch statement especificado.
- Solução: https://play.golang.org/p/TyIGp4Hi8B

### Na prática: exercício #9

- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
- Solução: https://play.golang.org/p/4-iTPZwfEz

### Na prática: exercício #10

- Anote (à mão) o resultado das expressões:
    - fmt.Println(true && true) 
    - fmt.Println(true && false) 
    - fmt.Println(true || true) 
    - fmt.Println(true || false) 
    - fmt.Println(!true)
- Ninja nível 3! Parabéns!

## 08 – Agrupamentos de Dados

### Array

- Estruturas de dados, ou agrupamentos de dados, nos permitem agrupar valores diferentes. Estes valores podem ser ou não do mesmo tipo.
- As estruturas que veremos são: arrays, slices, structs e maps.
- Vamos começar com arrays. Arrays em Go são uma fundação, e não algo que utilizamos todo dia.
- Seu tamanho deve estar presente na declaração: var x [n]int
- Atribui-se valores a suas posições com: x[i] = y (0-based)
- Para ver o tamanho usa-se: len(x)
- ref/spec: "The length is part of the array's type" → [5]int != [6]int
- Effective Go: Arrays são úteis para [umas coisas que a gente não vai fazer nunca] e servem de fundação para slices. Use slices ao invés de arrays.
- Go Playground: 

### Slice: literal composta

- O que são tipos de dados compostos? 
    - Wikipedia: Composite_data_type
    - Effective Go: Composite literals
    - ref/spec: Composite literals
- Uma slice agrupa valores de um único tipo.
- Criando uma slice: literal composta → x := []type{values}
- Go Playground: 

### Slice: for range

- Slices:
    - Tamanho: len(x)
    - Índice específico: x[i] (0-based)
- Para ver todos os itens de uma slice utilizamos o loop for com range.
- Range significa alcance, faixa, extensão.
- For range: for i, v := range x {}
- Go Playground: 

### Slice: fatiando uma fatia

- x[:], x[a:], x[:b], x[a:b]
- "a" é incluso;
- "b" não é.
- Exemplo: cabeça magnética de um disco rígido (relógio, fita).
    - Off-by-one error.
- É fatiando que se deleta um item de uma slice.
- Exercício: tente acessar todos os itens de uma slice *sem* utilizar range.
- Solução:

### Slice: anexando a uma slice

- Effective Go: append (package builtin)
- x = append(slice, ...values)
- x = append(slice, slice...)
- Todd: unfurl → desdobrar, desenrolar
- Nome oficial: enumeration
- Go Playground: 

### Slice: deletando de uma slice

- x := append(x[:i], x[:i]...)
- Go Playground: 

### Slice: make

- Slices são feitas de arrays.
- Elas são dinâmicas, podem mudar de tamanho.
- Sempre que isso acontece, um novo array é criado e os dados são copiados.
- É conveniente, mas tem um custo computacional.
- Para otimizar as coisas, podemos utilizar make.
- make([]T, len, cap)
- "The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume."
- len(x), cap(x)
- x[n] onde n > len é out of range. Use append.
- Append > cap modifica o array subjacente.
- pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
- Effective Go.
- Go Playground: 

### Slice: slice multi-dimensional

- Slices multi-dimensionais são slices que contem slices.
- São como planilhas.
- [][]type
- Go Playground: 

### Slice: o array subjacente

- Toda slice tem um array subjacente.
- Um slice é: um ponteiro/endereço para um array, mais len e cap (que é o len to array).
- Exemplo:
    - x := []int{...números}
    - y := append(x[:i], x[:i]...)
    - pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
    - Ou seja, y utiliza o mesmo array subjacente que x.
    - O que nos dá um resultado inesperado.
- Ou seja, bom saber de antemão pra não ter que aprender na marra.
- Go Playground: 

### Maps: introdução

- Utiliza o formato key:value.
- E.g. nome e telefone
- Performance excelente para lookups.
- map[key]value{ key: value }
- Acesso: m[key]
- Key sem value retorna zero. Isso pode trazer problemas.
- Para verificar: comma ok idiom.
    - v, ok := m[key]
    - ok é um boolean, true/false
- Na prática: if v, ok := m[key]; ok { }
- Maps *não tem ordem.*
- Go Playground: 

### Maps: adicionando elementos & range

- Para adicionar um item: m[v] = value
- Range: for k, v := range map { }
- Reiterando: maps *não tem ordem* e um range usará uma ordem aleatória.
- Go Playground: 

### Maps: deletando

- delete(map, key)
- Deletar uma key não-existente não retorna erros!
- Go Playground: 

## 09 – Exercícios: Ninja Nível 4

### Na prática: exercício #1

- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
- Solução: 

### Na prática: exercício #2

- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
- Solução: 

### Na prática: exercício #3

- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
- Solução: 

### Na prática: exercício #4

- Começando com a seguinte slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
    - y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x.
- Solução: 

### Na prática: exercício #5

- Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]
- Solução: 

### Na prática: exercício #6

- Crie uma slice usando make que possa conter todos os estados do Brasil.
    - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice *sem utilizar range.*
- Solução: 


### Na prática: exercício #7

- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
- Solução: 

### Na prática: exercício #8

- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
- Solução: 

### Na prática: exercício #9

- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
- Solução: 

### Na prática: exercício #10

- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
- Solução: 

## 10 – Structs

### Struct

- Struct é um tipo de dados composto que nos permite armazenar valores de tipos *diferentes.*
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.
- Go Playground: 

### Structs embutidos

- Structs dentro de structs dentro de structs.
- Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) *e tambem* um competidor (nome, equipe, pontos).
- Go Playground: 

### Lendo a documentação

- É importante se familiarizar com a documentação da linguagem Go.
- Neste vídeo vamos ver um pouco sobre o que a documentação diz sobre structs.
- Veremos:
    - ref/spec
        - Já vimos mais da metade dos tipos em Go!
        - Struct types.
            - x, y int
            - anonymous fields
            - promoted fields
- Go Playground: 

### Structs anônimos

- São structs sem identificadores.
- x := struct { name type }{ name: value }
- Go Playground: 

### Colocando ordem na casa

- Go Playground: 

## 11 – Exercícios: Ninja Nível 5

### Na prática: exercício #1

- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
- Solução: 

### Na prática: exercício #2

- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
- Solução: 

Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

### Na prática: exercício #3

- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "quatroRodas"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
- Solução: 

### Na prática: exercício #4

- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
- Solução: 

## 12 – Funções

### Sintaxe

- Qual a utilidade de funções?
    - Abstrair funcionalidade
    - Reutilização de código
- func (receiver) identifier(parameters) (returns) { code }
- A diferença entre parâmetros e argumentos:
    - Funções são definidas com parâmetros
    - Funções são chamadas com argumentos
- Tudo em Go é *pass by value.*
    - Pass by reference, pass by copy, ... não.
- Exemplos:
    - Função básica. 
        - Go Playground: 
    - Função que aceita um argumento. 
        - Go Playground: 
    - Função com retorno. 
        - Go Playground: 
    - Função com múltiplos retornos. 
        - Go Playground: 

### Parâmetro variádico

- ...variádico.
- Exemplo: função que retorna o total de todos os ints recebidos.
- Go Playground: 

### Desenrolando (enumerando) uma slice

- Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
- Exemplos:
    - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
        - Go Playground: 
    - Pode-se passar *zero* ou mais valores
        - Go Playground: 
    - O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
        - Go Playground: 
    - Pode-se passar uma "slice..."
        - Go Playground: 

### Defer

- Funções são ótimas pois tornam nosso código modular. Podemos alterar partes do nosso programa sem afetar o resto!
- Uma declaração defer chama uma função cuja execução ocorrerá no momento em que a função da qual ela faz parte finalizar.
- Essa finalização pode ocorrer devido a um return, ao fim do code block da função, ou no caso de pânico em uma goroutine correspondente.
- "Deixa pra última hora!"
- ref/spec
- Sempre usamos para fechar um arquivo após abri-lo.
- Go Playground: 
 
### Métodos

- Um método é uma função anexada a um tipo.
- Quando se anexa uma função a um tipo, ela se torna um método desse tipo.
- Pode-se anexar uma função a um tipo utilizando seu receiver.
- Utilização: valor.método()
- Exemplo: o tipo "pessoa" pode ter um método oibomdia()
- Go Playground: 

### Interfaces & polimorfismo

- Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo *e tambem* o tipo da interface.
- Em Go, valores podem ter mais que um tipo.
- Exemplos:
    - Os tipos *profissão1* e *profissão2* contem o tipo *pessoa*
    - Cada um tem seu método *oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()*
    - Implementam a interface *gente*
    - Ambos podem acessar o método *serhumano()* que chama o método *oibomdia()* de cada *gente*
    - Tambem podemos no método *oibomdia()* tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }* 
    - Go Playground: https://play.golang.org/p/VLbo_1uE-U
- Onde se utiliza?
    - Área de formas geométricas (gobyexample.com)
    - Sort
    - DB
    - Writer interface: arquivos locais, http request/response
- Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.

### Funções anônimas

- Anonymous self-executing functions → Funções anônimas auto-executáveis.
- func(p params) { ... }()
- Vamos ver bastante quando falarmos de goroutines.
- Go Playground:

### Func como expressão

- f := func(p params){ ... }
- f()
- Go Playground:

### Retornando uma função

- Pode-se usar uma função como retorno de uma função
- Declaração: func f() return
- Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
    - ????: fmt.Println(f()())
- Go Playground:

### Callback

- Primeiro veja se você entende isso: https://play.golang.org/p/QkAtwMZU-z
- Callback é passar uma função como argumento.
- Exemplo:
    - Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.
    - Go Playground:
- Desafio: Crie uma função no programa acima que utilize somente os números *ímpares.*
- Solução: 

### Closure

- Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto. Já vimos:
    - Package-level scope
    - Function-level scope
    - Code-block-in-code-block scope
- Exemplo de closure:
    - func i() func() int { x := 0; return func() int { x++; return x } }
    - Quando fizermos a := i() teremos um scope, um valor para x.
    - Quando fizermos b := i() teremos outro scope, e x terá um valor independente do x acima.
- Go Playground:

### Recursividade

- WP: "The most common application of recursion is in mathematics and computer science, where a function being defined is applied within its own definition."
- Exemplos de recursividade: Fractais, matrioscas, efeito Droste (o efeito produzido por uma imagem que aparece dentro dela própria), GNU (“GNU is Not Unix”), etc.
- No estudo de funções: é uma função que chama a ela própria.
- Exemplo: fatoriais.
    - 4! = 4 * 3 * 2 * 1 (e no zero, deu.)
    - Com recursividade. Go Playground:
    - Com loops. Go Playground:

## 13 – Exercícios: Ninja Nível 6

### Na prática: exercício #1

- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
- Solução: 

- Revisão:
    - Funções!
        - Servem para abstrair código
        - E para reutilizar código
    - A ordem das coisas é:
        - func (receiver) identifier (parameters) (returns) { code }
    - Parâmetros vs. argumentos
    - Funções variádicas
        - Múltiplos parâmetros
        - Múltiplos argumentos
    - Métodos
    - Interfaces & polimorfismo
    - Defer
        - "Deixa pra depois!"
    - Returns
        - Múltiplos returns
        - Returns com nome (blé!)
    - Funcs como expressões
        - Atribuindo uma função a uma variável
    - Callbacks
        - Passando uma função como argumento para outra função
    - Closure
        - Capturando um scope
        - Variáveis declaradas em scopes externos são visíveis em scopes internos
    - Recursividade
        - Uma função que chama a ela mesma
        - Fatoriais

### Na prática: exercício #2

- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
- Solução: 

### Na prática: exercício #3

- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
- Solução: 

### Na prática: exercício #4

- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
- Solução: 

### Na prática: exercício #5

- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
- Solução: 

### Na prática: exercício #6

- Crie e utilize uma função anônima.
- Solução: 

### Na prática: exercício #7

- Atribua uma função a uma variável.
- Chame essa função.
- Solução: 

### Na prática: exercício #8

- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
- Solução: 

### Na prática: exercício #9

- Callback: passe uma função como argumento a outra função.
- Solução: 

### Na prática: exercício #10

- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
- Solução: 

### Na prática: exercício #11

- Uma das melhores maneiras de aprender é ensinando.
- Para este exercício escolha o seu código favorito dentre os que vimos estudando funções. Pode ser das aulas ou da seção de exercícios. Então:
    - Faça download e instale isso aqui: https://obsproject.com/ 
    - Grave um vídeo onde *você* ensina o tópico em questão
    - Faça upload do vídeo no YouTube
    - Compartilhe o vídeo no Twitter e me marque no tweet (@ellenkorbes)

## 14 – Ponteiros

### O que são ponteiros?

- Todos os valores ficam armazenados na memória.
- Toda localização na memória possui um endereço.
- Um pointeiro se refere a esse endereço.
- Notações:
    - &variável mostra o endereço de uma variável
        - %T: variável vs. &variǻvel
    - *variável faz de-reference, mostra o valor que consta nesse endereço
    - ????: *&var funciona!
    - *type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador
- Exemplo: a := 0; b := &a; *b++
- Go Playground: 

### Quando usar ponteiros

- Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
    - Não queremos passar grandes volumes de dados pra lá e pra cá
    - Queremos mudar um valor em sua localização original (tudo em Go é pass by value!)
- Exemplos:
    - x := 0; funçãoquemudaovalordoargumentopra1(x); Print(x)
    - x := 0; funçãoquemudaovalordo*argumentopra1(&x); Print(x)
- Go Playground: 

### Conjuntos de métodos

- Conjuntos de métodos são conjuntos de métodos :)
- Um receiver não-ponteiro recebe tanto pointeiros como não-ponteiros.
- Um receiver ponteiro recebe somente ponteiros.
- Go Playground: 

## 15 – Exercícios: Ninja Nível 7

### Na prática: exercício #1

- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.
- Solução: 

### Na prática: exercício #2

- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar o valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
    - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), 
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado.
- Solução: 

## 16 – Aplicação

### Documentação JSON

- Go Playground: 

### JSON marshal (ordenação)

- Go Playground: 

### JSON unmarshal (desordenação)

- Go Playground: 

### A interface Writer

- Go Playground: 

### O pacote sort

- Go Playground: 

### Customizando o sort

- Go Playground: 

### bcrypt

- Go Playground: 

## 17 – Exercícios: Ninja Nível 8

### Na prática: exercício #1

- Solução: 

### Na prática: exercício #2

- Solução: 

### Na prática: exercício #3

- Solução: 

### Na prática: exercício #4

- Solução: 

### Na prática: exercício #5

- Solução: 

## 18 – Concorrência

### Concorrência vs. paralelismo

- Go Playground: 

### WaitGroup

- Go Playground: 

### Conjuntos de métodos, parte 2

- Go Playground: 

### Documentação

- Go Playground: 

### Condição de corrida

- Go Playground: 

### Mutex

- Go Playground: 

### Atomic

- Go Playground: 

## 19 – Seu Ambiente de Desenvolvimento

### O terminal



### Bash no Windows



### Instalando Go



### Go workspace



### Environment variables



### IDE's



### Comandos Go



### Repositórios no GitHub



### Explorando o GitHub



### Compilação cruzada



### Pacotes



## 20 – Exercícios: Ninja Nível 9

### Na prática: exercício #1

- Solução: 

### Na prática: exercício #2

- Solução: 

### Na prática: exercício #3

- Solução: 

### Na prática: exercício #4

- Solução: 

### Na prática: exercício #5

- Solução: 

### Na prática: exercício #6

- Solução: 

### Na prática: exercício #7

- Solução: 

## 21 – Canais

### Entendendo canais



### Canais direcionais



### Utilizando canais



### Range



### Select



### A expressão comma ok



### Convergência



### Divergência



### Context



## 22 – Exercícios: Ninja Nível 10

### Na prática: exercício #1

- Solução: 

### Na prática: exercício #2

- Solução: 

### Na prática: exercício #3

- Solução: 

### Na prática: exercício #4

- Solução: 

### Na prática: exercício #5

- Solução: 

### Na prática: exercício #6

- Solução: 

### Na prática: exercício #7

- Solução: 

## 23 – Tratamento de Erros

### Entendendo erros



### Verificando erros



### Print & log



### Recover



### Erros com informações adicionais



## 24 – Exercícios: Ninja Nível 11

### Na prática: exercício #1

- Solução: 

### Na prática: exercício #2

- Solução: 

### Na prática: exercício #3

- Solução: 

### Na prática: exercício #4

- Solução: 

### Na prática: exercício #5

- Solução: 

## 25 – Escrevendo Documentação

### Introdução



### go doc



### godoc



### godoc.org



### Escrevendo documentação



## 26 – Exercícios: Ninja Nível 12

### Na prática: exercício #1

- Solução: 

### Na prática: exercício #2

- Solução: 

### Na prática: exercício #3

- Solução: 

## 27 – Testes & Benchmarking

### Introdução



### Testes em tabela



### Testes como exemplos



### Golint



### Benchmark



### Cobertura



### Exemplos de benchmarks

- Ponteiros

### Revisão



## 28 – Exercícios: Ninja Nível 13

### Na prática: exercício #1

- Solução: 

### Na prática: exercício #2

- Solução: 

### Na prática: exercício #3

- Solução: 

## 29 – Considerações Finais

 