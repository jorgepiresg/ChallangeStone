package model_account

import (
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	tests := map[string]struct {
		input Create
		err   error
	}{
		"success": {
			input: Create{
				Name:     "John Doe",
				CPF:      "11111111111",
				Password: "any",
			},
		},
		"error: name invalid": {
			input: Create{
				Name: "",
			},
			err: fmt.Errorf("name invalid"),
		},
		"error: cpf invalid": {
			input: Create{
				Name: "John Doe",
			},
			err: fmt.Errorf("cpf invalid"),
		},
		"error: password invalid": {
			input: Create{
				Name: "John Doe",
				CPF:  "11111111111",
			},
			err: fmt.Errorf("password invalid"),
		},
		"error: cpf invalid only numbers": {
			input: Create{
				Name:     "John Doe",
				CPF:      "1111111111A",
				Password: "any",
			},
			err: fmt.Errorf("cpf invalid"),
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			err := tt.input.Valid()

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}
		})
	}

}
