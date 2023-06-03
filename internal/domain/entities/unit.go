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
		err = errors.ErrInvalidIdWrap(u.ID)
	case u.ID == 0:
		err = errors.ErrNullIdWrap(u.ID)
	case u.ModelID < 0:
		err = errors.ErrInvalidReferenceWrap(u.ModelID)
	case u.ModelID == 0:
		err = errors.ErrNullReferenceWrap(u.ModelID)
	case u.RegDate == "":
		err = errors.ErrInvalidDateWrap(u.RegDate)
	case u.Text == "":
		err = errors.ErrInvalidContentWrap(u.Text)
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
