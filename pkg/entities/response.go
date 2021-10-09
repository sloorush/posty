package entities

type Error struct {
	Status  int
	Message string
}

type Success struct {
	Status  int
	Message string
	Data    *User
}
