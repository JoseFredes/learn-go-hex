package right

import (
	"database/sql"
	"fmt"

	"github.com/JoseFredes/learn-go-hex/internal/core/domain"
	"github.com/JoseFredes/learn-go-hex/internal/core/ports"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) (ports.UserRepository, error) {
	return &MySQLUserRepository{db: db}, nil
}

func (r *MySQLUserRepository) GetUserByID(id int) (*domain.User, error) {
	query := "SELECT id, name, last_name, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user domain.User

	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *MySQLUserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users (name, last_name, email) VALUES (?, ?, ?)"

	result, err := r.db.Exec(query, user.Name, user.LastName, user.Email)
	println(query)
	if err != nil {
		return nil, err
	}

	print(result)

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return nil, err
	} else {
		fmt.Println(rows)
	}

	fmt.Println(result)
	user.ID = int(id)

	return user, nil
}