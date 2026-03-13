package entity

import (
	"time"

	"github.com/google/uuid"
)

type QuizSession struct {
	ID                  uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserID              uuid.UUID `gorm:"type:uuid;not null"`
	User                User
	UserCareerSessionID uuid.UUID `gorm:"type:uuid;not null"`
	UserCareerSession   UserCareerSession
	SkillID             uuid.UUID `gorm:"type:uuid;not null"`
	Skill               Skill
	Status              StatusEnum `gorm:"type:varchar(50);default:'on_process'"`
	Score               float64    `gorm:"type:float8"`
	StartedAt           time.Time  `gorm:"autoCreateTime"`
	CompletedAt         *time.Time
}

func (q *QuizSession) BeforeCreate() error {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	return nil
}
