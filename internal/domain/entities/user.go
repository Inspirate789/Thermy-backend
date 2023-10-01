package entities

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"hash/fnv"
	"time"
)

const entityDateFormat = time.DateTime

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Role     string `db:"role"`
	RegDate  string `db:"registration_date"`
}

func NewUser(request *AuthRequest) *User {
	return &User{
		ID:       0,
		Name:     request.Username,
		Password: request.Password,
		Role:     "",
		RegDate:  "",
	}
}

func (u *User) IsValid() error {
	var err error
	switch {
	case u.ID < 0:
		err = errors.ErrInvalidIdWrap(u.ID)
	case u.ID == 0:
		err = errors.ErrNullIdWrap(u.ID)
	case u.Name == "":
		return errors.ErrInvalidNameWrap(u.Name)
	case u.Password == "":
		return errors.ErrInvalidPasswordWrap(u.Password)
	case u.Role == "":
		return errors.ErrInvalidRoleWrap(u.Role)
	case u.RegDate == "":
		err = errors.ErrInvalidDateWrap(u.RegDate)
	}

	return err
}

func (u *User) GetHash() (uint64, error) {
	h := fnv.New64a()

	_, err := h.Write([]byte(u.Name))
	if err != nil {
		return 0, err
	}

	return h.Sum64(), err
}

func (u *User) SetRegDate(t time.Time) {
	u.RegDate = t.Format(entityDateFormat)
}

func (u *User) GetRegDate() (time.Time, error) {
	return time.Parse(entityDateFormat, u.RegDate)
}
