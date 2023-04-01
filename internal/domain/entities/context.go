package entities

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"strings"
	"time"
)

type Context struct {
	ID      int    `db:"id"`
	RegDate string `db:"registration_date"`
	Text    string `db:"text"`
}

func (c *Context) IsValid() error {
	var err error
	switch {
	case c.ID < 0:
		err = errors.ErrInvalidID
	case c.ID == 0:
		err = errors.ErrNullID
	case c.RegDate == "":
		err = errors.ErrInvalidDate
	case c.Text == "":
		err = errors.ErrInvalidContent
	}

	return err
}

func (c *Context) Contains(other *Context) bool {
	return strings.Contains(c.Text, other.Text)
}

func (c *Context) IsDuplicate(other *Context) bool {
	return strings.Contains(c.Text, other.Text) || strings.Contains(other.Text, c.Text)
}

func (c *Context) ContainsUnit(unit *Unit) bool {
	return strings.Contains(c.Text, unit.Text)
}

func (c *Context) SetRegDate(t time.Time) {
	c.RegDate = t.Format(entityDateFormat)
}

func (c *Context) GetRegDate() (time.Time, error) {
	return time.Parse(entityDateFormat, c.RegDate)
}
