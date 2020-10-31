package notes

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

func (s *Storage) NewNote(note model) error {
	stmt, err := s.db.Prepare(sqlNewNote)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if note.IDUser < 0 || note.IDType < 0 {
		return helper.ErrIDInvalid
	}

	rs, err := stmt.Exec(
		helper.StringNull(note.Title),
		helper.StringNull(note.Body),
		note.IDUser,
		note.IDType,
	)
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

func (s *Storage) DeleteNote(IDS model) error {
	stmt, err := s.db.Prepare(sqlDeleteNote)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if IDS.IDUser < 0 || IDS.IDType < 0 {
		return helper.ErrIDInvalid
	}

	rs, err := stmt.Exec(IDS.ID, IDS.IDUser)
	if err != nil {
		return helper.ErrIDInvalid
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

func (s *Storage) UpdateNote(note model) error {
	stmt, err := s.db.Prepare(sqlUpdateNote)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if note.IDUser < 0 || note.IDType < 0 {
		return helper.ErrIDInvalid
	}
	rs, err := stmt.Exec(
		helper.StringNull(note.Title),
		helper.StringNull(note.Body),
		note.IDType,
		note.ID,
		note.IDUser,
	)
	if err != nil {
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

func (s *Storage) GetAllNotes(ID int) (notes []model, err error) {
	stmt, err := s.db.Prepare(sqlGetAllNotes)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(ID)
	if err != nil {
		return
	}
	for rows.Next() {
		note := model{}
		err := rows.Scan(
			&note.ID,
			&note.Title,
			&note.Body,
			&note.IDUser,
			&note.IDType,
		)
		if err != nil {
			return notes, err
		}
		notes = append(notes, note)
	}
	return
}
