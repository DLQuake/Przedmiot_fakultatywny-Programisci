package models

import "github.com/jmoiron/sqlx"

//Create new user
func (u *User) Create(db *sqlx.DB) error {
	const query = `
		INSERT INTO users (username, password, email, is_active, created, updated)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Username, u.Password, u.Email, u.IsActive)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = uint64(id)
	return nil
}

//Update : update user
func (u *User) Update(db *sqlx.DB) error {
	stmt, err := db.Prepare(`
		UPDATE users
		SET username = ?,
			password = ?,
			is_active = ?,
			updated = NOW()
		WHERE id = ?
	`)
	_, err = stmt.Exec(u.Username, u.Password, u.IsActive, u.ID)
	return err
}

//Delete : delete user
func (u *User) Delete(db *sqlx.DB) (bool, error) {
	stmt, err := db.Prepare(`DELETE FROM users WHERE id = ?`)
	_, err = stmt.Exec(u.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
