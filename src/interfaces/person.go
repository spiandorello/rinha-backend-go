package interfaces

import (
	"context"

	"github.com/google/uuid"
	"rinha-de-backend/src/dtos/person_dto"
	"rinha-de-backend/src/structs"
)

type PersonService interface {
	GetByID(ctx context.Context, ID uuid.UUID) (person_dto.Info, error)
	List(ctx context.Context, params person_dto.ListRequestParams) ([]person_dto.Info, error)
	Create(ctx context.Context, payload person_dto.CreatePayload) (person_dto.Info, error)
}

type PersonRepository interface {
	GetByID(ctx context.Context, ID uuid.UUID) (structs.Person, error)
	GetByNickname(ctx context.Context, nickname string) (structs.Person, error)
	List(ctx context.Context, params person_dto.ListRequestParams) ([]structs.Person, error)
	Save(ctx context.Context, payload person_dto.CreatePayload) (structs.Person, error)
}
