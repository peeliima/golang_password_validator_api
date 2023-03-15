
# Password Validator API

Uma api de validação de senhas de acordo com regras utilizando Golang e GraphQL


## Documentação da API

#### Retorna todos os itens


| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `password` | `string` | **Obrigatório**. Password que deseja testar |
| `rules`    | `input`  | Regras que deseja testar |
|  `rule`    |  `string` | Nome da regra |
|  `value`   | `int`  | Quantidade de caracteres da regra que deve ser testado |

# Payload example
mutation {
  validatorPassword(input: {
    password: "qwertQWERT!@#012",
    rules: [ 
      {
        rule: "minSize",
        value: 2
      },
      {
        rule: "minUppercase",
        value: 3
      },
      {
        rule: "minLowercase",
        value: 3
      },
      {
        rule: "minSpecialChars",
        value: 3
      },
      {
        rule: "minDigit",
        value: 3
      },
      {
        rule: "noRepeted",
        value: 3
      }
    ]
  })
}

## Aprendizados

Foi colocado em prática a sintaxe da linguagem Golang utilizando o GraphQL, também foi desenvolvido testes unitários para as funcionalidades da API utilizando a biblioteca nativa do Golang.

A utilização de algumas funcionalidades do Golang como go channels foram implementados apenas por motivos academicos.

## Estudos futuros e pontos de melhorias

#### - Primeiro ponto de melhoria: corrigir as falhas relacionadas a Regex que estão comentadas no código, não perdi muito tempo com elas pra não ficar travado no estudo, quando eu entender como o golang lida com Regex eu volto pra corrigir.

#### - Segundo ponto de melhoria: a arquitetura da api pode ser melhor organizada, preciso entender melhor as possibilidades das stacks (golang e graphql).

#### - Terceiro ponto de melhoria: A API foi desenvolvida durante os estudando de Golang, porem estou aprofundando os estudos no GraphQL, percebo que utilizei algumas coisas de forma errônea, voltarei quando tiver mais acostumado com esse modelo (Utilização de mutation em vez de query, o generated da lib graphgen está gerando um arquivo novo com nome generico, quando tento remover a query Password gera muitos conflitos, não consigo/não sei se na implementação é possível escolher os campos que retorna já que retorna de forma obrigatória um array).

#### - Quarto ponto de melhoria: Dockerizar esse código com o imagem que eu estou montando, porem ela possui problemas que vou pedir ajuda/pesquisar para resolver e utilizar ela para esse projeto, os problemas da imagem estão descritos no readme dela.

link do Dockerfile: https://github.com/peeliima/golang-ubuntu-docker

### OBS: Irei adicionar nessa sessão todos os feedbacks que recebi durante os reviews dos membros do meu squad para documentar a evolução do meu conhecimento e das correções do código
