package type_notes

var (
	sqlGetTypes = "SELECT id, type FROM tb_typeOfnote"
	sqlCreateType = "INSERT INTO tb_typeOfnote ( type ) VALUES ( ? )"
	sqlDeleteType = "DELETE FROM tb_typeOfnote WHERE id = ?"
)
