package person_dto

type CreatePayload struct {
	Name      string   `json:"nome" validate:"required,min=1,max=32"`
	Nickname  string   `json:"apelido" validate:"required,min=1,max=100"`
	BirthDate string   `json:"nascimento" validate:"required,birth"`
	Stack     []string `json:"stack" validate:"dive,min=1,max=32"`
}
