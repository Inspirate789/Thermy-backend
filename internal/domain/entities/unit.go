package entities

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"time"
)

type Unit struct {
	ID      int    `db:"id"`
	ModelID int    `db:"model_id"`
	RegDate string `db:"registration_date"`
	Text    string `db:"text"`
}

type UnitsMap map[string][]Unit

func (u Unit) IsValid() error {
	var err error
	switch {
	case u.ID < 0:
		err = errors.ErrInvalidID
	case u.ID == 0:
		err = errors.ErrNullID
	case u.ModelID < 0:
		err = errors.ErrInvalidReference
	case u.ModelID == 0:
		err = errors.ErrNullReference
	case u.RegDate == "":
		err = errors.ErrInvalidDate
	case u.Text == "":
		err = errors.ErrInvalidContent
	}

	return err
}

func (u Unit) SetModel(model *Model) error {
	err := model.IsValid()
	if err != nil {
		return err
	}

	u.ModelID = model.ID

	return nil
}

func (u Unit) SetRegDate(t time.Time) {
	u.RegDate = t.Format(entityDateFormat)
}

func (u Unit) GetRegDate() (time.Time, error) {
	return time.Parse(entityDateFormat, u.RegDate)
}

func (u Unit) GetNuclearElem() (*Unit, error) {
	return nil, errors.ErrFeatureNotExistWrap("GetNuclearElem")
}
