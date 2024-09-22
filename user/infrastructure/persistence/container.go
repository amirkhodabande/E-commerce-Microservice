package persistence

import (
	"fmt"
	"reflect"

	"github.com/ecommerce/user/domain/contracts"
)

type Container struct {
	bindings map[reflect.Type]interface{}
}

func NewContainer() *Container {
	return &Container{
		bindings: map[reflect.Type]interface{}{
			reflect.TypeOf((*contracts.UserRepository)(nil)): NewSqlUserRepository(),
		},
	}
}

func (c *Container) Resolve(interfaceType reflect.Type) (interface{}, error) {
	if implementation, ok := c.bindings[interfaceType]; ok {
		return implementation, nil
	}

	return nil, fmt.Errorf("no binding found for type %v", interfaceType)
}
