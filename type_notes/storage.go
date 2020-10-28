package type_notes

import (
	"database/sql"
	"fmt"
	"github.com/Leonardo-Antonio/api-notes/helper"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) GetTypes() (TypeNotes []model, err error) {
	stmt, err := s.db.Prepare(sqlGetTypes)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		TypeNote := model{}
		err := rows.Scan(
				&TypeNote.ID,
				&TypeNote.TypeOfNote,
			)
		if err != nil {
			return TypeNotes, err
		}
		TypeNotes = append(TypeNotes, TypeNote)
	}
	return
}

func (s *Storage) CreateActivity(name string) error {
	stmt, err := s.db.Prepare(sqlCreateType)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(helper.StringNull(name))
	if err != nil {
		return helper.ErrDataInvalid
	}

	rA, err := rs.RowsAffected()
	if err != nil {
		return err
	}

	if rA != 1 {
		return helper.ErrInsertionFailed
	}
	return nil
}

func (s *Storage) DeleteType(id int) error {
	stmt, err := s.db.Prepare(sqlDeleteType)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		return helper.ErrDataInvalid
	}
	rA, err := rs.RowsAffected()
	if err != nil {
		return err
	}

	if rA != 1 {
		return helper.ErrRowNotAffected
	}
	return nil
}