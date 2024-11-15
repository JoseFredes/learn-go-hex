package right

import (
	"database/sql"
	"fmt"

	"github.com/JoseFredes/learn-go-hex/internal/core/domain"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(dsn string) (*MySQLUserRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySQLUserRepository{db: db}, nil
}

func (r *MySQLUserRepository) GetUserByID(id int) (domain.User, error) {
	query := "SELECT id, name, lastname, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user domain.User

	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, fmt.Errorf("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}
