package user

import (
	"database/sql"
	"github.com/Leonardo-Antonio/api-notes/helper"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) CreateUser(user User) error {
	stmt, err := s.db.Prepare(sqlNewUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

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

func (s *Storage) Login() error {
	stmt, err := s.db.Prepare(sqlLogin)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}