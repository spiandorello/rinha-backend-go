package person_dto

import "github.com/google/uuid"

type ListRequestParams struct {
	Params string `json:"params"`
	Size   int    `json:"size"`
}

type Info struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"nome"`
	Nickname  string    `json:"apelido"`
	BirthDate string    `json:"nascimento"`
	Stack     *[]string `json:"stack"`
}
