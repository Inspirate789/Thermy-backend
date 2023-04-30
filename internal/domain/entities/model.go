package entities

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"regexp"
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

func (m *Model) IsValidName() error {
	matched, _ := regexp.MatchString(`^([a-z]+\+)*?[a-z]$`, m.Name)

	if !matched {
		return errors.ErrInvalidNameWrap(m.Name)
	}
	return nil
}

func (m *Model) IsValid() error {
	err := m.IsValidName()

	switch {
	case err != nil:
		break
	case m.ID < 0:
		err = errors.ErrInvalidIdWrap(m.ID)
	case m.ID == 0:
		err = errors.ErrNullIdWrap(m.ID)
	case m.Name == "":
		err = errors.ErrInvalidContentWrap(m.Name)
	}

	return err
}

func (e *ModelElement) IsValid() error {
	var err error
	switch {
	case e.ID < 0:
		err = errors.ErrInvalidIdWrap(e.ID)
	case e.ID == 0:
		err = errors.ErrNullIdWrap(e.ID)
	case e.Name == "":
		err = errors.ErrInvalidContentWrap(e.ID)
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
