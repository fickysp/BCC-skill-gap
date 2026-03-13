package entity

import (
	"github.com/google/uuid"
)

type Skill struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid"`
	CareerId      uuid.UUID `gorm:"type:uuid;not null"`
	Career        Career
	Name          string    `gorm:"type:varchar(255);not null"`
	Desc          string    `gorm:"type:text"`
	Priority      int       `gorm:"not null"`
	RequiredLevel LevelEnum `gorm:"type:varchar(50);not null"`
}

func (s *Skill) BeforeCreate() error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}
