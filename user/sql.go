package user

const (
	sqlSignUp = `INSERT INTO tb_user ( name, last_name, phone, email, password, profile )
					VALUES
						   (?, ?, ?, ?, ?, ?)`
	sqlSignIn = `SELECT
					   id, name, last_name,
					   phone, email, profile
				FROM tb_user
				WHERE
					  email = ?
				AND
					  password = ?`
	sqlDelete = `DELETE FROM tb_user
				WHERE
					  email = ?
				  AND
					  password = ?`
	sqlUpdate = `UPDATE tb_user
				SET
					name = ?, last_name = ?,
					phone = ?, profile = ?
				WHERE id = ?;`
	sqlGetAll = `SELECT 
					   id, name, last_name, phone, 
					   email, password, profile 
				FROM tb_user`
)
