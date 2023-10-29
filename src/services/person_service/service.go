package person_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"rinha-de-backend/pkg/opentelemetry"
	"rinha-de-backend/src/dtos/person_dto"
	"rinha-de-backend/src/interfaces"
)

type Service struct {
	repository interfaces.PersonRepository
}

func New(repository interfaces.PersonRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetByID(ctx context.Context, ID uuid.UUID) (person_dto.Info, error) {
	ctx, parentSpan := opentelemetry.Tracer.Start(ctx, "service:person:get-by-id")
	defer parentSpan.End()

	person, err := s.repository.GetByID(ctx, ID)
	if err != nil {
		return person_dto.Info{}, err
	}

	stacks := make([]string, len(person.Stack), len(person.Stack))
	for key, stack := range person.Stack {
		stacks[key] = stack.Name
	}

	return person_dto.Info{
		ID:        person.ID,
		Name:      person.Name,
		Nickname:  person.Nickname,
		BirthDate: person.BirthDate.String(),
		Stack:     stacks,
	}, nil
}

func (s *Service) List(ctx context.Context, params person_dto.ListRequestParams) ([]person_dto.Info, error) {
	ctx, parentSpan := opentelemetry.Tracer.Start(ctx, "service:person:list")
	defer parentSpan.End()

	persons, err := s.repository.List(ctx, params)
	if err != nil {
		return []person_dto.Info{}, err
	}

	personsDTO := make([]person_dto.Info, len(persons), len(persons))
	for key, p := range persons {
		stacks := make([]string, len(p.Stack), len(p.Stack))
		for key, stack := range p.Stack {
			stacks[key] = stack.Name
		}

		personsDTO[key] = person_dto.Info{
			ID:        p.ID,
			Name:      p.Name,
			Nickname:  p.Nickname,
			BirthDate: p.BirthDate.String(),
			Stack:     stacks,
		}
	}

	return personsDTO, nil
}

func (s *Service) Create(ctx context.Context, payload person_dto.CreatePayload) (person_dto.Info, error) {
	ctx, parentSpan := opentelemetry.Tracer.Start(ctx, "service:person:create")
	defer parentSpan.End()

	person, err := s.repository.GetByNickname(ctx, payload.Nickname)
	if err != nil {
		return person_dto.Info{}, err
	}

	if person.ID != uuid.Nil {
		return person_dto.Info{}, fmt.Errorf("person already exists")
	}

	person, err = s.repository.Save(ctx, payload)
	if err != nil {
		return person_dto.Info{}, err
	}

	stacks := make([]string, len(person.Stack), len(person.Stack))
	for key, stack := range person.Stack {
		stacks[key] = stack.Name
	}

	return person_dto.Info{
		ID:        person.ID,
		Name:      person.Name,
		Nickname:  person.Nickname,
		BirthDate: person.BirthDate.String(),
		Stack:     stacks,
	}, nil
}
