package entities

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"strings"
)

const modelElementsDelimiter = "+"

type Model struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type ModelElement struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (m *Model) IsValid() error {
	var err error
	switch {
	case m.ID < 0:
		err = errors.ErrInvalidID
	case m.ID == 0:
		err = errors.ErrNullID
	case m.Name == "":
		err = errors.ErrInvalidContent
	}

	return err
}

func (e *ModelElement) IsValid() error {
	var err error
	switch {
	case e.ID < 0:
		err = errors.ErrInvalidID
	case e.ID == 0:
		err = errors.ErrNullID
	case e.Name == "":
		err = errors.ErrInvalidContent
	}

	return err
}

func (m *Model) Contains(other *Model) bool {
	return strings.Contains(m.Name, other.Name)
}

func (m *Model) IsDuplicate(other *Model) bool {
	return strings.Contains(m.Name, other.Name) || strings.Contains(other.Name, m.Name)
}

func (m *Model) ContainsElement(element *ModelElement) bool {
	return strings.Contains(m.Name, element.Name)
}

func (m *Model) AddElement(element *ModelElement) *Model {
	return &Model{
		ID:   m.ID,
		Name: m.Name + modelElementsDelimiter + element.Name,
	}
}

func (m *Model) GetNuclearElem() (*Model, error) {
	return nil, errors.ErrFeatureNotExistWrap("GetNuclearElem")
}
