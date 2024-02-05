package product

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

func (b *Base) GenerateID()
func (b *Base) SetCreatedAt()
func (b *Base) SetUpdatedAt()
func (b *Base) TableName() string
func (b *Base) GetMap() map[string]interface{}
func (b *Base) GetFildetId() map[string]interface{}
