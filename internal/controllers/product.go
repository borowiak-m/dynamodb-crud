package product

import (
	"dynamodb-crud/internal/entities/product"
	"dynamodb-crud/internal/repository/adapter"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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
	response, err := c.repository.FindOne(entity.GetFilterId(), entity.TableName())
	if err != nil {
		return entity, err
	}
	return product.ParseDynamicAttributeToStruct(response.Item)
}

func (c *Controller) ListAll() (entities []product.Product, err error) {
	entities = []product.Product{}
	var entity product.Product

	filter := expression.Name("name").NotEqual(expression.Value(""))

	condition, err := expression.NewBuilder().WithFilter(filter).Build()

	if err != nil {
		return entities, err
	}

	response, err := c.repository.FindAll(condition, entity.TableName())
	if err != nil {
		return entities, err
	}

	if response != nil {
		for _, value := range response.Items {
			entity, err := product.ParseDynamoAttributeToStruct(value)
			if err != nil {
				return entities, err
			}
			entities = append(entities, entity)
		}
	}
	return entities, nil
}

func (c *Controller) Create(entity product.Product) (uuid.UUID, error) {
	entity.CreatedAt = time.Now()
	_, err := c.repository.CreateOrUpdate(entity.GetMap(), entity.TableName())
	return entity.ID, err
}

func (c *Controller) Update(ID uuid.UUID, entity *product.Product) error {
	found, err := c.ListOne(ID)
	if err != nil {
		return err
	}

	found.ID = ID
	found.Name = entity.Name
	found.UpdatedAt = time.Now()
	_, err = c.repository.CreateOrUpdate(found.GetMap(), entity.TableName())
	return err
}

func (c *Controller) Remove(ID uuid.UUID) error {
	entity, err := c.ListOne(ID)
	if err != nil {
		return err
	}
	_, err = c.repository.Delete(entity.GetFildetId(), entity.TableName())
	return err
}
