package product

import (
	"dynamodb-crud/internal/entities"
	"encoding/json"
)

type Product struct {
	entities.Base
	Name string `jason:"name"`
}

func InterfaceToModel(data interface{}) (instance *Product, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return instance, err
	}

	return instance, json.Unmarshal(bytes, &instance)
}

func (p *Product) GetFildetId() map[string]interface{}

func (p *Product) TableName() string

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{}

func ParseDynamoAttributeToStruct() {

}
