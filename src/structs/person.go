package structs

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"nome" gorm:"type:varchar(32);not null,index:idx_name"`
	Nickname  string    `json:"apelido" gorm:"type:varchar(100);unique;not null,index:idx_nickname"`
	BirthDate time.Time `json:"nascimento"`
	Stack     []Stack   `json:"stack" gorm:"foreignKey:PersonID"`
}
