package structs

import (
	"github.com/google/uuid"
)

type Stack struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"nome" gorm:"type:varchar(32);not null,index:idx_name"`
	PersonID uuid.UUID `json:"person_id" gorm:"type:uuid;not null"`
}
