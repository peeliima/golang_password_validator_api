package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"golang_password_validator_api/graph/model"
	"regexp"
	"unicode"
	"unicode/utf8"
)

// ValidatorPassword is the resolver for the validatorPassword field.
func (r *mutationResolver) ValidatorPassword(ctx context.Context, input model.ValidatorPasswordInput) ([]string, error) {
	password := input.Password
	valid_password := false

	errors := []string{}

	for _, rules := range input.Rules {
		switch rules.Rule {
		case "minSize":
			valid_password = minSize(password, rules.Value)

			if !valid_password {
				errors = append(errors, "minSize")
			}
		case "minUppercase":
			valid_password = minUppercase(password, rules.Value)

			if !valid_password {
				errors = append(errors, "minUppercase")
			}
		case "minLowercase":
			valid_password = minLowercase(password, rules.Value)

			if !valid_password {
				errors = append(errors, "minLowercase")
			}
		case "minSpecialChars":
			valid_password = minSpecialChars(password, rules.Value)

			if !valid_password {
				errors = append(errors, "minSpecialChars")
			}
		case "minDigit":
			valid_password = minDigit(password, rules.Value)

			if !valid_password {
				errors = append(errors, "minDigit")
			}
		case "noRepeted":
			valid_password = noRepeted(password, rules.Value)

			if !valid_password {
				errors = append(errors, "noRepeted")
			}
		}
	}

	if len(errors) > 0 {
		return errors, nil
	}

	errors = append(errors, "Senha Válida!")
	return errors, nil
}

func minSize(password string, value int) bool {
	count := utf8.RuneCountInString(password)

	return count >= value
}

func minUppercase(password string, value int) bool {
	count := 0

	for _, password_splited := range password {
		if unicode.IsUpper(password_splited) {
			count++
		}
	}

	return count >= value
}

func minLowercase(password string, value int) bool {
	count := 0

	for _, password_splited := range password {
		if unicode.IsLower(password_splited) {
			count++
		}
	}

	return count >= value
}

func minDigit(password string, value int) bool {
	re := regexp.MustCompile("[0-9]")
	digits := re.FindAllString(password, -1)

	return len(digits) >= value
}

// Essa func precisa de um ajuste no Regex||Codigo, está com problemas na validação dos contra barras \ nas senhas
// Preciso entender como resolver problemas vinculados a forma que o Golang lida com \
// Usar \\ gera o scape do contra barra mas continua sendo problematico em tempo de excução, crase tambem nao resolveu o problema
// Provavelmente o problema está na funcao que estou utilizando para validar esse caso
func minSpecialChars(password string, value int) bool {
	re := regexp.MustCompile("[!@#$%^&*()-\\/+{}\\[\\]]")
	special_caracteres := re.FindAllString(password, -1)

	return len(special_caracteres) >= value
}

// Essa func eu queria implementar utilizando Regex porem não conseguir utilizar as funcoes do Go corretamente
// pra não perder muito tempo fiz da maneira abaixo, futuramente volto nesse ponto para implementar da maneira mais elegante
// Att: Refiz a funcao utilizando Regex, mas vou deixar essa func no código por motivos academicos
func minSpecialCharsDepreciated(password string, value int) bool {
	count := 0
	special_caracteres := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '\\', '/', '+', '{', '}'}

	for _, pass := range password {
		for _, char := range special_caracteres {
			if pass == char {
				count++
			}
		}
	}

	return count >= value
}

// Pelo site https://regexr.com/ o regex a seguir -> /([a-zA-Z0-9])\1+/g resolve o problema de caracteres normais e numericos repetidos
// Porem o Golang não interpreta a string literal com o contra barra mesmo usando crases (posso ter utilizado da maneira errada), depois eu volto nessa funcao, mesmo problema da func de caracteres especiais
// Provavelmente o problema está na funcao que estou utilizando para validar esse caso
func noRepeted(password string, value int) bool {
	regex := "[a-zA-Z0-9]"
	re := regexp.MustCompile(regex)

	// re := regexp.MustCompile("/^([a-z])\1+$/")
	caracteres_repeted := re.FindAllString(password, -1)

	fmt.Println(caracteres_repeted)

	return true

}

// Password is the resolver for the password field.
func (r *queryResolver) Password(ctx context.Context, password string) (*model.Password, error) {
	fmt.Println(password)
	return nil, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
