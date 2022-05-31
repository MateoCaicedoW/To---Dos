package models

import "github.com/google/uuid"

type Player struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Level     int64
}

func (p *Player) Validate() (result Template) {
	if p.FirstName == "" {
		result = Template{Status: 400, Data: ListPlayers{}, Message: "You inserted a Int on FirstName"}
	} else if p.LastName == "" {
		result = Template{Status: 400, Data: ListPlayers{}, Message: "You inserted a Int on LastName"}
	} else if p.Level == 0 {
		result = Template{Status: 400, Data: ListPlayers{}, Message: "You inserted a string on Level"}
	}
	return
}

type Template struct {
	Status  int
	Data    ListPlayers
	Message string
}

type ListPlayers []Player

var P = ListPlayers{
	Player{uuid.MustParse("deae98b1-feab-47d0-a64b-5808bb12d612"), "Juan", "Hernandez", 12},
	Player{uuid.MustParse("9a42e66f-f020-471b-893e-2e04c2e24307"), "Marcos", "Llorente", 90},
	Player{uuid.MustParse("07ec6cb3-71fa-458d-bd0a-afcf1f7c3e5d"), "Luis", "Avila", 82},
	Player{uuid.MustParse("18207fe0-c9af-4b98-a1cb-494d874450f0"), "Alberto", "Mercado", 20},
	Player{uuid.MustParse("5e98f402-d087-4134-9a01-59af9af0f6d1"), "Rosa", "Mercado", 30},
	Player{uuid.MustParse("dba0f115-c8e6-401f-a14b-9e47a01367b3"), "Marta", "Rosario", 30},
	Player{uuid.MustParse("f86166fc-1481-44ee-8435-e1e78529d25d"), "Isaias", "Perez", 40},
	Player{uuid.MustParse("642da712-0eb2-4e8f-a0b9-5f6f597eaaaf"), "Samuel", "Benitez", 45},
	Player{uuid.MustParse("1b82526c-1e7b-46d6-a47b-06e8e15784e4"), "Gonzalo", "Higuain", 25},
	Player{uuid.MustParse("ae205ca3-84a9-485e-8b94-949c21aac14b"), "Alberto", "Rosado", 75},
	Player{uuid.MustParse("f3cf2904-0780-4e4e-bc44-0b9f8f1babb6"), "Ismael", "Perez", 59},
}
