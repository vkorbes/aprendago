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
Solução: https://play.golang.org/p/yYXnWXIQNa

### Na prática: exercício #2

- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
    1. Identificador "x" deverá ter tipo int
    2. Identificador "y" deverá ter tipo string
    3. Identificador "z" deverá ter tipo bool
- Na função main:
    1. Demonstre os valores de cada identificador
    2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
Solução: https://play.golang.org/p/pAFd-F7uGZ

### Na prática: exercício #3

- Utilizando a solução do exercício anterior:
    1. Em package-level scope, atribua os seguintes valores às variáveis:
        1. para "x" atribua 42
        2. para "y" atribua "James Bond"
        3. para "z" atribua true
    2. Na função main:
        1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
        2. Demonstre a variável "s".
    Solução: https://play.golang.org/p/QFctSQB_h3

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
Solução: https://play.golang.org/p/snm4WuuYmG

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
Solução: https://play.golang.org/p/uq81T_fasP

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
Go Playground: https://play.golang.org/p/7joj615nZw

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
Solução: https://play.golang.org/p/IiwgT0v3Mp

### Na prática: exercício #5

- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
Solução: https://play.golang.org/p/RkpqPpRWuo

### Na prática: exercício #6

- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
Solução: https://play.golang.org/p/zRBEooRvo4

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
Solução: https://play.golang.org/p/REm2WHyzzz

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

    - pode avaliar tipos
    - pode haver uma expressão de inicialização

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
Solução: https://play.golang.org/p/MkdZiDW8SQ

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
Solução: https://play.golang.org/p/QlP6nVchq-

### Na prática: exercício #3

- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
Solução: https://play.golang.org/p/PzUqGLODmW

### Na prática: exercício #4

- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
Solução: https://play.golang.org/p/zCU9QWkzJ9

### Na prática: exercício #5

- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
Solução: https://play.golang.org/p/zcEsXqnBr8

### Na prática: exercício #6

- Crie um programa que demonstre o funcionamento da declaração if.
Solução: https://play.golang.org/p/OGPgTJZbiy

### Na prática: exercício #7

- Utilizando a solução anterior, adicione as opções else if e else.
Solução: https://play.golang.org/p/_cO7E-yL0o

### Na prática: exercício #8

- Crie um programa que utilize a declaração switch, sem switch statement especificado.
Solução: https://play.golang.org/p/TyIGp4Hi8B

### Na prática: exercício #9

- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
Solução: https://play.golang.org/p/4-iTPZwfEz

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

### Slice: literal composta

### Slice: for range

### Slice: fatiando uma fatia

### Slice: anexando a uma slice

### Slice: deletando de uma slice

### Slice: make

### Slice: slice multi-dimensional

### Slice: o array subjacente

### Maps: introdução

### Maps: adicionando elementos & range

### Maps: deletando



## 09 – Exercícios: Ninja Nível 4

### Na prática: exercício #1

### Na prática: exercício #2

### Na prática: exercício #3

### Na prática: exercício #4

### Na prática: exercício #5

### Na prática: exercício #6

### Na prática: exercício #7

### Na prática: exercício #8

### Na prática: exercício #9

### Na prática: exercício #10



## 10 – Structs

### Struct

### Structs embutidos

### Lendo a documentação

### Structs anônimos

### Colocando ordem na casa



## 11 – Exercícios: Ninja Nível 5

### Na prática: exercício #1

### Na prática: exercício #2

### Na prática: exercício #3

### Na prática: exercício #4



## 12 – Funções

### Sintaxe

### Parâmetro variádico

### Desenrolando uma slice

### Defer

### Métodos

### Interfaces & polimorfismo

### Funções anônimas

### Func como expressão

### Retornando uma função

### Callback

### Closure

### Recursividade



## 13 – Exercícios: Ninja Nível 6

### Na prática: exercício #1

### Na prática: exercício #2

### Na prática: exercício #3

### Na prática: exercício #4

### Na prática: exercício #5

### Na prática: exercício #6

### Na prática: exercício #7

### Na prática: exercício #8

### Na prática: exercício #9

### Na prática: exercício #10

### Na prática: exercício #11

### Na prática: exercício #12



## 14 – Ponteiros

### O que são ponteiros?

### Quando usar ponteiros

### Conjuntos de métodos



## 15 – Exercícios: Ninja Nível 7

### Na prática: exercício #1

### Na prática: exercício #2



## 16 – Aplicação

### Documentação JSON

### JSON marshal (ordenação)

### JSON unmarshal (desordenação)

### A interface Writer

### O pacote sort

### Customizando o sort

### bcrypt



## 17 – Exercícios: Ninja Nível 8

### Na prática: exercício #1

### Na prática: exercício #2

### Na prática: exercício #3

### Na prática: exercício #4

### Na prática: exercício #5



## 18 – Concorrência

### Concorrência vs. paralelismo

### WaitGroup

### Conjuntos de métodos, parte 2

### Documentação

### Condição de corrida

### Mutex

### Atomic



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

### Na prática: exercício #2

### Na prática: exercício #3

### Na prática: exercício #4

### Na prática: exercício #5

### Na prática: exercício #6

### Na prática: exercício #7



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

### Na prática: exercício #2

### Na prática: exercício #3

### Na prática: exercício #4

### Na prática: exercício #5

### Na prática: exercício #6

### Na prática: exercício #7



## 23 – Tratamento de Erros

### Entendendo erros

### Verificando erros

### Print & log

### Recover

### Erros com informações adicionais



## 24 – Exercícios: Ninja Nível 11

### Na prática: exercício #1

### Na prática: exercício #2

### Na prática: exercício #3

### Na prática: exercício #4

### Na prática: exercício #5



## 25 – Escrevendo Documentação

### Introdução

### go doc

### godoc

### godoc.org

### Escrevendo documentação



## 26 – Exercícios: Ninja Nível 12

### Na prática: exercício #1

### Na prática: exercício #2

### Na prática: exercício #3



## 27 – Testes & Benchmarking

### Introdução

### Testes em tabela

### Testes como exemplos

### Golint

### Benchmark

### Cobertura

### Exemplos de benchmarks

### Revisão



## 28 – Exercícios: Ninja Nível 13

### Na prática: exercício #1

### Na prática: exercício #2

### Na prática: exercício #3



## 29 – Considerações Finais

 