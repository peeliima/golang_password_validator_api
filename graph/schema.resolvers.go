package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"golang_password_validator_api/graph/model"
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
