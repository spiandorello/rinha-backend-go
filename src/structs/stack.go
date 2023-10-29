package structs

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stack struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"nome" gorm:"type:varchar(32);not null,index:idx_name"`
	PersonID uuid.UUID `json:"person_id" gorm:"type:uuid;not null"`
}

func (s *Stack) BeforeCreate(tx *gorm.DB) error {
	if s.ID != uuid.Nil {
		return nil
	}
	s.ID = uuid.New()
	return nil
}
