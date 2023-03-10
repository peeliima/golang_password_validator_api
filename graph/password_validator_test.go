package graph

import (
	"context"
	"golang_password_validator_api/graph/model"
	"testing"
)

func TestValidatorPassword(t *testing.T) {
	context := context.Background()

	resolver := Resolver{}

	input_valid := model.ValidatorPasswordInput{
		Password: "qwertQWERT!@#012",
		Rules: []*model.RuleInput{
			{Rule: "minSize", Value: 2},
			{Rule: "minUppercase", Value: 3},
		},
	}

	response_success := []string{"Senha Válida!"}

	response, _ := resolver.Mutation().ValidatorPassword(context, input_valid)

	if response[0] != response_success[0] {
		t.Errorf("Error")
	}
}
