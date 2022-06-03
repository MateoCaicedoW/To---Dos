package models

import (
	"testing"

	"github.com/google/uuid"
)

type test struct {
	input  []Player
	output string
}

func Test_validate(t *testing.T) {
	tests := []test{
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "",
					LastName:          "",
					Level:             0,
					Age:               0,
					Position:          "",
					PhysicalCondition: "",
				},
			},
			output: "FirstName cant not be empty.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "",
					Level:             80,
					Age:               32,
					Position:          "Delantero",
					PhysicalCondition: "A+",
				},
			},
			output: "LastName cant not be empty.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "Caicedo",
					Level:             100,
					Age:               32,
					Position:          "Delantero",
					PhysicalCondition: "A+",
				},
			},
			output: "Level must be between 1 and 99.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "Caicedo",
					Level:             99,
					Age:               32,
					Position:          "Sosa",
					PhysicalCondition: "A+",
				},
			},
			output: "Insert a valid Position.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "Caicedo",
					Level:             99,
					Age:               32,
					Position:          "Sosa",
					PhysicalCondition: "A+",
				},
			},
			output: "Insert a valid Position.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo$",
					LastName:          "Caicedo",
					Level:             99,
					Age:               32,
					Position:          "Delantero",
					PhysicalCondition: "A+",
				},
			},
			output: "FirstName cant not contains caracters.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo3",
					LastName:          "Caicedo",
					Level:             99,
					Age:               32,
					Position:          "Delantero",
					PhysicalCondition: "A+",
				},
			},
			output: "FirstName cant not be a number.",
		},
	}

	for _, test := range tests {
		for _, player := range test.input {
			response := player.Validate()
			if response.Message != test.output {
				t.Errorf("Expected %s, got %s", test.output, response.Message)
			}
		}
	}

}
