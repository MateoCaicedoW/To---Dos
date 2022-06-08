package models

import (
	"log"
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
			output: "FirstName can not be empty.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "",
					Level:             80,
					Age:               32,
					Position:          "Forward",
					PhysicalCondition: "A+",
				},
			},
			output: "LastName can not be empty.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "Caicedo",
					Level:             100,
					Age:               32,
					Position:          "Forward",
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
					Position:          "Forward",
					PhysicalCondition: "A+",
				},
			},
			output: "FirstName can not contains caracters.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo3",
					LastName:          "Caicedo",
					Level:             99,
					Age:               32,
					Position:          "Forward",
					PhysicalCondition: "A+",
				},
			},
			output: "FirstName can not be a number.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "Caicedo",
					Level:             99,
					Age:               32,
					Position:          "Forward",
					PhysicalCondition: "A+",
					Teams: []Team{
						{},
					},
				},
			},
			output: "",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "Caicedo",
					Level:             99,
					Age:               32,
					Position:          "Forward",
					PhysicalCondition: "A+",
					Teams: []Team{
						{Name: "junior"},
						{Name: "junior"},
						{Name: "junior"},
					},
				},
			},
			output: "Teams can not be greater than 2.",
		},
		{
			input: []Player{
				{
					ID:                uuid.New(),
					FirstName:         "Mateo",
					LastName:          "Caicedo",
					Level:             99,
					Age:               32,
					Position:          "Forward",
					PhysicalCondition: "A+",
					Teams: []Team{
						{Name: "junior"},
						{Name: "junior"},
					},
				},
			},
			output: "Teams can not be the same.",
		},
	}

	for i, test := range tests {
		for _, player := range test.input {
			response := player.Validate()
			if response.Message != test.output {
				t.Errorf("Test %d: Expected %s, got %s", i, test.output, response.Message)
				//t.Errorf("Expected %s, got %s", test.output, response.Message)
			}
			log.Printf("Test %d: %v", i, response)
		}
	}

}
