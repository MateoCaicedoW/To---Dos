package models

import (
	"log"
	"testing"

	"github.com/google/uuid"
)

type listTest struct {
	input  []Team
	output string
}

func Test_validateTeam(t *testing.T) {

	var tests = []listTest{
		{
			input: []Team{
				{
					IDTeam:  uuid.New(),
					Name:    "",
					Type:    "",
					Country: "",
				},
			},
			output: "Name cant not be empty.",
		},
		{
			input: []Team{
				{
					IDTeam:  uuid.New(),
					Name:    "Junior",
					Type:    "Club",
					Country: "",
				},
			},
			output: "Country cant not be empty if Type is Club.",
		},
		{
			input: []Team{
				{
					IDTeam:  uuid.New(),
					Name:    "Colombia#",
					Type:    "Seleccion",
					Country: "",
				},
			},
			output: "Name cant not contains caracters.",
		},
		{
			input: []Team{
				{
					IDTeam:  uuid.New(),
					Name:    "Colombia",
					Type:    "Seleccion$",
					Country: "",
				},
			},
			output: "Type cant not contains caracters.",
		},
		{
			input: []Team{
				{
					IDTeam:  uuid.New(),
					Name:    "Colombia",
					Type:    "Seleccion",
					Country: "Colombia",
				},
			},
			output: "Country must be empty if Type is Seleccion.",
		},
	}

	for i, test := range tests {
		for _, team := range test.input {
			err := team.Validate()
			if err.Message != test.output {
				t.Error(err)
			}
			log.Printf("Test %d: %v", i, err)
		}
	}

}
