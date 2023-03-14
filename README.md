# golang_password_validator_api
Uma api de validação de senhas de acordo com regras utilizando Golang e GraphQL

# Para executar
execute o comando ´´´go run server.go´´´ na pasta raiz do projeto - go1.20.1

# Testes
vá a pasta graph e execute o comando ´´´go test´´´

# Payload example
´´´
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
´´´