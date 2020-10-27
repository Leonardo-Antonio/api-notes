package helper

import "database/sql"

func StringNull(value string) sql.NullString {
	null := sql.NullString{String: value}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func Int32Null(value int32) sql.NullInt32 {
	null := sql.NullInt32{Int32: value}
	if null.Int32 != 0 {
		null.Valid = true
	}
	return null
}