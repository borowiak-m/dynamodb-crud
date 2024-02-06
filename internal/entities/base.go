package entities

import (
	"time"

	"github.com/google/uuid"
)

type Interface interface {
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
	TableName() string
	GetMap() map[string]interface{}
	GetFildetId() map[string]interface{}
}

type Base struct {
	ID        uuid.UUID `json:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (b *Base) GenerateID() {
	b.ID = uuid.New()
}

func (b *Base) SetCreatedAt() {
	b.CreatedAt = time.Now()
}
func (b *Base) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}

func GetTimeFormat() string {
	return "2010-01-05T15:04:05-0700"
}

//func (b *Base) TableName() string
//func (b *Base) GetMap() map[string]interface{}
//func (b *Base) GetFildetId() map[string]interface{}
