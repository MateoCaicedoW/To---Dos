package models

import (
	"testing"

	"github.com/google/uuid"
)

type tests struct {
	input  []PlayerTeam
	output string
}

func Test_validatePlayerTeam(t *testing.T) {
	var tests = []tests{
		{
			input: []PlayerTeam{
				{
					PlayerID: uuid.New(),
					TeamID:   uuid.New(),
				},
			},
			output: "",
		},
		{
			input: []PlayerTeam{
				{
					PlayerID: uuid.Nil,
					TeamID:   uuid.New(),
				},
			},
			output: "PlayerID can not be empty.",
		},
		{
			input: []PlayerTeam{
				{
					PlayerID: uuid.New(),
					TeamID:   uuid.Nil,
				},
			},
			output: "TeamID can not be empty.",
		},
	}

	for _, test := range tests {
		for _, playerTeam := range test.input {
			response := playerTeam.Validate()
			if response.Message != test.output {
				t.Errorf("Expected %s, got %s", test.output, response.Message)
			}
		}

	}

}
