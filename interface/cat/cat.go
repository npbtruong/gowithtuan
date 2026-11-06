package cat

import (
	"errors"
	"strings"
)

// cái gì muốn export cho file khác gọi viết hoa chử đầu
type Cat struct {
	Name  string
	Owner string
}

func New(name string, owner string) (*Cat, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("name is too long (max 50 characters)")
	}

	return &Cat{
		Name:  name,
		Owner: owner,
	}, nil
}

func (c *Cat) EatSound() string {
	return "Crm Crm Crm!"
}

func (c *Cat) Sound() string {
	return "MOEW MOEW"
}

func (c *Cat) GetName() string {
	return c.Name
}
