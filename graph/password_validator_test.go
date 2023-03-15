package graph

import (
	"context"
	"golang_password_validator_api/graph/model"
	"testing"

	"golang.org/x/exp/slices"
)

func TestValidatorPassword(t *testing.T) {
	context := context.Background()

	resolver := Resolver{}

	input_valid := model.ValidatorPasswordInput{
		Password: "qwertQWERT!@#012",
		Rules: []*model.RuleInput{
			{Rule: "minSize", Value: 2},
			{Rule: "minUppercase", Value: 3},
			{Rule: "minLowercase", Value: 3},
			{Rule: "minSpecialChars", Value: 3},
			{Rule: "minDigit", Value: 3},
			{Rule: "noRepeted", Value: 3},
		},
	}

	response_success := []string{"Senha Válida!"}

	response, _ := resolver.Mutation().ValidatorPassword(context, input_valid)

	if response[0] != response_success[0] {
		for _, res := range response {
			t.Errorf(res)
		}
	}
}

func TestFailValidatorPassword(t *testing.T) {
	context := context.Background()

	resolver := Resolver{}

	input_failed := model.ValidatorPasswordInput{
		Password: "qwert",
		Rules: []*model.RuleInput{
			{Rule: "minSize", Value: 10},
			{Rule: "minUppercase", Value: 10},
			{Rule: "minLowercase", Value: 10},
			{Rule: "minSpecialChars", Value: 3},
			{Rule: "minDigit", Value: 10},
		},
	}

	expect_response := []string{
		"minSize",
		"minUppercase",
		"minLowercase",
		"minSpecialChars",
		"minDigit",
	}

	response, _ := resolver.Mutation().ValidatorPassword(context, input_failed)

	// fmt.Println(response, expect_response)
	for _, expect := range expect_response {
		if !slices.Contains(response, expect) {
			t.Errorf("A regra %s está com problemas", expect)
		}
	}
}
