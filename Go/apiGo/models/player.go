package models

type Player struct {
	ID        string
	FirstName string
	LastName  string
	Level     int64
}

type ListPlayers []Player

func Players() ListPlayers {
	P := ListPlayers{
		Player{"deae98b1-feab-47d0-a64b-5808bb12d612", "Juan", "Hernandez", 12},
		Player{"9a42e66f-f020-471b-893e-2e04c2e24307", "Marcos", "Llorente", 90},
		Player{"07ec6cb3-71fa-458d-bd0a-afcf1f7c3e5d", "Luis", "Avila", 82},
		Player{"18207fe0-c9af-4b98-a1cb-494d874450f0", "Alberto", "Mercado", 20},
		Player{"5e98f402-d087-4134-9a01-59af9af0f6d1", "Rosa", "Mercado", 30},
		Player{"dba0f115-c8e6-401f-a14b-9e47a01367b3", "Marta", "Rosario", 30},
		Player{"f86166fc-1481-44ee-8435-e1e78529d25d", "Isaias", "Perez", 40},
		Player{"642da712-0eb2-4e8f-a0b9-5f6f597eaaaf", "Samuel", "Benitez", 45},
		Player{"1b82526c-1e7b-46d6-a47b-06e8e15784e4", "Gonzalo", "Higuain", 25},
		Player{"ae205ca3-84a9-485e-8b94-949c21aac14b", "Alberto", "Rosado", 75},
		Player{"f3cf2904-0780-4e4e-bc44-0b9f8f1babb6", "Ismael", "Perez", 59},
	}
	return P

}
