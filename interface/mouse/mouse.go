package mouse

import (
	"errors"
	"go/go/service"
	"strings"
)

type Mouse struct {
	Name string
}

func New(name string) (service.AnimalAction, error) {

	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("name is too long (max 50 characters)")
	}

	return &Mouse{
		Name: name,
	}, nil

}

func (m Mouse) Sound() string {
	return "BARK BARK"
}

func (m *Mouse) GetName() string {
	return m.Name
}

func (m Mouse) Bite() string {
	return "cratch attack!!!!"
}
