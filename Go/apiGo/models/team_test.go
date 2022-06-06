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
					ID:      uuid.New(),
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
					ID:      uuid.New(),
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
					ID:      uuid.New(),
					Name:    "Colombia#",
					Type:    "National",
					Country: "",
				},
			},
			output: "Name cant not contains caracters.",
		},
		{
			input: []Team{
				{
					ID:      uuid.New(),
					Name:    "Colombia",
					Type:    "National$",
					Country: "",
				},
			},
			output: "Type cant not contains caracters.",
		},
		{
			input: []Team{
				{
					ID:      uuid.New(),
					Name:    "Colombia",
					Type:    "National",
					Country: "Colombia",
				},
			},
			output: "Country must be empty if Type is National.",
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
