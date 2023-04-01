package entities

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"strings"
)

type Property struct {
	ID   int    `db:"id"`
	Name string `db:"property"`
}

type Properties []Property

func (p *Property) IsValid() error {
	var err error
	switch {
	case p.ID < 0:
		err = errors.ErrInvalidID
	case p.ID == 0:
		err = errors.ErrNullID
	case p.Name == "":
		err = errors.ErrInvalidContent
	}

	return err
}

func (p *Property) IsDuplicate(other *Property) bool {
	return strings.Contains(p.Name, other.Name) || strings.Contains(other.Name, p.Name)
}

func (p *Property) Contains(other *Property) bool {
	return strings.Contains(p.Name, other.Name)
}

func (p Properties) Contains(other *Property) bool {
	for _, property := range p {
		if property.IsDuplicate(other) {
			return true
		}
	}

	return false
}

func (p Properties) RemoveDuplicates() Properties {
	set := make([]Property, len(p)/2)
	for _, property := range p {
		if !p.Contains(&property) {
			set = append(set, property)
		}
	}

	return set
}
