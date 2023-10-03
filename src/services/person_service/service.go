package person_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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
	person, err := s.repository.GetByID(ctx, ID)
	if err != nil {
		return person_dto.Info{}, err
	}

	var stacks []string
	for _, stack := range person.Stack {
		stacks = append(stacks, stack.Name)
	}

	return person_dto.Info{
		ID:        person.ID,
		Name:      person.Name,
		Nickname:  person.Nickname,
		BirthDate: person.BirthDate.String(),
		Stack:     &stacks,
	}, nil
}

func (s *Service) List(ctx context.Context, params person_dto.ListRequestParams) ([]person_dto.Info, error) {
	persons, err := s.repository.List(ctx, params)
	if err != nil {
		return []person_dto.Info{}, err
	}

	personsDTO := make([]person_dto.Info, 0)
	for _, p := range persons {
		var stacks []string
		for _, stack := range p.Stack {
			stacks = append(stacks, stack.Name)
		}

		personsDTO = append(personsDTO, person_dto.Info{
			ID:        p.ID,
			Name:      p.Name,
			Nickname:  p.Nickname,
			BirthDate: p.BirthDate.String(),
			Stack:     &stacks,
		})
	}

	return personsDTO, nil
}

func (s *Service) Create(ctx context.Context, payload person_dto.CreatePayload) (person_dto.Info, error) {
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

	var stacks []string
	for _, stack := range person.Stack {
		stacks = append(stacks, stack.Name)
	}

	return person_dto.Info{
		ID:        person.ID,
		Name:      person.Name,
		Nickname:  person.Nickname,
		BirthDate: person.BirthDate.String(),
		Stack:     &stacks,
	}, nil
}
