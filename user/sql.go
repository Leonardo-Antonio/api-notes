package user

const (
	sqlNewUser = `
		INSERT INTO tb_users ( name, last_name, nick_name, profile, email, password )
		VALUES ( ?, ?, ?, ?, ?, ? )
	`
	sqlLogin = `
		SELECT
			   id, name, last_name, nick_name,
			   profile, email, password
		FROM tb_users
		WHERE
			  email = ?
		and
			  password = ?
	`
)