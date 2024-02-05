package product

import (
	"dynamodb-crud/internal/entities/product"
	"dynamodb-crud/internal/repository/adapter"

	"github.com/google/uuid"
)

type Controller struct {
	repository adapter.Interface
}

type Iterface interface {
	ListOne(ID uuid.UUID) (entity product.Product, err error)
	ListAll() (entities []product.Product, err error)
	Create(entity product.Product) (uuid.UUID, error)
	Update(ID uuid.UUID, entity product.Product) error
	Remove(ID uuid.UUID) error
}

func NewController(repository adapter.Interface) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) ListOne(ID uuid.UUID) (entity product.Product, err error) {
	entity.ID = ID
	response, err := c.repository.FindOne(entity.GetFilerId(), entity.TableName)
	if err != nil {
		return entity, err
	}
	return product.ParseDynamicAttributeToStruct(response.Item)
}

func ListAll() (entities []product.Product, err error) {

}

func (c *Controller) Create(entity product.Product) (uuid.UUID, error) {

}

func (c *Controller) Update(ID uuid.UUID, entity *product.Product)

func (c *Controller) Remove(ID uuid.UUID) error {

}
