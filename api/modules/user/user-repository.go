package user

import (
	"database/sql"
	"errors"

	"github.com/rof20004/brasil-servicos/database"
)

// List - query all users
func List() ([]User, error) {
	var (
		id   sql.NullInt64
		nome sql.NullString
	)

	sql := "SELECT * FROM users"

	rows, err := database.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		err = rows.Scan(&id, &nome)
		if err != nil {
			return nil, err
		}
		users = append(users, User{ID: id.Int64, Nome: nome.String})
	}

	return users, nil
}

// Insert - create user
func Insert(user *User) error {
	sql := "INSERT INTO users(nome) VALUES(?)"

	_, err := database.DB.Exec(sql, user.Nome)
	if err != nil {
		return err
	}

	return nil
}

// Delete - remove user
func Delete(id string) error {
	sql := "DELETE FROM users WHERE id = ?"

	result, err := database.DB.Exec(sql, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("Usuário não existe na base de dados => ID: " + id)
	}

	return nil
}
