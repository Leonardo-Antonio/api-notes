package user

import (
	"database/sql"
	"github.com/Leonardo-Antonio/api-notes/helper"
	"github.com/asaskevich/govalidator"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) CreateUser(user model) error {
	stmt, err := s.db.Prepare(sqlNewUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if !govalidator.IsEmail(user.Email) {
		return helper.ErrEmailInvalid
	}
	if len(user.Password) < 10 {
		return helper.ErrPasswordNotSecure
	}
	if len(user.NickName) < 5 {
		return helper.ErrNickNameIvalid
	}
	rs, err := stmt.Exec(
			helper.StringNull(user.Name),
			helper.StringNull(user.LastName),
			helper.StringNull(user.NickName),
			helper.StringNull(string(user.Profile)),
			helper.StringNull(user.Email),
			helper.StringNull(user.Password),
		)
	if err != nil {
		return helper.ErrDataInvalid
	}
	rA, err := rs.RowsAffected()
	if err != nil {
		return helper.ErrRowsAffected
	}
	if rA != 1 {
		return helper.ErrInsertionFailed
	}
	return nil
}

func (s *Storage) Login(data model) (user model, err error) {
	stmt, err := s.db.Prepare(sqlLogin)
	if err != nil {
		return
	}
	defer stmt.Close()

	if !govalidator.IsEmail(data.Email) {
		return user, helper.ErrEmailInvalid
	}
	if len(data.Password) < 10 {
		return user, helper.ErrPasswordNotSecure
	}
	profileNull := sql.NullString{}
	err = stmt.QueryRow(
		data.Email,
		data.Password,
		).Scan(
				&user.ID,
				&user.Name,
				&user.LastName,
				&user.NickName,
				&profileNull,
				&user.Email,
				&user.Password,
			)
	if err != nil {
		return user, helper.ErrDataInvalid
	}
	user.Profile = []byte(profileNull.String)
	return
}