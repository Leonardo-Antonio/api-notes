package notes

const (
	sqlNewNote = `INSERT INTO tb_notes ( title, body, id_user, id_type ) 
					VALUES (?, ?, ?, ?)`
	sqlGetAllNotes = `SELECT
							id, title, body,
							id_user, id_type
						FROM tb_notes
						WHERE id_user = ?`
	sqlDeleteNote = `DELETE FROM tb_notes WHERE id = ? AND id_user = ?`
	sqlUpdateNote = `UPDATE tb_notes 
						SET title = ?, 
						body = ?, 
						id_type = ? 
						WHERE id = ? 
							AND
						id_user = ?`
)
